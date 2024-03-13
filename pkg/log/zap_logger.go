package log

import "go.uber.org/zap"

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() (Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return &ZapLogger{
		logger: logger,
	}, nil
}

func (z *ZapLogger) Info(message string, args ...LogArgs) {
	z.logger.Info(message, convertLogArgsToZapFields(args)...)
}

func (z *ZapLogger) Error(message string, args ...LogArgs) {
	z.logger.Error(message, convertLogArgsToZapFields(args)...)
}

func (z *ZapLogger) Warn(message string, args ...LogArgs) {
	z.logger.Warn(message, convertLogArgsToZapFields(args)...)
}

func convertLogArgsToZapFields(args []LogArgs) []zap.Field {
	var zapFields []zap.Field
	for _, arg := range args {
		zapFiled := zap.Any(arg.Key, arg.Value)
		zapFields = append(zapFields, zapFiled)
	}
	return zapFields
}
