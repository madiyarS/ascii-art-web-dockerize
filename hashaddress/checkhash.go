package hashaddress

import (
	"fmt"
	"crypto/sha256"
	"os"
	"io"
	"errors"
)


func CalculateFileHash(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to calculate hash: %v", err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// ValidateFileHash validates the file hash against the expected hash.
func ValidateFileHash(filepath, expectedHash string) error {
	hash, err := CalculateFileHash(filepath)
	if err != nil {
		return err
	}
	if hash != expectedHash {
		return errors.New("file integrity check failed: hash mismatch")
	}
	return nil
}
