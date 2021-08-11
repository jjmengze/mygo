package gorm

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gorm.io/gorm"
	"strings"
)

const (
	dbTableKey        = attribute.Key("db.sql.table")
	dbRowsAffectedKey = attribute.Key("db.rows_affected")
	dbOperationKey    = semconv.DBOperationKey
	dbStatementKey    = semconv.DBStatementKey
)

func dbTable(name string) attribute.KeyValue {
	return dbTableKey.String(name)
}

func dbStatement(stmt string) attribute.KeyValue {
	return dbStatementKey.String(stmt)
}

func dbCount(n int64) attribute.KeyValue {
	return dbRowsAffectedKey.Int64(n)
}

func dbOperation(op string) attribute.KeyValue {
	return dbOperationKey.String(op)
}


// extractQuery SQL Statement explain to species SQL e.g. (MySQL,Postgresql e.t.c.)
func extractQuery(tx *gorm.DB) string {
	return tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)
}

// operationForQuery split operation name
func operationForQuery(query string, op string) string {
	if op != "" {
		return op
	}

	return strings.ToUpper(strings.Split(query, " ")[0])
}
