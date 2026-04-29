package logger

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

type SimpleLogger struct{}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) Info(msg string) {
	println("[INFO] " + msg)
}

func (l *SimpleLogger) Debug(msg string) {
	println("[DEBUG] " + msg)
}

func (l *SimpleLogger) Error(msg string) {
	println("[ERROR] " + msg)
}