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
)

const (
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
	debugL = log.New(DefaultOpts.Output, prefDebug, log.LstdFlags)
	infoL  = log.New(DefaultOpts.Output, prefInfo, log.LstdFlags)
	errL   = log.New(DefaultOpts.Output, prefError, log.LstdFlags)
	warnL  = log.New(DefaultOpts.Output, prefWarn, log.LstdFlags)
)

type Opts struct {
	Level         int
	Output        io.Writer
	IncludeSource bool
}

type Logger struct {
	opts *Opts
}

func init() {
	SetLevel(DefaultOpts.Level)
}

func New(opts *Opts) *Logger {
	flags := log.LstdFlags
	if opts.IncludeSource {
		flags = log.LstdFlags | log.Lshortfile
	}
	debugL.SetFlags(flags)
	infoL.SetFlags(flags)
	errL.SetFlags(flags)
	warnL.SetFlags(flags)

	return &Logger{
		opts: opts,
	}
}

func (l *Logger) Debug(args ...any) {
	if l.opts.Level <= LevelDebug {
		debugL.Println(args...)
	}
}

func (l *Logger) Debugf(msg string, args ...any) {
	if l.opts.Level <= LevelDebug {
		debugL.Printf(msg, args...)
	}
}

func (l *Logger) Info(args ...any) {
	if l.opts.Level <= LevelInfo {
		infoL.Println(args...)
	}
}

func (l *Logger) Infof(msg string, args ...any) {
	if l.opts.Level <= LevelInfo {
		infoL.Printf(msg, args...)
	}
}

func (l *Logger) Error(err error) {
	if l.opts.Level <= LevelError {
		errL.Println(err.Error())
	}
}

func (l *Logger) Errorf(msg string, args ...any) {
	if l.opts.Level <= LevelError {
		errL.Printf(msg, args...)
	}
}

func (l *Logger) Warn(args ...any) {
	if l.opts.Level <= LevelError {
		warnL.Println(args...)
	}
}

func (l *Logger) Warnf(msg string, args ...any) {
	if l.opts.Level <= LevelWarn {
		warnL.Printf(msg, args...)
	}
}

func (l *Logger) Fatal(args ...any) {
	log.Fatal(args...)
}

func (l *Logger) Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args...)
}

func (l *Logger) Panic(args ...any) {
	log.Panic(args...)
}

func (l *Logger) Panicf(msg string, args ...any) {
	log.Panicf(msg, args...)
}

func Debug(args ...any) {
	if DefaultOpts.Level <= LevelDebug {
		debugL.Println(args...)
	}
}

func Debugf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelDebug {
		debugL.Printf(msg, args...)
	}
}

func Info(args ...any) {
	if DefaultOpts.Level <= LevelInfo {
		infoL.Println(args...)
	}
}

func Infof(msg string, args ...any) {
	if DefaultOpts.Level <= LevelInfo {
		infoL.Printf(msg, args...)
	}
}

func Error(err error) {
	if DefaultOpts.Level <= LevelError {
		errL.Println(err.Error())
	}
}

func Errorf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelError {
		errL.Printf(msg, args...)
	}
}

func Warn(args ...any) {
	if DefaultOpts.Level <= LevelError {
		warnL.Println(args...)
	}
}

func Warnf(msg string, args ...any) {
	if DefaultOpts.Level <= LevelWarn {
		warnL.Printf(msg, args...)
	}
}

func Fatal(args ...any) {
	log.Fatal(args...)
}

func Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args...)
}

func Panic(args ...any) {
	log.Panic(args...)
}

func Panicf(msg string, args ...any) {
	log.Panicf(msg, args...)
}

func SetLevel(level int) {
	DefaultOpts.Level = level
}

func SetOutput(out io.Writer) {
	log.SetOutput(out)
	DefaultOpts.Output = out
}

// SetIncludeSource is not properly implemented yet
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

func ParseLevel(level string) int {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "error":
		return LevelError
	case "warn":
		return LevelWarn
	default:
		panic("unmatched level")
	}
}
