package pattern

/*
	Реализовать паттерн «цепочка вызовов».
	Преимущества:
		Ослабление связанности между отправителем запроса и его обработчиками.
		Возможность добавления новых обработчиков в цепочку без изменения существующего кода.
	Недостатки:
		Возможная неопределенность, будет ли обработан запрос.
		Может быть сложно отследить и отладить, как запрос проходит через цепочку обработчиков.
	Использование:
		Обработка запросов в веб-приложениях.
		Логирование и мониторинг.
		Обработка событий в графических интерфейсах.
		Модели управления доступом.
*/
import "fmt"

// LogLevel определяет уровни логирования.
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

// Logger интерфейс для логгеров.
type Logger interface {
	SetNext(logger Logger)
	LogMessage(level LogLevel, message string)
}

// BaseLogger базовый логгер с реализацией цепочки.
type BaseLogger struct {
	next Logger
}

func (l *BaseLogger) SetNext(logger Logger) {
	l.next = logger
}

func (l *BaseLogger) LogNext(level LogLevel, message string) {
	if l.next != nil {
		l.next.LogMessage(level, message)
	}
}

// InfoLogger логгер для информационных сообщений.
type InfoLogger struct {
	BaseLogger
}

func (l *InfoLogger) LogMessage(level LogLevel, message string) {
	if level == INFO {
		fmt.Printf("INFO: %s\n", message)
	}
	l.LogNext(level, message)
}

// WarnLogger логгер для предупреждений.
type WarnLogger struct {
	BaseLogger
}

func (l *WarnLogger) LogMessage(level LogLevel, message string) {
	if level == WARN {
		fmt.Printf("WARN: %s\n", message)
	}
	l.LogNext(level, message)
}

// ErrorLogger логгер для ошибок.
type ErrorLogger struct {
	BaseLogger
}

func (l *ErrorLogger) LogMessage(level LogLevel, message string) {
	if level == ERROR {
		fmt.Printf("ERROR: %s\n", message)
	}
	l.LogNext(level, message)
}
