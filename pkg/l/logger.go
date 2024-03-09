package l

type Logger interface {
	Info(message string, args ...LogArgs)
	Error(message string, args ...LogArgs)
	Warn(message string, args ...LogArgs)
}

type LogArgs struct {
	Key   string
	Value any
}
