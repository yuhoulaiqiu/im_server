package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"regexp"
	"strings"
	"yim_server/common/etcd"
)

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string `json:"addr"`
	Etcd string `json:"etcd"`
	Log  logx.LogConf
}

var config Config

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		UserID int `json:"userId"`
		Role   int `json:"role"`
	} `json:"data"`
}

// 网关函数
func gateway(res http.ResponseWriter, req *http.Request) {
	//匹配请求前缀 /api/user/xxx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) < 2 {
		res.Write([]byte("invalid request"))
		return
	}
	service := addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		logx.Errorf(" %s 不匹配的服务", service)
		res.Write([]byte("invalid request"))
		return
	}
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	logx.Infof("请求服务地址:%s, 客户端地址:%s", addr, remoteAddr[0])
	// 创建一个缓冲区并将req.Body的内容复制到其中
	bodyBytes, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 把原始的req.Body内容放回去

	// 请求认证服务地址
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	authReq, err := http.NewRequest("POST", authUrl, bytes.NewReader(bodyBytes)) // 使用缓冲区的内容创建新的请求
	authReq.Header = req.Header
	authReq.Header.Set("ValidPath", req.URL.Path)
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Error(err)
		res.Write([]byte("认证服务异常"))
		return
	}

	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	err = json.Unmarshal(byteData, &authResponse)
	if err != nil {
		logx.Error(err)
		res.Write([]byte("认证服务异常"))
		return
	}
	if authResponse.Code != 0 {
		res.Write(byteData)
		return
	}
	// 设置请求头,返回认证数据
	if authResponse.Data != nil {
		req.Header.Set("User-ID", fmt.Sprintf("%d", authResponse.Data.UserID))
		req.Header.Set("Role", fmt.Sprintf("%d", authResponse.Data.Role))
	}

	url := "http://" + addr + req.URL.Path
	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(bodyBytes)) // 使用缓冲区的内容创建新的请求
	if err != nil {
		logx.Errorf("创建请求失败:%v", err)
		return
	}
	for name, headers := range req.Header {
		for _, h := range headers {
			proxyReq.Header.Add(name, h)
		}
	}
	fmt.Println(req.Method, url)
	authReq.Header.Set("Content-Type", "application/json")
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	response, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("服务异常"))
		return
	}

	io.Copy(res, response.Body)
}

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config)
	logx.SetUp(config.Log)
	//回调函数
	http.HandleFunc("/", gateway)
	fmt.Printf("gateway server running at %s\n", config.Addr)
	//绑定服务
	err := http.ListenAndServe(config.Addr, nil)
	if err != nil {
		return
	}
}
