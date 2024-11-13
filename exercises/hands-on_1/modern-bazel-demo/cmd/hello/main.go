// cmd/hello/main.go
package main

import (
    "example.com/modern-bazel-demo/pkg/calculator"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func newLogger() *zap.Logger {
    config := zap.Config{
        Encoding:         "json",
        Level:           zap.NewAtomicLevelAt(zap.DebugLevel),
        OutputPaths:     []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
        EncoderConfig: zapcore.EncoderConfig{
            MessageKey:     "message",
            LevelKey:      "level",
            TimeKey:       "time",
            EncodeLevel:   zapcore.CapitalLevelEncoder,
            EncodeTime:    zapcore.ISO8601TimeEncoder,
            EncodeCaller:  zapcore.ShortCallerEncoder,
        },
    }
    
    logger, _ := config.Build()
    return logger
}

func main() {
    logger := newLogger()
    defer logger.Sync()

    logger.Info("Starting calculator application")
    
    // Perform multiple calculations
    numbers := [][2]int{{5, 3}, {10, 7}, {15, 8}}
    
    for _, pair := range numbers {
        a, b := pair[0], pair[1]
        
        addResult := calculator.Add(a, b)
        logger.Info("Addition performed",
            zap.Int("a", a),
            zap.Int("b", b),
            zap.Int("result", addResult),
        )
        
        subResult := calculator.Subtract(a, b)
        logger.Info("Subtraction performed",
            zap.Int("a", a),
            zap.Int("b", b),
            zap.Int("result", subResult),
        )
    }
    
    logger.Info("Calculator application finished")
}