package common

import "github.com/fatih/color"

func Notice(msg string) {
	noticeMsg := color.New(color.Bold, color.FgBlue).PrintlnFunc()

	noticeMsg(msg)
}

func Error(msg string) {
	errorMsg := color.New(color.Bold, color.BgRed, color.FgWhite).PrintlnFunc()

	errorMsg(msg)
}

func Info(msg string) {
	infoMsg := color.New(color.FgYellow).PrintlnFunc()

	infoMsg(msg)
}
