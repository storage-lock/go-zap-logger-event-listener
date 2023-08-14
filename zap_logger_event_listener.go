package zap_logger_event_listener

import (
	"context"
	"github.com/storage-lock/go-events"
	"go.uber.org/zap"
)

type ZapLoggerListener struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

var _ events.Listener = &ZapLoggerListener{}

const ZapLoggerListenerName = "zap-logger-listener"

func NewZapLoggerListener() *ZapLoggerListener {
	return &ZapLoggerListener{}
}

func (x *ZapLoggerListener) Name() string {
	return ZapLoggerListenerName
}

func (x *ZapLoggerListener) On(ctx context.Context, e *events.Event) {

	if x.sugaredLogger != nil {
		x.sugaredLogger.Info(e.ToJsonString())
	}

	if x.logger != nil {
		x.logger.Info(e.ToJsonString())
	}

}

func (x *ZapLoggerListener) SetLogger(logger *zap.Logger) *ZapLoggerListener {
	x.logger = logger
	return x
}

func (x *ZapLoggerListener) SetSugaredLogger(sugaredLogger *zap.SugaredLogger) *ZapLoggerListener {
	x.sugaredLogger = sugaredLogger
	return x
}
