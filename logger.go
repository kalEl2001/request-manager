package main

import (
    "net"
	"os"

	"github.com/bshuster-repo/logrus-logstash-hook"
    "github.com/sirupsen/logrus"
)

func initLogger() *logrus.Logger {
	LOGSTASH_TCP := os.Getenv("LOGSTASH_URL")
	if len(LOGSTASH_TCP) == 0 {
		LOGSTASH_TCP = "elk.faishol.net:5000"
	}

	log := logrus.New()
    conn, err := net.Dial("tcp", LOGSTASH_TCP) 
    if err != nil {
        log.Fatal(err)
    }
    hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type": "logs",
		"service": "request manager",
	}))
    log.Hooks.Add(hook)
	return log
}

func createLog(log *logrus.Logger, level string, message string, fields map[string]interface{}) {
	if log == nil {
		log = initLogger()
	}

	ctx := log.WithFields(fields)

	if level == "Warn" {
		ctx.Warn(message)
	} else if level == "Error" {
		ctx.Error(message)
	} else if level == "Info" {
		ctx.Info(message)
	} else if level == "Fatal" {
		ctx.Fatal(message)
	} else if level == "Panic" {
		ctx.Panic(message)
	} else if level == "Debug" {
		ctx.Debug(message)
	}
}