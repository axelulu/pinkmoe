/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:02
 * @FilePath: /pinkmoe_server/dao/mysql/loginLog.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
)

func GetLoginLogList(info *request.SearchLoginLogParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdLoginLog{})
	if info.Ip != "" {
		db = db.Where("ip = ?", info.Ip)
	}

	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}

	if info.Explorer != "" {
		db = db.Where("explorer LIKE ?", "%"+info.Explorer+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorLoginLogListGet, nil, 0
	}
	var comment []model.XdLoginLog
	if err = db.Limit(limit).Offset(offset).Find(&comment).Error; err != nil {
		return response.ErrorLoginLogListGet, nil, 0
	}
	return err, comment, total
}

func GetLoginLogByTime(days int) (total int64) {
	db := global.XD_DB.Model(&model.XdLoginLog{})

	db = db.Where("to_days(now()) - to_days(updated_at) <= ?", days)

	if err := db.Count(&total).Error; err != nil {
		return 0
	}
	return total
}

func GetLoginLogTrend() (today []int, yesterday []int) {
	db := global.XD_DB.Model(&model.XdLoginLog{})

	db = db.Where("to_days(now()) - to_days(updated_at) <= ?", 1)

	var logs []model.XdLoginLog
	if err := db.Find(&logs).Error; err != nil {
		return today, yesterday
	}

	db2 := global.XD_DB.Model(&model.XdLoginLog{})

	db2 = db2.Where("to_days(now()) - to_days(updated_at) <= ?", 1)

	var logs2 []model.XdLoginLog
	if err := db2.Find(&logs).Error; err != nil {
		return today, yesterday
	}
	today = append(today, getTrendNum(0, 6, logs))
	today = append(today, getTrendNum(6, 7, logs))
	today = append(today, getTrendNum(7, 8, logs))
	today = append(today, getTrendNum(8, 9, logs))
	today = append(today, getTrendNum(9, 10, logs))
	today = append(today, getTrendNum(10, 11, logs))
	today = append(today, getTrendNum(11, 12, logs))
	today = append(today, getTrendNum(12, 13, logs))
	today = append(today, getTrendNum(13, 14, logs))
	today = append(today, getTrendNum(14, 15, logs))
	today = append(today, getTrendNum(15, 16, logs))
	today = append(today, getTrendNum(16, 17, logs))
	today = append(today, getTrendNum(17, 18, logs))
	today = append(today, getTrendNum(18, 19, logs))
	today = append(today, getTrendNum(19, 20, logs))
	today = append(today, getTrendNum(20, 21, logs))
	today = append(today, getTrendNum(21, 22, logs))
	today = append(today, getTrendNum(22, 23, logs))

	yesterday = append(yesterday, getTrendNum(0, 6, logs2))
	yesterday = append(yesterday, getTrendNum(6, 7, logs2))
	yesterday = append(yesterday, getTrendNum(7, 8, logs2))
	yesterday = append(yesterday, getTrendNum(8, 9, logs2))
	yesterday = append(yesterday, getTrendNum(9, 10, logs2))
	yesterday = append(yesterday, getTrendNum(10, 11, logs2))
	yesterday = append(yesterday, getTrendNum(11, 12, logs2))
	yesterday = append(yesterday, getTrendNum(12, 13, logs2))
	yesterday = append(yesterday, getTrendNum(13, 14, logs2))
	yesterday = append(yesterday, getTrendNum(14, 15, logs2))
	yesterday = append(yesterday, getTrendNum(15, 16, logs2))
	yesterday = append(yesterday, getTrendNum(16, 17, logs2))
	yesterday = append(yesterday, getTrendNum(17, 18, logs2))
	yesterday = append(yesterday, getTrendNum(18, 19, logs2))
	yesterday = append(yesterday, getTrendNum(19, 20, logs2))
	yesterday = append(yesterday, getTrendNum(20, 21, logs2))
	yesterday = append(yesterday, getTrendNum(21, 22, logs2))
	yesterday = append(yesterday, getTrendNum(22, 23, logs2))
	return today, yesterday
}

func getTrendNum(start int, end int, logs []model.XdLoginLog) (num int) {
	num = 0
	for _, log := range logs {
		if log.UpdatedAt.Hour() > start && log.UpdatedAt.Hour() <= end {
			num++
		}
	}
	return num
}

func CreateLoginLog(p *model.XdLoginLog) (err error) {
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorLoginLogCreate
	}
	return err
}

func UpdateLoginLog(u *model.XdLoginLog) (err error) {
	if err = global.XD_DB.Where("id = ?", u.ID).First(&model.XdLoginLog{}).Updates(&model.XdLoginLog{
		UserName:      u.UserName,
		Ip:            u.Ip,
		LoginLocation: u.LoginLocation,
		Explorer:      u.Explorer,
		Os:            u.Os,
		Status:        u.Status,
		Msg:           u.Msg,
		Module:        u.Module,
	}).Error; err != nil {
		return response.ErrorLoginLogUpdate
	}
	return err
}

func DeleteLoginLog(p *model.XdLoginLog) (err error) {
	if err = global.XD_DB.Delete(&p).Error; err != nil {
		return response.ErrorLoginLogDelete
	}
	return
}
