package template

import "fmt"

type FTPDownloader struct {
	*DownloadTemplate
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.DownloadTemplate = template
	return template
}

func (hd *FTPDownloader) Download() {
	fmt.Printf("download %s via ftp\n", hd.uri)
}
