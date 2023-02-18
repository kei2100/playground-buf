package example_marshal_zap

import (
	foov1 "github.com/kei2100/playground-buf/example_marshal_zap/gen/go/foo/v1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestFoo(t *testing.T) {
	var msg zapcore.ObjectMarshaler = &foov1.FooMessage{
		Message:       "message",
		SecretMessage: "secret",
	}
	log, _ := zap.NewProduction()
	log.Info("write msg", zap.Object("msg", msg))
	// e.g
	// {
	//  "level": "info",
	//  "ts": 1676454357.246934,
	//  "caller": "example_marshal_zap/foo_test.go:16",
	//  "msg": {
	//    "message": "message",
	//    "secret_message": "[MASKED]"
	//  }
	//}
}
