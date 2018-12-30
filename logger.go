package main

import "log"

type logger interface {
	Criticalf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type stdLogger struct{}

func newStdLogger() *stdLogger {
	return &stdLogger{}
}

func (l *stdLogger) Criticalf(format string, args ...interface{}) {
	log.Printf("CRITICAL: "+format, args...)
}
func (l *stdLogger) Debugf(format string, args ...interface{}) {
	log.Printf("DEBUG: "+format, args...)
}
func (l *stdLogger) Errorf(format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}
func (l *stdLogger) Infof(format string, args ...interface{}) {
	log.Printf("INFO: "+format, args...)
}
func (l *stdLogger) Warningf(format string, args ...interface{}) {
	log.Printf("WARNING: "+format, args...)
}

var _ logger = (*stdLogger)(nil)
