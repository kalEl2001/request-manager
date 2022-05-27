package main

import (
	"fmt"
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
	
	req := Request{
		User: user,
		Slug: slug,
		Status: numFile,
		NumFiles: numFile,
		Files: fileQueues,
	}
	result := dbConn.Create(&req)
	errorLog(result.Error, "Failed to create object")
	infoLog("Receive create request", body)

	for _, val := range req.Files {
		createDownloadJobMessage(req.Slug, val.Link, val.ID)
	}
	updateStatusFileProvider(req.ID, "update_progress", "Starting to download...")
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

	updateString := fmt.Sprintf("%d of %d file(s) downloaded", req.NumFiles - req.Status, req.NumFiles)
	updateStatusFileProvider(req.ID, "update_progress", updateString)
	if req.status == 0 {
		createCompressJobMessage(req.Slug, req.ID)
		updateStatusFileProvider(req.ID, "update_progress", "Compressing...")
	}

	infoLog("Receive download request", map[string]interface{}{"correlationId": id})
}

func compressResponse(id int, body map[string]interface{}) {
	var req Request
	dbConn.Where("id = ?", id).First(&req)
	req.Status = -1
	req.OutputPath = body["folder"].(string)

	dbConn.Save(&req)

	updateStatusFileProvider(req.ID, "update_url", req.OutputPath)

	infoLog("Receive compress request", map[string]interface{}{"correlationId": id})
}