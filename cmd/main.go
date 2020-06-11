package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/seerx/base/pkg/configure"
	"github.com/seerx/httproxy/pkg/block"
	"github.com/seerx/httproxy/pkg/config"
	"github.com/seerx/httproxy/pkg/xval"
)

var invalideHostMap = map[string]int{}
var mapLocker sync.Mutex

const invalideMaxTimes = 3

func invalide(remoteAddr string) {
	idx := strings.Index(remoteAddr, ":")
	if idx > 0 {
		remoteAddr = remoteAddr[:idx]
	}
	go func() {
		mapLocker.Lock()
		defer mapLocker.Unlock()
		c, ok := invalideHostMap[remoteAddr]
		if !ok {
			invalideHostMap[remoteAddr] = 1
		} else if c+1 >= invalideMaxTimes {
			delete(invalideHostMap, remoteAddr)
			block.RejectIP(remoteAddr)
		} else {
			invalideHostMap[remoteAddr] = c + 1
		}
	}()
}

func newReverseProxy(proxyMap map[string]*config.ProxyMap) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(request *http.Request) {
			// fmt.Println("host", request.Host)
			p, ok := proxyMap[request.Host]
			if ok {
				request.URL.Scheme = p.Scheme
				request.URL.Host = p.Target
			} else {
				// request.RemoteAddr
				invalide(request.RemoteAddr)
			}
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(fmt.Sprintf("[%s] Error %s!", r.Host, err.Error())))
			}
		},
	}
}

func main() {
	// 读取配置文件
	data := configure.GetConfigureData("f", "e")
	cfg, err := config.Parse(data)
	if err != nil {
		log.Fatal(err)
	}
	// 启动主站
	go xval.Start(cfg.Home.Port)

	// 解析反向代理表
	pmap := map[string]*config.ProxyMap{}
	for _, p := range cfg.ProxyMaps {
		pmap[p.Host] = p
	}
	http.ListenAndServe(":80", newReverseProxy(pmap))
}
