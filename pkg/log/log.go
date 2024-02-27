package log

import (
	"database/sql"
	"fmt"
	"time"
	//loggerMeli "github.com/melisource/fury_go-meli-toolkit-goutils/logger"
)

var Logger LoggerIntrfc

type logger struct {
	db *sql.DB
}

func Init(db *sql.DB) LoggerIntrfc {
	return logger{db}
}

type LoggerIntrfc interface {
	Success(log string, duration int64, timestamp time.Time)
	Error(log string, duration int64, timestamp time.Time)
}

func (l logger) Success(log string, duration int64, timestamp time.Time) {
	query := "INSERT INTO logs (message, duration_ms, timestamp) VALUES (?, ?, ?);"
	stmt, err := l.db.Prepare(query)
	if err != nil {
		return
	}

	if _, err := stmt.Exec(fmt.Sprintf("Success - %s", log), int(duration), timestamp); err != nil {
		return
	}

	//loggerMeli.Info(fmt.Sprintf("Success - %s", log))
}

func (l logger) Error(log string, duration int64, timestamp time.Time) {
	query := "INSERT INTO logs (message, duration_ms, timestamp) VALUES (?, ?, ?);"
	stmt, err := l.db.Prepare(query)
	if err != nil {
		return
	}

	if _, err := stmt.Exec(fmt.Sprintf("Error - %s", log), int(duration), timestamp); err != nil {
		return
	}

	//loggerMeli.Info(fmt.Sprintf("Error - %s", log))
}
