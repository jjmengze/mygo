# 單利模式-singleton Pattern

透過 `sync.Once` 確保只會觸發一次物件生成，由於 instance 為 private 不能直接拿來使用，外部要使用物件時必須透過 `GetInstance` 生成，之後復用同一個 instance 即可。