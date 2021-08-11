package gorm

import (
	"fmt"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func (op *OtelPlugin) before(operation string) gormHookFunc {
	return func(tx *gorm.DB) {
		tx.Statement.Context, _ = op.tracer.Start(tx.Statement.Context, op.spanName(tx, operation), oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	}
}

func (op *OtelPlugin) after(operation string) gormHookFunc {
	return func(tx *gorm.DB) {
		span := oteltrace.SpanFromContext(tx.Statement.Context)
		if !span.IsRecording() {
			// skip the reporting if not recording
			return
		}
		defer span.End()

		// Error
		if tx.Error != nil {
			span.SetStatus(codes.Error, tx.Error.Error())
		}

		// extract the db operation
		query := extractQuery(tx)
		operation = operationForQuery(query, operation)

		if tx.Statement.Table != "" {
			span.SetAttributes(dbTable(tx.Statement.Table))
		}

		span.SetAttributes(
			dbStatement(query),
			dbOperation(operation),
			dbCount(tx.Statement.RowsAffected),
		)
	}
}

func (op *OtelPlugin) spanName(tx *gorm.DB, operation string) string {
	query := extractQuery(tx)
	operation = operationForQuery(query, operation)

	target := op.cfg.dbName
	if target == "" {
		target = tx.Dialector.Name()
	}

	if tx.Statement != nil && tx.Statement.Table != "" {
		target += "." + tx.Statement.Table
	}

	return fmt.Sprintf("%s %s", operation, target)
}

