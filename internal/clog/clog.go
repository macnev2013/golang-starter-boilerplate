package clog

import (
	"log"
	"os"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
)

type clog struct {
	cfg      config.Config
	warnLog  *log.Logger
	debugLog *log.Logger
	errorLog *log.Logger
}

func NewLogger(cfg config.Config) core.Clog {
	return &clog{
		cfg:      cfg,
		warnLog:  log.New(os.Stdout, "WARN\t", log.Ldate|log.Ltime),
		debugLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *clog) Warnf(msg string, args ...interface{}) {
	c.warnLog.Printf(msg, args...)
}
func (c *clog) Debugf(msg string, args ...interface{}) {
	c.debugLog.Printf(msg, args...)
}
func (c *clog) Errorf(msg string, args ...interface{}) {
	c.errorLog.Printf(msg, args...)
}
