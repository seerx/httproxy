package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/seerx/base/pkg/configure"
	"github.com/seerx/httproxy/pkg/config"
	"github.com/seerx/httproxy/pkg/xval"
)

func newReverseProxy(proxyMap map[string]*config.ProxyMap) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(request *http.Request) {
			// fmt.Println("host", request.Host)
			p, ok := proxyMap[request.Host]
			if ok {
				request.URL.Scheme = p.Scheme
				request.URL.Host = p.Target
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