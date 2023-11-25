package helper

import (
	"net/url"
	"path"
)

func GetFileName(fileURL string) (string, error) {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}

	return path.Base(parsedURL.Path), nil
}
