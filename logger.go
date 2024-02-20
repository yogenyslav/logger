package logger

import (
	"io"
	"log"
	"os"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelError
	LevelWarn
)

var (
	defaultOpts = &Opts{
		Level:         LevelDebug,
		Output:        os.Stdout,
		IncludeSource: false,
	}
)

type Opts struct {
	Level         int
	Output        io.Writer
	IncludeSource bool
}

type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
	opts        *Opts
}

func New(opts *Opts) *Logger {
	flags := log.LstdFlags
	if opts.IncludeSource {
		flags = log.LstdFlags | log.Lshortfile
	}
	debug := log.New(opts.Output, "DEBUG:", flags)
	info := log.New(opts.Output, "INFO:", flags)
	err := log.New(opts.Output, "ERROR:", flags)
	warn := log.New(opts.Output, "WARN:", flags)

	return &Logger{
		debugLogger: debug,
		infoLogger:  info,
		errorLogger: err,
		warnLogger:  warn,
		opts:        opts,
	}
}

func (l *Logger) Debug(args ...any) {
	if l.opts.Level <= LevelDebug {
		l.debugLogger.Println(args)
	}
}

func (l *Logger) Debugf(msg string, args ...any) {
	if l.opts.Level <= LevelDebug {
		l.debugLogger.Printf(msg, args)
	}
}

func (l *Logger) Info(args ...any) {
	if l.opts.Level <= LevelInfo {
		l.infoLogger.Println(args)
	}
}

func (l *Logger) Infof(msg string, args ...any) {
	if l.opts.Level <= LevelInfo {
		l.infoLogger.Printf(msg, args)
	}
}

func (l *Logger) Error(err error) {
	if l.opts.Level <= LevelError {
		l.errorLogger.Println(err.Error())
	}
}

func (l *Logger) Errorf(msg string, args ...any) {
	if l.opts.Level <= LevelError {
		l.errorLogger.Printf(msg, args)
	}
}

func (l *Logger) Warn(args ...any) {
	if l.opts.Level <= LevelError {
		l.warnLogger.Println(args)
	}
}

func (l *Logger) Warnf(msg string, args ...any) {
	if l.opts.Level <= LevelWarn {
		l.warnLogger.Printf(msg, args)
	}
}

func (l *Logger) Fatal(args ...any) {
	l.warnLogger.Fatal(args)
}

func (l *Logger) Fatalf(msg string, args ...any) {
	l.warnLogger.Fatalf(msg, args)
}

func (l *Logger) Panic(args ...any) {
	l.warnLogger.Panic(args)
}

func (l *Logger) Panicf(msg string, args ...any) {
	l.warnLogger.Panicf(msg, args)
}

func Debug(args ...any) {
	if defaultOpts.Level <= LevelDebug {
		log.Println(args)
	}
}

func Debugf(msg string, args ...any) {
	if defaultOpts.Level <= LevelDebug {
		log.Printf(msg, args)
	}
}

func Info(args ...any) {
	if defaultOpts.Level <= LevelInfo {
		log.Println(args)
	}
}

func Infof(msg string, args ...any) {
	if defaultOpts.Level <= LevelInfo {
		log.Printf(msg, args)
	}
}

func Error(err error) {
	if defaultOpts.Level <= LevelError {
		log.Println(err.Error())
	}
}

func Errorf(msg string, args ...any) {
	if defaultOpts.Level <= LevelError {
		log.Printf(msg, args)
	}
}

func Warn(args ...any) {
	if defaultOpts.Level <= LevelError {
		log.Println(args)
	}
}

func Warnf(msg string, args ...any) {
	if defaultOpts.Level <= LevelWarn {
		log.Printf(msg, args)
	}
}

func Fatal(args ...any) {
	log.Fatal(args)
}

func Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args)
}

func Panic(args ...any) {
	log.Panic(args)
}

func Panicf(msg string, args ...any) {
	log.Panicf(msg, args)
}

func SetLevel(level int) {
	defaultOpts.Level = level
}

func SetOutput(out io.Writer) {
	defaultOpts.Output = out
}

func SetIncludeSource(include bool) {
	defaultOpts.IncludeSource = include
}
