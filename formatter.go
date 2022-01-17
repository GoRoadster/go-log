package log

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	defaultLogFormat       = "[%lvl%]: %time% - %msg%"
	defaultTimestampFormat = "02-Jan-06 15:04:05 MST"
)

type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.UTC().Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	output += getFieldsString(entry.Data)

	output += "\n"
	return []byte(output), nil
}

func getFieldsString(fields logrus.Fields) string {
	var output string
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := fields[key]
		switch v := val.(type) {
		case error:
			output += fmt.Sprintf("\t\t%s = %s", key, v.Error())
		case string:
			output += fmt.Sprintf("\t\t%s = %s", key, v)
		case int:
			s := strconv.Itoa(v)
			output += fmt.Sprintf("\t\t%s = %s", key, s)
		case bool:
			s := strconv.FormatBool(v)
			output += fmt.Sprintf("\t\t%s = %s", key, s)
		}
	}

	return output
}
