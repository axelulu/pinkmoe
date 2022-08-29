/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:25
 * @FilePath: /pinkmoe_server/util/ip.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	body, _ := ioutil.ReadAll(r.Body)
	tmp, _ := GbkToUtf8(body)
	var ipJson map[string]interface{}
	_ = json.Unmarshal(tmp, &ipJson)
	if err != nil {
		return ""
	}
	if ipJson["code"] == 0 {
		city := fmt.Sprintf("%s %s", ipJson["pro"], ipJson["city"])
		return city
	} else {
		return ""
	}
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}
