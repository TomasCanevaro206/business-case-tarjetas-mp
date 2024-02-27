package log

import (
	"database/sql"
	"time"
)

type LoggerDummy struct {
	db *sql.DB
}

func NewLoggerDummy() LoggerIntrfc {
	return LoggerDummy{db: nil}
}

func (l LoggerDummy) Success(log string, duration int64, timestamp time.Time) {
}

func (l LoggerDummy) Error(log string, duration int64, timestamp time.Time) {
}
