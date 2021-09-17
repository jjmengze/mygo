# 範本模式 template Pattern

範本模式定義一個操作中的執行骨架,將一些步驟延遲到子類別中。


template Pattern 可以把一些業務邏輯中不變的行為移轉到 Super Class 換句話說就是把一個功能拆成多個小步驟，每個小步驟就是一個 method ，然後用一個 template method 把這些小步驟組合起來變成一個功能。
> 這些小步驟可能會是一些預設行為，或是由子類別複寫。


例如下載檔案這個業務邏輯，有幾個步驟需要實作。
1. 確認網路連線
2. 下載檔案
3. 儲存檔案
4. 確認檔案是否損毀

```go
type Implement interface {
	Ping()
	Download()
	Save()
	Check()
}
```

實際上外部使用者只關心這個物件是否能夠執行 `Download` 因此還會再定義另一個 interface 給外部使用者。 

```go
type Downloader interface {
	Download(url string)
}
```
此時會定義一個樣板，用來訂製業務邏輯的流程，
依照下載的範例樣板會描述下載時檢察網路環境、下載檔案、儲存檔案、以及檢查檔案是否損毀的流程。

以及提供給外部呼叫的業務邏輯進入點。

```go
type DownloadTemplate struct {
	//子類實作下載細節
	Implement
	uri string
}

func newTemplate(impl Implement) *DownloadTemplate {
	return &DownloadTemplate{
		//傳入子累實作下載細節的物件
		Implement: impl,
	}
}
//定義業務邏輯的執行流程，與決定哪些步驟為預設行為，以及提供外部呼叫業務邏輯的進入點
func (dt *DownloadTemplate) Download(uri string) {
	dt.uri = uri
	fmt.Print("checking network environment\n")
    //預設行為
	dt.Ping()
	
	
	fmt.Print("prepare downloading\n")
	//由子類複寫下載
	dt.Implement.Download()
    //由子類複寫儲存
	dt.Implement.Save()
	fmt.Print("finish downloading\n")
	
	
	fmt.Print("checking file\n")
    //預設行為
	dt.Check()
}
//預設行為
func (dt *DownloadTemplate) Ping() {
    fmt.Print("default Check network environment via ping\n")
}
//預設行為
func (dt *DownloadTemplate) Save() {
    fmt.Print("default Save\n")
}
//預設行為
func (dt *DownloadTemplate) Check() {
    fmt.Print("default Check file checksum of MD5\n")
}

```


在 template Pattern 的設計之下，已經將業務羅拆分能針對個別的流程實作相對應的物件，例如 FTP 的下載方法只改變了流程中的下載行為，我們可能會這麼定義。

```go
//在 FTPDownloader 遷入模板，因為FTPDownloader是基礎模板的延伸
type FTPDownloader struct {
	*DownloadTemplate
}

func NewFTPDownloader() Downloader {
	//建立FTPDownloader物件
	downloader := &FTPDownloader{}
	//模板內部分實作套用 FTPDownloader 物件
	template := newTemplate(downloader)
	//設定FTPDownloader物件的模板為(模板內部分實作套用 FTPDownloader 物件)
	downloader.DownloadTemplate = template
	return template
}

//複寫模板的下載行為
func (hd *FTPDownloader) Download() {
	fmt.Printf("download %s via ftp\n", hd.uri)
}
```