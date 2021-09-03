package cmd

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

var _ tmlog.Logger = (*ZeroLogWrapper)(nil)

// ZeroLogWrapper provides a wrapper around a zerolog.Logger instance. It implements
// Tendermint's Logger interface.
type ZeroLogWrapper struct {
	zerolog.Logger
}

// Info implements Tendermint's Logger interface and logs with level INFO. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Info(msg string, keyVals ...interface{}) {
	z.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements Tendermint's Logger interface and logs with level ERR. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Error(msg string, keyVals ...interface{}) {
	z.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements Tendermint's Logger interface and logs with level DEBUG. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Debug(msg string, keyVals ...interface{}) {
	z.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (z ZeroLogWrapper) With(keyVals ...interface{}) tmlog.Logger {
	return ZeroLogWrapper{z.Logger.With().Fields(getLogFields(keyVals...)).Logger()}
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}

const (
	black = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white

	bold     = 1
	darkgray = 90
)

func logFormatLevel(color bool) zerolog.Formatter {
	return func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = colorize("|TRACE|", blue, color)
			case "debug":
				l = colorize("|DEBUG|", cyan, color)
			case "info":
				l = colorize("|INFO |", green, color)
			case "warn":
				l = colorize("|WARN |", yellow, color)
			case "error":
				l = colorize("|ERROR|", red, color)
			case "fatal":
				l = colorize("|FATAL|", red, color)
			case "panic":
				l = colorize("|PANIC|", red, color)
			default:
				l = colorize("???", bold, color)
			}
		} else {
			if i == nil {
				l = colorize("???", bold, color)
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return l
	}
}

func colorize(s interface{}, c int, enable bool) string {
	if enable {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
	}
	return fmt.Sprintf("%s", s)
}
