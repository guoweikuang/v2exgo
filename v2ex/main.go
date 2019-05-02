package main

import (
	"github.com/hashicorp/logutils"
	"log"
	"os"
	"github.com/guoweikuang/v2exgo/cmd"
)

func main() {
	filter := &logutils.LevelFilter{
		Levels: []logutils.LogLevel{"D", "E", "W", "I"},
		MinLevel: logutils.LogLevel("W"),
		Writer: os.Stderr,
	}

	log.SetOutput(filter)
	cmd.Execute()
}