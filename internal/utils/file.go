package utils

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func UnzipFile(zipFile, destPath string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		targetPath := filepath.Join(destPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(targetPath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
				return err
			}

			targetFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err = io.Copy(targetFile, src); err != nil {
				return err
			}
		}
	}
	return nil
}
