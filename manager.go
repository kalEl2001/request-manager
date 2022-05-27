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
		// TODO: create download job
	}
	// TODO: send update to file provider

	dbConn.Create(&Request{
		User: user,
		Slug: slug,
		Status: numFile,
		NumFiles: numFile,
		Files: fileQueues,
	})
	dbConn.Commit()
	infoLog("Receive create request", body)
}

func downloadResponse(id int) {
	var req Request
	var fileQueue FileQueue
	dbConn.Where("id = ?", id).First(&fileQueue)
	dbConn.Where("id = ?", fileQueue.RequestId).First(&req)

	if fileQueue.Status == false {
		req.Status -= 1
	}
	fileQueue.Status = true

	dbConn.Save(&fileQueue)
	dbConn.Save(&req)
	dbConn.Commit()

	// TODO: Send update to file provider
	// TODO: Check if req.status = 0, then create compress job

	infoLog("Receive download request", map[string]interface{}{"correlationId": id})
}

func compressResponse(id int, body map[string]interface{}) {
	var req Request
	dbConn.Where("id = ?", id).First(&req)
	req.Status = -1
	req.OutputPath = body["folder"].(string)

	dbConn.Save(&req)
	dbConn.Commit()

	// TODO: Send update to file provider

	infoLog("Receive compress request", map[string]interface{}{"correlationId": id})
}