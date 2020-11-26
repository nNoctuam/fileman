package internal

import "go.uber.org/zap"

var logger = NewLogger()

func NewLogger() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"log.log",
	}
	cfg.Encoding = "console"
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return l.Sugar()
}
