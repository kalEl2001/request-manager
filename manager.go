package main

import (
	"github.com/google/uuid"
)

func createRequest(body map[string]interface{}) {
	slug := uuid.New().String()
	user := body["username"].(string)
	links := body["data"].([]interface{})
	numFile := len(links)

	var fileQueues []FileQueue
	for _, val := range links {
		tmp := FileQueue{Link: val.(string), Status: false}
		fileQueues = append(fileQueues, tmp)
	}

	dbConn.Create(&Request{
		User: user,
		Slug: slug,
		Status: numFile,
		NumFiles: numFile,
		Files: fileQueues,
	})
	infoLog("Receive create request", body)
	dbConn.Commit()
}