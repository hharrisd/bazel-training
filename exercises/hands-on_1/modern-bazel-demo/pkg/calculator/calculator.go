// pkg/calculator/calculator.go
package calculator

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
    logger, _ = zap.NewProduction()
}

func Add(a, b int) int {
    result := a + b
    logger.Debug("Add operation performed",
        zap.Int("a", a),
        zap.Int("b", b),
        zap.Int("result", result),
    )
    return result
}

func Subtract(a, b int) int {
    result := a - b
    logger.Debug("Subtract operation performed",
        zap.Int("a", a),
        zap.Int("b", b),
        zap.Int("result", result),
    )
    return result
}