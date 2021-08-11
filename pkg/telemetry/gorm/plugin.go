package gorm

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	defaultTracerName = "go.opentelemetry.io/contrib/instrumentation/github.com/go-gorm/gorm/otelgorm"

	callBackBeforeName = "otel:before"
	callBackAfterName  = "otel:after"

	opCreate = "INSERT"
	opQuery  = "SELECT"
	opDelete = "DELETE"
	opUpdate = "UPDATE"
)

type gormHookFunc func(tx *gorm.DB)

type OtelPlugin struct {
	cfg    *config
	tracer trace.Tracer
}

// NewPlugin initialize a new gorm.DB plugin that traces queries
// You may pass optional Options to the function
func NewPlugin(opts ...Option) *OtelPlugin {
	cfg := &config{}
	for _, o := range opts {
		o.apply(cfg)
	}

	if cfg.tracerProvider == nil {
		cfg.tracerProvider = otel.GetTracerProvider()
	}

	return &OtelPlugin{
		cfg: cfg,
		tracer: cfg.tracerProvider.Tracer(
			defaultTracerName,
		),
	}
}

// Name Plugin GORM plugin interface
func (op *OtelPlugin) Name() string {
	return "OpenTelemetryPlugin"
}

// Name Plugin GORM plugin interface
func (op *OtelPlugin) Initialize(db *gorm.DB) error {
	type registerCallback interface {
		Register(name string, fn func(*gorm.DB)) error
	}

	registerHooks := []struct {
		callback registerCallback
		hook     gormHookFunc
		name     string
	}{
		// before hooks
		{callback: db.Callback().Create().Before("gorm:before_create"), hook: op.before(opCreate), name: beforeName("create")},
		{callback: db.Callback().Query().Before("gorm:query"), hook: op.before(opQuery), name: beforeName("query")},
		{callback: db.Callback().Delete().Before("gorm:before_delete"), hook: op.before(opDelete), name: beforeName("delete")},
		{callback: db.Callback().Update().Before("gorm:before_update"), hook: op.before(opUpdate), name: beforeName("update")},
		{callback: db.Callback().Row().Before("gorm:row"), hook: op.before(""), name: beforeName("row")},
		{callback: db.Callback().Raw().Before("gorm:raw"), hook: op.before(""), name: beforeName("raw")},

		// after hooks
		{callback: db.Callback().Create().After("gorm:after_create"), hook: op.after(opCreate), name: afterName("create")},
		{callback: db.Callback().Query().After("gorm:after_query"), hook: op.after(opQuery), name: afterName("select")},
		{callback: db.Callback().Delete().After("gorm:after_delete"), hook: op.after(opDelete), name: afterName("delete")},
		{callback: db.Callback().Update().After("gorm:after_update"), hook: op.after(opUpdate), name: afterName("update")},
		{callback: db.Callback().Row().After("gorm:row"), hook: op.after(""), name: afterName("row")},
		{callback: db.Callback().Raw().After("gorm:raw"), hook: op.after(""), name: afterName("raw")},
	}
	for _, h := range registerHooks {
		if err := h.callback.Register(h.name, h.hook); err != nil {
			return fmt.Errorf("register %s hook: %w", h.name, err)
		}
	}

	return nil
}

// beforeName return callBackBeforeName and call back before operation
func beforeName(name string) string {
	return callBackBeforeName + "_" + name
}

// afterName return callBackBeforeName and call back after operation
func afterName(name string) string {
	return callBackAfterName + "_" + name
}
