/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 17:59:52
 * @FilePath: /xanaduCms/pinkmoe_server/dao/mysql/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package mysql

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Reg(p model.XdUser) (err error, user model.XdUser) {
	if !errors.Is(global.XD_DB.Where("email = ?", p.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return response.ErrorUserEmailExist, user
	}
	// 否则 附加uuid 密码md5简单加密 注册
	password, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return response.ErrorUserCreate, user
	}
	p.Password = string(password)
	p.UUID = uuid.NewV4()
	if err = global.XD_DB.Create(&p).Error; err != nil {
		return response.ErrorUserCreate, user
	}
	return err, user
}

func UpdateUser(p model.XdUser) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if errors.Is(tx.Where("uuid = ?", p.UUID).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return response.ErrorUserNotExist
		}
		AuthorityId := p.AuthorityId
		if err = tx.Delete(&[]model.XdUseAuthority{}, "xd_user_uuid = ?", user.UUID).Error; err != nil {
			return response.ErrorUserUpdate
		}
		if err = tx.Updates(&p).Error; err != nil {
			return response.ErrorUserUpdate
		}
		// 设置用户密码权限
		if err = tx.Where("uuid = ?", p.UUID).First(&model.XdUser{}).Update("authority_id", AuthorityId).Error; err != nil {
			return response.ErrorUserUpdate
		}
		return nil
	}); err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UserCheckInStatusGet(uuid uuid.UUID) (err error, status bool) {
	var checkIn model.XdCheckIn
	now := time.Now()
	hasCheckin := !errors.Is(global.XD_DB.Order("last_checkin_time desc").Where("uuid = ?", uuid).First(&checkIn).Error, gorm.ErrRecordNotFound)
	isToday := checkIn.CreatedAt.Day() == now.Day() && checkIn.CreatedAt.Month() == now.Month() && checkIn.CreatedAt.Year() == now.Year()
	if hasCheckin && isToday {
		return err, true
	}
	return err, false
}

func UpdateUserCheckIn(uuid uuid.UUID, ip string, credit int, checkType string) (err error, crt int) {
	var checkIn []model.XdCheckIn
	now := time.Now()
	isToday := false
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		hasCheckin := !errors.Is(tx.Order("last_checkin_time desc").Where("uuid = ?", uuid).Find(&checkIn).Error, gorm.ErrRecordNotFound)
		if len(checkIn) > 0 {
			isToday = checkIn[0].CreatedAt.Day() == now.Day() && checkIn[0].CreatedAt.Month() == now.Month() && checkIn[0].CreatedAt.Year() == now.Year()
		}
		if hasCheckin && isToday {
			return response.ErrorCheckIned
		}
		if err = tx.Create(&model.XdCheckIn{
			XD_MODEL:        global.XD_MODEL{},
			Uuid:            uuid,
			Credit:          credit,
			CheckType:       checkType,
			Status:          1,
			LastLoginIp:     ip,
			LastCheckinTime: global.XdTime{Time: time.Now()},
			CheckinCount:    len(checkIn),
		}).Error; err != nil {
			return response.ErrorCheckIn
		}
		var user model.XdUser
		if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).First(&user).Error; err != nil {
			return response.ErrorCheckIn
		}
		if checkType == "cash" {
			if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update("cash", user.Cash+credit).Error; err != nil {
				return response.ErrorCheckIn
			}
		} else {
			if err = tx.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update("credit", user.Credit+credit).Error; err != nil {
				return response.ErrorCheckIn
			}
		}
		return err
	}); err != nil {
		return response.ErrorCheckIned, 0
	}
	return err, credit
}

func UpdateUserExp(exp int, uuid uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return response.ErrorUserUpdate
	}
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"exp": user.Exp + exp,
	}).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserCredit(credit int, uuid uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return response.ErrorUserUpdate
	}
	if user.Credit+credit < 0 {
		return response.ErrorUserUpdate
	}
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"credit": user.Credit + credit,
		"exp":    user.Exp + credit*10,
	}).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserAuthority(authorityId string, uuid uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("uuid = ?", uuid).First(&user).Error; err != nil { // 判断用户名是否注册
			return response.ErrorUserUpdate
		}
		if err = tx.Delete(&[]model.XdUseAuthority{}, "xd_user_uuid = ?", user.UUID).Error; err != nil {
			return response.ErrorUserUpdate
		}
		if err = tx.Create(&model.XdUseAuthority{
			XdAuthorityAuthorityId: authorityId,
			XdUserId:               uuid,
		}).Error; err != nil {
			return response.ErrorUserUpdate
		}
		if err = tx.Where("uuid = ?", uuid).First(&model.XdUser{}).Update("authority_id", authorityId).Error; err != nil {
			return response.ErrorUserUpdate
		}
		return err
	}); err != nil {
		return response.ErrorUserUpdate
	}
	return
}

func UpdateUserCash(cash int, uuid uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return response.ErrorUserUpdate
	}
	if user.Cash+cash < 0 {
		return response.ErrorUserUpdate
	}
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"cash": user.Cash + cash,
		"exp":  user.Exp + cash*100,
	}).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserAvatar(avatar string, uuid uuid.UUID) (err error) {
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update("avatar", avatar).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserDetail(nickName string, desc string, uuid uuid.UUID) (err error) {
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"nick_name": nickName,
		"desc":      desc,
	}).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserEmail(email string, uuid uuid.UUID) (err error) {
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update("email", email).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserPwd(password string, oldPassword string, uuid uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return response.ErrorUserUpdate
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return response.ErrorUserUpdate
	}
	// 否则 附加uuid 密码md5简单加密 注册
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return response.ErrorUserCreate
	}
	if err = global.XD_DB.Model(&model.XdUser{}).Where("uuid = ?", uuid).Update("password", pwd).Error; err != nil {
		return response.ErrorUserUpdate
	}
	return err
}

func UpdateUserPwdByEmail(password string, email string) (user model.XdUser, err error) {
	if err = global.XD_DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, response.ErrorUserUpdate
	}
	// 否则 附加uuid 密码md5简单加密 注册
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return user, response.ErrorUserCreate
	}
	if err = global.XD_DB.Model(&model.XdUser{}).Where("email = ?", email).Update("password", pwd).Error; err != nil {
		return user, response.ErrorUserUpdate
	}
	return user, err
}

func DeleteUser(delUUID uuid.UUID, userId uuid.UUID) (err error) {
	var user model.XdUser
	if err = global.XD_DB.Transaction(func(tx *gorm.DB) error {
		if errors.Is(tx.Where("uuid = ?", delUUID).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return response.ErrorUserNotExist
		}
		if userId == user.UUID {
			return response.ErrorUserNotDel
		}
		if err = tx.Where("uuid = ?", delUUID).Delete(&user).Error; err != nil {
			return response.ErrorUserDel
		}
		if err = tx.Delete(&[]model.XdUseAuthority{}, "xd_user_uuid = ?", delUUID).Error; err != nil {
			return response.ErrorUserAuthDel
		}
		return nil
	}); err != nil {
		return response.ErrorUserAuthDel
	}
	return err
}

func Login(p request.Login) (err error, user model.XdUser) {
	if err = global.XD_DB.Where("email = ?", p.EmailCode).Preload("Authorities").Preload("Authority").First(&user).Error; err != nil {
		return response.ErrorUserNotExist, user
	}
	//验证（对比）判读密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password)); err != nil {
		return response.ErrorInvalidPassword, user
	}
	return err, user
}

func GetUserInfoByUserId(userId string) (err error, user model.XdUser) {
	if err = global.XD_DB.Where("uuid = ?", userId).Preload("Authorities").Preload("Authority").First(&user).Error; err != nil {
		return response.ErrorUserNotExist, user
	}
	return err, user
}

func GetUserInfoList(u request.SearchUserParams) (err error, list interface{}, total int64) {
	limit := u.PageSize
	offset := u.PageSize * (u.Page - 1)
	db := global.XD_DB.Model(&model.XdUser{})
	var userList []model.XdUser
	if u.Username != "" {
		db = db.Where("username LIKE ?", "%"+u.Username+"%")
	}

	if u.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+u.NickName+"%")
	}

	if u.Phone != "" {
		db = db.Where("phone = ?", u.Phone)
	}

	if u.Email != "" {
		db = db.Where("email = ?", u.Email)
	}

	if u.Sex != "0" && u.Sex != "" {
		db = db.Where("sex = ?", u.Sex)
	}

	if u.AuthorityId != "" {
		db = db.Where("authority_id = ?", u.AuthorityId)
	}

	if len(u.OrderKey) > 0 {
		if u.Desc {
			db = db.Order(u.OrderKey + " desc")
		} else {
			db = db.Order(u.OrderKey + " asc")
		}
	}

	err = db.Count(&total).Error
	if err != nil {
		return response.ErrorUserMetaGet, nil, 0
	}
	if err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error; err != nil {
		return response.ErrorUserMetaGet, nil, 0
	}
	return err, userList, total
}
