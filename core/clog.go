package core

type Clog interface {
	Warnf(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
}
