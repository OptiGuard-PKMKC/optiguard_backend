package helpers

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func StoreImage(imageBlob string) (string, error) {
	// Decode the base64 string
	imageData, err := base64.StdEncoding.DecodeString(imageBlob)
	if err != nil {
		return "", fmt.Errorf("failed to decode image blob: %v", err)
	}

	// Generate a unique file name using the current timestamp
	fileName := fmt.Sprintf("image_%d.jpg", time.Now().UnixNano())
	filePath := filepath.Join("storage/images/fundus", fileName)

	// Ensure the directory exists
	if err := os.MkdirAll("storage/images/fundus", os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Write the image data to the file
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to write image file: %v", err)
	}

	// Return the file path
	return filePath, nil
}
