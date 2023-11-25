package helper

import "net/http"

func SupportsRange(url string) (bool, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Range", "bytes=0-0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusPartialContent, nil
}
