package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadPart(url string, part Part, dest string, wg *sync.WaitGroup, progress *Progress) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	rangeHeader := fmt.Sprintf("bytes=%d-%d", part.Start, part.End)
	req.Header.Add("Range", rangeHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	counter := &WriteCounter{
		Progress: progress,
		Size:     uint64(part.End - part.Start + 1),
	}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		panic(err)
	}
}

func MergeFiles(parts []string, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, part := range parts {
		in, err := os.Open(part)
		if err != nil {
			return err
		}

		io.Copy(out, in)
		in.Close()

		os.Remove(part)
	}

	return nil
}

func DownloadFileParts(url string, dest string, parts []Part) {
	var wg sync.WaitGroup
	progress := &Progress{}

	partFiles := make([]string, len(parts))

	for i, part := range parts {
		partFile := fmt.Sprintf("%s.part%d", dest, i)
		partFiles[i] = partFile

		wg.Add(1)
		go DownloadPart(url, part, partFile, &wg, progress)
	}

	wg.Wait()

	MergeFiles(partFiles, dest)
}
