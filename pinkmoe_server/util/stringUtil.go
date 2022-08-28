/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:51:07
 * @FilePath: /pinkmoe_server/util/stringUtil.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */

package util

import (
	"encoding/hex"
	"math/rand"
	"server/global"
	"sort"
	"time"
)

// Reverse 方式二: SliceStable
func Reverse(s interface{}) {
	sort.SliceStable(s, func(i, j int) bool {
		return true
	})
}

func GetPageCount(total int64, PageSize int) (tt int64) {
	if PageSize == 0 {
		return 1
	}
	if int(total)%PageSize == 0 {
		return int64(int(total) / PageSize)
	} else {
		return int64(int(total)/PageSize) + 1
	}
}

// RandUp 随机字符串，包含 英文字母和数字附加=_两个符号
func RandUp(n int) []byte {
	rand.Seed(time.Now().Unix())
	var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=_")
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 63
		b[i] = longLetters[arc]
	}
	s1 := global.XD_CONFIG.BasicConfig.Name + "_"
	return append([]byte(s1), b...)
}

// RandHex 生成16进制格式的随机字符串
func RandHex(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var need int
	if n&1 == 0 { // even
		need = n
	} else { // odd
		need = n + 1
	}
	size := need / 2
	dst := make([]byte, need)
	src := dst[size:]
	if _, err := rand.Read(src[:]); err != nil {
		return []byte{}
	}
	hex.Encode(dst, src)
	return dst[:n]
}
