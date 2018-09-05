package functional_options

import (
	"github.com/op/go-logging"
	"runtime"
	"strings"
)

// Logger OMIT
type Logger interface {
	Error(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
}
// Logger# OMIT


// Constructors OMIT
// NewLogger creates new logger for automatically retrieved module name.
func NewLogger() Logger {
	return NewLoggerForModule(getModuleName())
}

// NewLoggerForModule creates new logger for a given module name.
func NewLoggerForModule(module string) Logger {
	return logging.MustGetLogger(module)
}
// Constructors# OMIT


// getModuleName OMIT
func getModuleName() string {
	// 0 - this function
	// 1 - NewLogger
	// 2 - calling module
	const SkipStackFrames = 2

	pc, _, _, ok := runtime.Caller(SkipStackFrames)
	if !ok {
		return defaultLoggerName
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return defaultLoggerName
	}
	// componentName will be equal to something like:
	// dir_to_gohan/gohan/some_dirs/package_name/(*class_name).func_name
	componentName := f.Name()
	componentName = strings.Replace(componentName, "/", ".", -1)
	nameStart := strings.Index(componentName, "gohan.")
	nameStop := strings.LastIndex(componentName, "(") - 1
	if nameStop < 0 {
		nameStop = strings.LastIndex(componentName, ".")
		if nameStop < 0 {
			nameStop = len(componentName)
		}
	}
	if nameStart < 0 {
		nameStart = 0
	}
	return componentName[nameStart:nameStop]
}
// getModuleName# OMIT

// Constructors2 OMIT
func NewLogger() Logger
func NewLoggerForModule(module string) Logger
func NewTracingLogger(traceId string) Logger
func NewTracingLoggerForModule(traceId, module string) Logger
// Constructors2# OMIT

// Constructors3 OMIT
func NewLogger(traceId, module string) Logger {
	if traceId == "" {
		...
	}

	if module == "" {
		...
	}

}

logger := NewLogger("", "")

// Constructors3# OMIT

// Config OMIT
type Config struct {
	Module string
	TraceId string
}

func NewLogger(config ...Config) Logger
// Config# OMIT

// ConfigExample OMIT
defaultLogger := NewLogger()

customLogger := NewLogger(Config{
	Module: "myModule",
})
// ConfigExample# OMIT

// ConfigExampleBad OMIT
badLogger := NewLogger(Config{Module: "myModule"}, Config{TraceId: "abcd-1234"})
// ConfigExampleBad# OMIT


// Call OMIT
log := NewLogger()
log := NewLogger(TraceId(params.TraceID))
log := NewLogger(ModuleName(module), TraceId(request.getTraceID()))
// Call# OMIT

// ConfigObject OMIT
type Options struct {
	moduleName string
	traceId    string
}
// ConfigObject# OMIT

// OptType OMIT
type LoggingOption func(op *Options)
// OptType# OMIT


// ConcreteOptions OMIT
func ModuleName(name string) LoggingOption {
	return func(op *Options) {
		op.moduleName = name
	}
}

func TraceId(id string) LoggingOption {
	return func(op *Options) {
		op.traceId = id
	}
}
// ConcreteOptions# OMIT

// ApplyOptions OMIT
// Method OMIT
func NewLogger(opts ...LoggingOption) Logger {
// Method# OMIT
	// apply options
	options := Options{}
	for _, op := range opts {
		op(&options)
	}

	// handle defaults
	if options.moduleName == "" {
		options.moduleName = getModuleName()
	}

	// build object
	var logger Logger = logging.MustGetLogger(options.moduleName)

	if options.traceId != "" {
		logger = &tracingLogger{logger, options.traceId}
	}

	return logger
}
// ApplyOptions# OMIT
