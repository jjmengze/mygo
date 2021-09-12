# 簡單工廠模式 simple factory Pattern

在 Golang 滿少看到簡單工廠（因為`原生推薦的寫法就是 simple factory Pattern`），一般來說會透過 NewXxxxx 來建立對應的物件，我們也能夠透過 NewXxxxx 為基底做到各式封裝好的方法。

e.g.
```go
func New(){
    tlsConfig := &tls.Config{InsecureSkipVerify: true}
    return NewWithTLSConfig(tlsConfig, followNonLocalRedirects)
}

func NewWithTLSConfig(config *tls.Config, followNonLocalRedirects bool) Prober {
    // We do not want the probe use node's local proxy set.
    transport := utilnet.SetTransportDefaults(
    &http.Transport{
        TLSClientConfig:   config,
        DisableKeepAlives: true,
        Proxy:             http.ProxyURL(nil),
    })
    return httpProber{transport, followNonLocalRedirects}
}
```

從上述例子來看 New function 提供了已經封裝好的 TLS Config 方法，若是客官不滿意的話也能透過 NewWithTLSConfig function 自己攜帶 TLS Config，如此一來便能達到復用的效果。  

