package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Part struct {
	Start, End int64
}

type Progress struct {
	Total uint64
	Lock  sync.Mutex
}

func (p *Progress) Add(n uint64) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	p.Total += n
}

func (p *Progress) Print() {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	fmt.Printf("\rDownloading... %v bytes complete", p.Total)
}

type WriteCounter struct {
	Size     uint64
	Progress *Progress
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Progress.Add(uint64(n))
	fmt.Printf("progress is %f%% \n", float64(wc.Progress.Total)/float64(wc.Size)*100)
	return n, nil
}

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 文件原始大小
	size := resp.ContentLength

	counter := &WriteCounter{
		Size: uint64(size),
	}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		return err
	}

	fmt.Print("\n")
	return nil
}
