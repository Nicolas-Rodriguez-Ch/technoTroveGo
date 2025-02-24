package utils

import "github.com/gin-gonic/gin"

type CloudinaryResponse struct {
	url string
}

func ConvertFilesToImageUrls(c *gin.Context) []string {
	files, exist := c.Get("files")
	if !exist {
		return []string{}
	}

	var fileURLs []string
	if fileList, ok := files.([]string); ok {
		for _, file := range fileList {
			fileURLs = append(fileURLs, file)
		}
	}
	return fileURLs
}
