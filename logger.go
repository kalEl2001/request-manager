package main

import (
    "net"
	"os"

	"github.com/bshuster-repo/logrus-logstash-hook"
    "github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func initLogger() {
	LOGSTASH_TCP := os.Getenv("LOGSTASH_URL")
	if len(LOGSTASH_TCP) == 0 {
		LOGSTASH_TCP = "elk.faishol.net:5000"
	}

	logger = logrus.New()
    conn, err := net.Dial("tcp", LOGSTASH_TCP) 
    if err != nil {
        logger.Fatal(err)
    }
    hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"type": "logs",
		"service": "request manager",
	}))
    logger.Hooks.Add(hook)
}

func createLog(level string, message string, fields map[string]interface{}) {
	if logger == nil {
		initLogger()
	}

	ctx := logger.WithFields(fields)

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

func createErrorLog(level string, err error, msg string) {
	if err != nil {
        field := map[string]interface{}{
            "error_msg": err,
        }
        createLog(level, msg, field)
    }
}

func failLog(err error, msg string) {
    createErrorLog("Panic", err, msg)
}

func errorLog(err error, msg string) {
	createErrorLog("Error", err, msg)
}

func warningLog(msg string, field map[string]interface{}) {
	createLog("Warn", msg, field)
}

func infoLog(msg string, field map[string]interface{}) {
	createLog("Info", msg, field)
}