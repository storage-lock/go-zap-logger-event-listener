package zap_logger_event_listener

import (
	"context"
	"testing"

	"github.com/storage-lock/go-events"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapListener(t *testing.T) {
	listener := NewZapLoggerListener()
	if listener.Name() != ZapLoggerListenerName {
		t.Fatal(listener.Name())
	}
	listener.On(context.Background(), events.NewEvent("silent"))

	core, recorded := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)
	if listener.SetLogger(logger) != listener || listener.SetSugaredLogger(logger.Sugar()) != listener {
		t.Fatal("setters lost receiver")
	}
	listener.On(context.Background(), events.NewEvent("lock"))
	if recorded.Len() != 2 {
		t.Fatalf("records=%d", recorded.Len())
	}
}
