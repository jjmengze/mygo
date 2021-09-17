package template

import "fmt"

type Downloader interface {
	Download(url string)
}

type DownloadTemplate struct {
	Implement
	uri string
}

type Implement interface {
	Ping()
	Download()
	Save()
	Check()
}

func newTemplate(impl Implement) *DownloadTemplate {
	return &DownloadTemplate{
		Implement: impl,
	}
}

func (dt *DownloadTemplate) Download(uri string) {
	dt.uri = uri
	fmt.Print("checking network environment\n")
	dt.Ping()
	fmt.Print("prepare downloading\n")
	dt.Implement.Download()
	dt.Implement.Save()
	fmt.Print("finish downloading\n")
	fmt.Print("checking file\n")
	dt.Check()
}

func (dt *DownloadTemplate) Ping() {
	fmt.Print("default Check network environment via ping\n")
}

func (dt *DownloadTemplate) Save() {
	fmt.Print("default Save\n")
}

func (dt *DownloadTemplate) Check() {
	fmt.Print("default Check file checksum of MD5\n")
}
