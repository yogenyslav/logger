package logger

import (
	"io"
	"log"
	"os"
	"strings"
)

const (
	prefDebug = "DEBUG: "
	prefInfo  = "INFO: "
	prefError = "ERROR: "
	prefWarn  = "WARN: "

	LevelDebug = iota
	LevelInfo
	LevelError
	LevelWarn
)

var (
	DefaultOpts = &Opts{
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
	debug := log.New(opts.Output, prefDebug, flags)
	info := log.New(opts.Output, prefInfo, flags)
	err := log.New(opts.Output, prefError, flags)
	warn := log.New(opts.Output, prefWarn, flags)

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
	if DefaultOpts.Level <= LevelDebug {
		log.Println(args)
	}
}

func Debugf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelDebug {
		log.Printf(msg, args)
	}
}

func Info(args ...any) {
	if DefaultOpts.Level <= LevelInfo {
		log.Println(args)
	}
}

func Infof(msg string, args ...any) {
	if DefaultOpts.Level <= LevelInfo {
		log.Printf(msg, args)
	}
}

func Error(err error) {
	if DefaultOpts.Level <= LevelError {
		log.Println(err.Error())
	}
}

func Errorf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelError {
		log.Printf(msg, args)
	}
}

func Warn(args ...any) {
	if DefaultOpts.Level <= LevelError {
		log.Println(args)
	}
}

func Warnf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelWarn {
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
	switch level {
	case LevelDebug:
		log.SetPrefix(prefDebug)
	case LevelInfo:
		log.SetPrefix(prefInfo)
	case LevelError:
		log.SetPrefix(prefError)
	case LevelWarn:
		log.SetPrefix(prefWarn)
	}
	DefaultOpts.Level = level
}

func SetOutput(out io.Writer) {
	log.SetOutput(out)
	DefaultOpts.Output = out
}

func SetIncludeSource(include bool) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	DefaultOpts.IncludeSource = include
}

func SetFileOutput(path string, logConsole bool) {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	var out io.Writer
	if logConsole {
		out = io.MultiWriter(os.Stdout, logFile)
	}
	log.SetOutput(out)
	DefaultOpts.Output = out
}
