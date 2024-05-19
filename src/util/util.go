package util

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Function to create a zip archive from a list of file paths and JSON code
func CreateArchive(outputPath string) (*zip.Writer, *os.File) {

	zipFile, _ := os.Create(outputPath)
	zipWriter := zip.NewWriter(zipFile)

	return zipWriter, zipFile
}

func AddFile(w *zip.Writer, filePath string) FileInfo {
	FileInfo, _ := GetFileInfo(filePath)
	file, _ := os.Open(filePath)
	defer file.Close()
	fw, _ := w.Create(FileInfo.Filename)
	io.Copy(fw, file)

	return FileInfo
}

type FileInfo struct {
	Checksum   string
	DataFormat string
	Filename   string
}

// getFileInfo returns the MD5 checksum, data format, and formatted filename for the given file path
func GetFileInfo(filePath string) (FileInfo, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return FileInfo{}, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a new MD5 hash
	hash := md5.New()

	// Copy the file's content into the hash
	if _, err := io.Copy(hash, file); err != nil {
		return FileInfo{}, fmt.Errorf("failed to calculate checksum: %v", err)
	}

	// Get the MD5 checksum as a hex string
	checksum := hex.EncodeToString(hash.Sum(nil))

	// Get the file extension
	dataFormat := filepath.Ext(filePath)[1:]

	// Format the filename as "checksum + dataFormat"
	filename := checksum + "." + dataFormat

	return FileInfo{
		Checksum:   checksum,
		DataFormat: dataFormat,
		Filename:   filename,
	}, nil
}
