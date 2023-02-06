package logger

import (
	"fast-food-api-client/constants"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Package struct {
	appName  string
	fileName string
}

func Logger(appName string, fileName string) *Package {
	return &Package{
		appName:  appName,
		fileName: fileName,
	}
}

func (l *Package) Error(args ...interface{}) {
	sentryLog(l.appName, l.fileName, constants.TypeError, args...)
}

func (l *Package) Info(args ...interface{}) {
	sentryLog(l.appName, l.fileName, constants.TypeInfo, args...)
}

func (l *Package) Warn(args ...interface{}) {
	sentryLog(l.appName, l.fileName, constants.TypeWarn, args...)
}

func (l *Package) Debug(args ...interface{}) {
	sentryLog(l.appName, l.fileName, constants.TypeDebug, args...)
}

func sentryLog(appName string, fileName string, logType string, args ...interface{}) *time.Time {
	// format args to string message
	message := fmt.Sprint(args...)
	// format package name
	formatPackage(appName, fileName, logType)
	// format log message
	logText, logTime := formatLog(logType, message)
	fmt.Print(logText)

	return &logTime
}

func formatPackage(appName string, fileName string, logType string) {
	switch logType {
	case constants.TypeError:
		colorError := color.New(color.FgRed).SprintFunc()
		fmt.Printf("%s ", colorError(strings.ToUpper(appName), "(", fileName, ")"))
	case constants.TypeWarn:
		colorWarn := color.New(color.FgYellow).SprintFunc()
		fmt.Printf("%s ", colorWarn(strings.ToUpper(appName), "(", fileName, ")"))
	case constants.TypeInfo:
		colorInfo := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("[%s] ", colorInfo(strings.ToUpper(appName), "(", fileName, ")"))
	case constants.TypeDebug:
		colorDebug := color.New(color.FgBlue).SprintFunc()
		fmt.Printf("%s ", colorDebug(strings.ToUpper(appName), "(", fileName, ")"))
	default:
		colorDefault := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("%s ", colorDefault(strings.ToUpper(appName), "(", fileName, ")"))
	}
}

func formatLog(logType string, message string) (string, time.Time) {
	logTime := time.Now().UTC()
	logText := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%03d [%s] --> %v\n",
		logTime.Year(), logTime.Month(), logTime.Day(), logTime.Hour(), logTime.Minute(), logTime.Second(),
		logTime.Nanosecond()/1000000, logType, message)

	return logText, logTime
}
