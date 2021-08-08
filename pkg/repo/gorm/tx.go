package gorm

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"k8s.io/klog"
	errPkg "github.com/jjmengze/mygo/pkg/errors"
	"time"
)

// CtxKey 用來代表 context.Context 的 key
type CtxKey string

// AfterCommitCallback ...
type AfterCommitCallback func() error

const (
	// CtxKeyAfterCommitCallback 註冊 transaction commit 後要執行的動作的 context key
	CtxKeyAfterCommitCallback CtxKey = "afterCommit"
	// DefaultCostMilliSeconds 預設的db預期的執行時間，超過會出警示
	DefaultCostMilliSeconds int64 = 2000
	// DefaultAfterCommitFuncCostMilliSeconds 預設的db after commit func預期的執行時間，超過會出警示
	DefaultAfterCommitFuncCostMilliSeconds int64 = 2000
)

// ExecuteTx 執行一個 Database 交易，如果 `fn` 執行過程中遇到失敗，會自動執行 rollback, 如果成功則會自動 commit,
// 另外需要注意的地方是 db 的 connection 需要用 write 的 connection 來執行 tx
func ExecuteTx(ctx context.Context, db *gorm.DB, exceptCostMilliSeconds int64, fn func(*gorm.DB) error) error {
	var err error
	if exceptCostMilliSeconds == 0 {
		exceptCostMilliSeconds = DefaultCostMilliSeconds
	}
	now := time.Now()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		klog.Warningf("cost_milliseconds", costMilliSeconds)
		if costMilliSeconds > exceptCostMilliSeconds {
			klog.Warningf("db: transaction time too long: %v millsSeconds", costMilliSeconds)
		}
		if err == nil { // trigger after commit
			runAfterCommitCallback(ctx)
		}
	}(now)
	// Start a transactions.
	err = executeInTx(ctx, db, fn)
	if err != nil {
		return err
	}

	return nil
}

func executeInTx(ctx context.Context, db *gorm.DB, fn func(*gorm.DB) error) (err error) {
	panicked := true
	tx := db.Begin()
	if tx.Error != nil {
		return errors.Wrapf(errPkg.ErrInternalError, "executeInTx Begin error %+v", tx.Error.Error())
	}

	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked {
			err = errors.Wrap(errPkg.ErrInternalError, "executeInTx occurs panic, start Rollback transaction")
		} else if err != nil {
			err = errors.Wrapf(err, "executeInTx occurs error, start Rollback transaction.")
		}

		if err != nil {
			klog.Errorf("%+v", err)
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				klog.Errorf("executeInTx Rollback transaction then error again! error: %+v ", rollbackErr)
				return
			}
			return
		}
	}()

	err = fn(tx)

	if err == nil {
		if commitErr := tx.Commit().Error; commitErr != nil {
			klog.Errorf("executeInTx failed to commit, start Rollback transaction. error: %+v", err)
			return commitErr
		}
	}
	panicked = false
	return err
}

func runAfterCommitCallback(ctx context.Context) {
	var err error
	now := time.Now()
	defer func(now time.Time) {
		costMilliSeconds := time.Since(now).Milliseconds()
		klog.Info("db: runAfterCommitCallback cost_milliseconds", costMilliSeconds)
		if costMilliSeconds > DefaultAfterCommitFuncCostMilliSeconds {
			klog.Info("db: runAfterCommitCallback too long: %v1 millsSeconds", costMilliSeconds)
		}
	}(now)
	callbacks := ctx.Value(CtxKeyAfterCommitCallback)
	if callbacks == nil {
		return
	}

	fns, ok := callbacks.([]AfterCommitCallback)
	if !ok {
		klog.Warningf("runAfterCommitCallback convert failed")
		return
	}
	for i := range fns {
		fn := (fns)[i]
		if err = fn(); err != nil {
			klog.Errorf("fail to run after commit callback, err: %s", err.Error())
		}
	}
}

// AddAfterCommitCallback ...
func AddAfterCommitCallback(ctx context.Context, fn AfterCommitCallback) (context.Context, error) {
	if fn == nil {
		klog.Warningf("AddAfterCommitCallback fn can not be nil")
		return ctx, nil
	}

	callbacks := ctx.Value(CtxKeyAfterCommitCallback)
	if callbacks == nil {
		return ctx, errors.WithMessage(errPkg.ErrInternalError, "AddAfterCommitCallback failed, not set CtxKey")
	}

	fns, ok := callbacks.([]AfterCommitCallback)
	if !ok {
		klog.Warningf("AddAfterCommitCallback convert failed")
		return ctx, errors.WithMessage(errPkg.ErrInternalError, "AddAfterCommitCallback convert failed, the callbacks func unexpected")
	}

	fns = append(fns, fn)
	return context.WithValue(ctx, CtxKeyAfterCommitCallback, fns), nil
}

// InitAfterTxCommitCallback init after tx commit callback callback list
func InitAfterTxCommitCallback(ctx context.Context) context.Context {
	var afterCommitFunc []AfterCommitCallback
	return context.WithValue(ctx, CtxKeyAfterCommitCallback, afterCommitFunc)
}
