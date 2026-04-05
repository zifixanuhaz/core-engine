package core_engine

 import (
 	"crypto/rand"
 	"encoding/hex"
 	"fmt"
 	"os"
 	"path/filepath"
 	"strings"
 	"time"
 )

 // GenerateRandomString generates a random string of the specified length.
 func GenerateRandomString(length int) (string, error) {
 	b := make([]byte, length)
 	_, err := rand.Read(b)
 	if err != nil {
 		return "", fmt.Errorf("failed to generate random string: %w", err)
 	}
 	return hex.EncodeToString(b), nil
 }

 // CreateDirectoryIfNotExists creates a directory if it doesn't exist.
 func CreateDirectoryIfNotExists(path string) error {
 	if _, err := os.Stat(path); os.IsNotExist(err) {
 		err := os.MkdirAll(path, 0755)
 		if err != nil {
 			return fmt.Errorf("failed to create directory: %w", err)
 		}
 	}
 	return nil
 }

 // FileExists checks if a file exists at the given path.
 func FileExists(path string) bool {
 	_, err := os.Stat(path)
 	return !os.IsNotExist(err)
 }

 // GetTimestamp returns the current timestamp in a specific format.
 func GetTimestamp() string {
 	return time.Now().Format(time.RFC3339)
 }

 // SanitizeFilename sanitizes a filename by removing invalid characters and limiting length.
 func SanitizeFilename(filename string) string {
 	// Replace spaces with underscores
 	filename = strings.ReplaceAll(filename, " ", "_")

 	// Remove invalid characters (keep alphanumeric, underscore, hyphen, and period)
 	var validChars []rune
 	for _, r := range filename {
 		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' || r == '.' {
 			validChars = append(validChars, r)
 		}
 	}
 	filename = string(validChars)

 	// Limit filename length to avoid issues with some filesystems
 	const maxFilenameLength = 255
 	if len(filename) > maxFilenameLength {
 		filename = filename[:maxFilenameLength]
 	}

 	return filename
 }

 // GetAbsolutePath resolves a relative path to an absolute path.
 func GetAbsolutePath(path string) (string, error) {
 	if filepath.IsAbs(path) {
 		return path, nil
 	}

 	absPath, err := filepath.Abs(path)
 	if err != nil {
 		return "", fmt.Errorf("failed to get absolute path: %w", err)
 	}

 	return absPath, nil
 }