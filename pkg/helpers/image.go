package helpers

import (
	"encoding/base64"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
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

	// Ensure the directory exists with sudo
	if err := exec.Command("sudo", "mkdir", "-p", "storage/images/fundus").Run(); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Write the image data to the file with sudo
	cmd := exec.Command("sudo", "tee", filePath)
	cmd.Stdin = strings.NewReader(string(imageData))
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to write image file: %v", err)
	}

	// Return the file path
	return filePath, nil
}
