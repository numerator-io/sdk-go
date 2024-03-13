package log

type Logger interface {
	Info(message string, args ...LogArgs)
	Error(message string, args ...LogArgs)
	Warn(message string, args ...LogArgs)
}

type LogArgs struct {
	Key   string
	Value any
}

func Any(key string, value any) LogArgs {
	return LogArgs{
		Key:   key,
		Value: value,
	}
}

func Error(err error) LogArgs {
	return LogArgs{
		Key:   "error",
		Value: err,
	}
}
