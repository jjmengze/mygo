package main

import "github.com/jjmengze/mygo/pkg/design-pattern/behavioral/template"

func main() {

	var downloader template.Downloader
	downloader = template.NewHTTPDownloader()

	downloader.Download("http.google.com")

	downloader = template.NewFTPDownloader()
	downloader.Download("ftp://google.com/hello.zip\"")

}
