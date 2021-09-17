package template

import "fmt"

type HTTPDownloader struct {
	*DownloadTemplate
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.DownloadTemplate = template
	return template
}

func (hd *HTTPDownloader) Download() {
	fmt.Printf("download %s via http\n", hd.DownloadTemplate.uri)
}

func (hd *HTTPDownloader) Ping() {
	fmt.Print("default Check network environment via ping\n")
}

func (hd *HTTPDownloader) Save() {
	fmt.Print("http Save\n")
}

func (hd *HTTPDownloader) Check() {
	fmt.Print("http Check file checksum \n")
}
