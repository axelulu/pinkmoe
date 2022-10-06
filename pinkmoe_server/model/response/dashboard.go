/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:15
 * @FilePath: /pinkmoe_server/model/response/dashboard.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package response

type OrderLarge struct {
	Amount    int `json:"amount"`
	Decline   int `json:"decline"`
	Rise      int `json:"rise"`
	WeekLarge int `json:"weekLarge"`
}

type Saleroom struct {
	Amount       int     `json:"amount"`
	Degree       float64 `json:"degree"`
	WeekSaleroom int     `json:"weekSaleroom"`
}

type Visits struct {
	Amount    int     `json:"amount"`
	DayVisits int     `json:"dayVisits"`
	Decline   float64 `json:"decline"`
	Rise      float64 `json:"rise"`
}

type Ua struct {
	Amount   int     `json:"amount"`
	Decline  float64 `json:"decline"`
	Rise     float64 `json:"rise"`
	DayLarge int     `json:"dayLarge"`
}

type FluxTrend struct {
	Today     []int `json:"today"`
	Yesterday []int `json:"yesterday"`
}

type Console struct {
	OrderLarge `json:"orderLarge"`
	Saleroom   `json:"saleroom"`
	Visits     `json:"visits"`
	Ua         `json:"ua"`
	FluxTrend  `json:"fluxTrend"`
}
