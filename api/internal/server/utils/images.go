package utils

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type IImageService interface {
	SaveImage(id int64, isUser bool, file multipart.File, fileHeader multipart.FileHeader) (string, error)
}

func SaveImage(id int64, isUser bool, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Get the file extension
	fileExtension := strings.Split(fileHeader.Filename, ".")[1]

	// Generate new file name
	newFileName := fmt.Sprintf("%s.%s", uuid.NewV4().String(), fileExtension)

	imagePath := "images"
	if isUser {
		imagePath = filepath.Join(imagePath, "users", fmt.Sprintf("%d", id))
	} else {
		imagePath = filepath.Join(imagePath, "groups", fmt.Sprintf("%d", id))
	}

	// Create folder if not exists
	err := os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		log.Println("Error with creating folder", err)
		return "", err
	}

	// Create new file
	newFile, err := os.Create(fmt.Sprintf("%s/%s", imagePath, newFileName))
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer newFile.Close()

	// Copy the uploaded file to the created file
	if _, err := io.Copy(newFile, file); err != nil {
		log.Println(err)
		return "", err
	}

	return newFileName, nil
}
