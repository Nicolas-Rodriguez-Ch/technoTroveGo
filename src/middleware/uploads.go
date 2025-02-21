package middleware

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func UploadToCloudinary(fileBuffer []byte, fileType string) (*uploader.UploadResult, error) {

	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudinaryApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinaryApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	if cloudinaryCloudName == "" || cloudinaryApiKey == "" || cloudinaryApiSecret == "" {
		return nil, errors.New("Not all neccesarry cloudinary credentials are in the env files, please check")
	}

	cld, err := cloudinary.NewFromParams(cloudinaryCloudName, cloudinaryApiKey, cloudinaryApiSecret)
	if err != nil {
		return nil, err
	}

	base64Str := base64.StdEncoding.EncodeToString(fileBuffer)
	dataUri := fmt.Sprintf("data:%s;base64,%s", fileType, base64Str)

	uploadParams := uploader.UploadParams{
		ResourceType: "auto",
	}

	result, err := cld.Upload.Upload(context.Background(), dataUri, uploadParams)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ProcessFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.Next()
		return
	}

	files := form.File["file"]

	var uploadedFiles []string

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			continue
		}

		defer f.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(f)
		if err != nil {
			continue
		}
		uploadResult, err := UploadToCloudinary(buf.Bytes(), file.Header.Get("Content-Type"))
		if err != nil {
			continue
		}
		uploadedFiles = append(uploadedFiles, uploadResult.SecureURL)
	}
	c.Set("files", uploadedFiles)
	c.Next()
}
