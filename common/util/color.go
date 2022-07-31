package util

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Success           = color.New(color.FgHiGreen).SprintFunc()
	SuccessBackground = color.New(color.BgGreen).SprintFunc()

	Fail           = color.New(color.FgHiRed).SprintFunc()
	FailBackground = color.New(color.BgRed).SprintFunc()

	Warn           = color.New(color.FgHiYellow).SprintFunc()
	WarnBackground = color.New(color.BgYellow).SprintFunc()

	Info           = color.New(color.FgHiWhite).SprintFunc()
	InfoBackground = color.New(color.BgWhite).SprintFunc()
)

func PrintSuccess(a ...interface{}) {
	fmt.Print(Success(a...))
}

func PrintDangerous(a ...interface{}) {
	fmt.Print(Fail(a...))
}

func PrintWarn(a ...interface{}) {
	fmt.Print(Warn(a...))
}

func PrintDefault(a ...interface{}) {
	fmt.Print(Info(a...))
}

func PrintlnSuccess(a ...interface{}) {
	fmt.Println(Success(a...))
}

func PrintlnDangerous(a ...interface{}) {
	fmt.Println(Fail(a...))
}

func PrintlnWarn(a ...interface{}) {
	fmt.Println(Warn(a...))
}

func PrintlnDefault(a ...interface{}) {
	fmt.Println(Info(a...))
}

func PrintfSuccess(format string, a ...interface{}) {
	fmt.Fprintf(color.Output, format, Success(a...))
}

func PrintfDangerous(format string, a ...interface{}) {
	fmt.Fprintf(color.Output, format, Fail(a...))
}

func PrintfWarn(format string, a ...interface{}) {
	fmt.Fprintf(color.Output, format, Warn(a...))
}

func PrintfDefault(format string, a ...interface{}) {
	fmt.Fprintf(color.Output, format, Info(a...))
}

