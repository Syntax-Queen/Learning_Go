package main

import "fmt"

type LogLevel int

const (
	LevelTrace   LogLevel = iota
	LevelDebug            = 1
	LevelInfo             = 2
	LevelWarning          = 3
	LevelError            = 4
)

var LevelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return LevelNames[l]
}
func printLogLevel(level LogLevel) {
	fmt.Printf(" Log levl: %d %s\n", level, level.String())
}
func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelTrace)
	printLogLevel(LevelError)
	printLogLevel(10)

}