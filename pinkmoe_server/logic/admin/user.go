/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:55
 * @FilePath: /pinkmoe_server/logic/admin/user.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"context"
	"encoding/json"
	"math/rand"
	"server/dao/mysql"
	"server/dao/redis"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func PwdForget(p request.PwdForget) (err error) {
	if result, err := global.XD_REDIS.Get(context.Background(), redis.EmailPreKey+p.CaptchaType+"_"+p.EmailCode).Result(); err != nil {
		return response.ErrorValidateCode
	} else {
		if result == p.CaptchaCode {
			var user model.XdUser
			if user, err = mysql.UpdateUserPwdByEmail(p.Password, p.EmailCode); err != nil {
				return err
			}
			// 初始化积分通知
			err, userReward := mysql.GetSettingItem(model.XdSetting{
				Slug: "user_reward",
			})
			var reward response.ResUserReward
			if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
				return response.ErrorUserUpdate
			}
			if reward.ForgetPwdType == "cash" {
				err = mysql.UpdateUserCredit(reward.ForgetPwdCredit, user.UUID)
				if err == nil {
					CreateNotification(user.UUID, "", "", 0, reward.ForgetPwdCredit, 0, "forgetPwd", "")
				}
			} else {
				err = mysql.UpdateUserCash(reward.ForgetPwdCredit, user.UUID)
				if err == nil {
					CreateNotification(user.UUID, "", "", reward.ForgetPwdCredit, 0, 0, "forgetPwd", "")
				}
			}
		}
	}
	return err
}

func Reg(p request.Reg) (user model.XdUser, err error) {
	var u = model.XdUser{
		NickName: p.Username,
		Username: p.Username,
		Password: p.Password,
		Email:    p.EmailCode,
	}
	if result, err := global.XD_REDIS.Get(context.Background(), redis.EmailPreKey+p.CaptchaType+"_"+p.EmailCode).Result(); err != nil {
		return user, response.ErrorValidateCode
	} else {
		if result == p.CaptchaCode {
			if err, user = mysql.Reg(u); err != nil {
				return user, err
			}
			// 初始化积分通知
			err, userReward := mysql.GetSettingItem(model.XdSetting{
				Slug: "user_reward",
			})
			var reward response.ResUserReward
			if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
				return user, response.ErrorUserCreate
			}
			if reward.RegType == "cash" {
				err = mysql.UpdateUserCredit(reward.RegCredit, user.UUID)
				if err == nil {
					CreateNotification(user.UUID, "", "", 0, reward.RegCredit, 0, "reg", "")
				}
			} else {
				err = mysql.UpdateUserCash(reward.RegCredit, user.UUID)
				if err == nil {
					CreateNotification(user.UUID, "", "", reward.RegCredit, 0, 0, "reg", "")
				}
			}
		}
	}
	return user, err
}

func Login(p request.Login, userAgent string, clientIP string) (userMeta map[string]interface{}, err error) {
	userMeta = make(map[string]interface{})
	err, user := mysql.Login(p)
	userMeta["userInfo"] = user
	if err != nil {
		return nil, response.ErrorUserMeta
	}
	token, err := util.GenToken(user.AuthorityId, user.UUID, user.Username)
	if err != nil {
		return nil, response.ErrorUserNotLogin
	}
	userMeta["token"] = token

	// 记录登陆日志
	ua := user_agent.New(userAgent)
	browser, _ := ua.Browser()
	var msg = "登陆后台"
	var status = 1
	if err != nil {
		status = 0
		msg = err.Error()
		return nil, response.ErrorUserNotExist
	}
	if err = mysql.CreateLoginLog(&model.XdLoginLog{
		XD_MODEL:      global.XD_MODEL{},
		UserName:      user.Username,
		Ip:            clientIP,
		LoginLocation: util.GetCityByIp(clientIP),
		Explorer:      browser,
		Os:            ua.OS(),
		Status:        status,
		Msg:           msg,
		Module:        "系统后台",
	}); err != nil {
		return nil, response.ErrorLoginLogCreate
	}

	// 记录用户在线日志
	if err = mysql.CreateUserOnline(&model.XdUserOnline{
		XD_MODEL: global.XD_MODEL{},
		Uuid:     user.UUID,
		Token:    token,
		UserName: user.Username,
		Ip:       clientIP,
		Explorer: browser,
		Os:       ua.OS(),
	}); err != nil {
		return nil, response.ErrorUserOnlineCreate
	}
	return
}

func Logout(userId uuid.UUID, userToken string) (err error) {
	if err = mysql.CreateJwtBlack(userToken); err != nil {
		return response.ErrorUserLogout
	}
	if err = mysql.DeleteUserOnline(userId); err != nil {
		return response.ErrorUserLogout
	}
	return
}

func UserInfo(c *gin.Context) (userInfo map[string]interface{}, err error) {
	userInfo = make(map[string]interface{})
	userId, err := util.GetCurrentUserID(c)
	if err != nil {
		return nil, response.ErrorUserNotLogin
	} else {
		var user model.XdUser
		var menu []model.XdMenu
		err, user = mysql.GetUserInfoByUserId(userId.String())
		if err != nil {
			return
		}
		userInfo["userInfo"] = user
		err, menu = mysql.GetMenuTree(user.AuthorityId)
		if err != nil {
			return
		}
		userInfo["userMenu"] = menu
	}
	return
}

func GetUserInfoList(p request.SearchUserParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetUserInfoList(p)
	return err, list, total
}

func GetUserInfo(p request.SearchUserParams) (err error, list model.XdUser) {
	err, list = mysql.GetUserInfoByUserId(p.UUID)
	return err, list
}

func UserCheckInStatusGet(uuid uuid.UUID) (err error, status bool) {
	err, status = mysql.UserCheckInStatusGet(uuid)
	return err, status
}

func UpdateUserCheckIn(uuid uuid.UUID, ip string) (err error, credit int) {
	rand.Seed(time.Now().UnixNano())
	// 初始化积分通知
	err, userReward := mysql.GetSettingItem(model.XdSetting{
		Slug: "user_reward",
	})
	var reward response.ResUserReward
	if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
		return response.ErrorCheckIn, 0
	}
	credit = rand.Intn(reward.CheckInCreditFoot-reward.CheckInCreditHead) + reward.CheckInCreditHead
	// 随机获取20-100积分
	err, credit = mysql.UpdateUserCheckIn(uuid, ip, credit)
	if err != nil {
		return err, 0
	}
	if reward.CheckInType == "cash" {
		err = mysql.UpdateUserCredit(credit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", 0, credit, 0, "checkIn", "")
		}
	} else {
		err = mysql.UpdateUserCash(credit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", credit, 0, 0, "checkIn", "")
		}
	}
	return err, credit
}

func UpdateUserAvatar(avatar string, uuid uuid.UUID) (err error) {
	if err = mysql.UpdateUserAvatar(avatar, uuid); err != nil {
		return err
	}
	// 初始化积分通知
	err, userReward := mysql.GetSettingItem(model.XdSetting{
		Slug: "user_reward",
	})
	var reward response.ResUserReward
	if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
		return response.ErrorUserUpdate
	}
	if reward.UpdateAvatarType == "cash" {
		err = mysql.UpdateUserCredit(reward.UpdateAvatarCredit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", 0, reward.UpdateAvatarCredit, 0, "updateAvatar", "")
		}
	} else {
		err = mysql.UpdateUserCash(reward.UpdateAvatarCredit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", reward.UpdateAvatarCredit, 0, 0, "updateAvatar", "")
		}
	}
	return err
}

func UserBuyShop(p request.UserShop, uuid uuid.UUID) (err error) {
	switch p.ShopType {
	case "credit":
		if p.PriceType == "cash" {
			if err = mysql.UpdateUserCash(-p.ShopNum*p.ShopCredit, uuid); err != nil {
				return response.ErrorCash
			}
			if err = mysql.UpdateUserCredit(p.ShopNum*p.ShopValue, uuid); err != nil {
				return response.ErrorBuyCredit
			}
			CreateNotification(uuid, "", "", p.ShopNum*p.ShopValue, p.ShopNum*p.ShopCredit, 0, "shopCredit_cash", "")
		} else if p.PriceType == "key" {
			if err, key := mysql.GetKeyItem(&model.XdKey{
				Key:    p.ShopKey,
				Value:  p.ShopValue,
				Type:   p.ShopType,
				Status: "1",
			}); err != nil {
				return response.ErrorKeyUsed
			} else {
				if err := mysql.UpdateKeyItem(&model.XdKey{
					XD_MODEL: global.XD_MODEL{
						ID: key.ID,
					},
					Status: "2",
				}); err != nil {
					return response.ErrorKeyUsed
				}
				if err = mysql.UpdateUserCredit(p.ShopValue, uuid); err != nil {
					return response.ErrorBuyCredit
				}
				CreateNotification(uuid, "", "", 0, p.ShopValue, 0, "shopCredit_key", "")
			}
		} else {
			return response.ErrorBuyCredit
		}
	case "cash":
		if err, key := mysql.GetKeyItem(&model.XdKey{
			Key:    p.ShopKey,
			Value:  p.ShopValue,
			Type:   p.ShopType,
			Status: "1",
		}); err != nil {
			return response.ErrorKeyUsed
		} else {
			if key.Value == p.ShopValue && key.Type == p.ShopType {
				if err := mysql.UpdateKeyItem(&model.XdKey{
					XD_MODEL: global.XD_MODEL{
						ID: key.ID,
					},
					Status: "2",
				}); err != nil {
					return response.ErrorKeyUsed
				}
				if err = mysql.UpdateUserCash(p.ShopValue, uuid); err != nil {
					return response.ErrorBuyCash
				}
				CreateNotification(uuid, "", "", 0, p.ShopValue, 0, "shopCash_key", "")
			} else {
				return response.ErrorPayChoose
			}
		}
	}
	return
}

func UserBuyVip(p request.XdAuthorityParams, uuid uuid.UUID) (err error) {
	err, user := mysql.GetUserInfoByUserId(uuid.String())
	if err != nil {
		return response.ErrorBuyVip
	}
	err, authority := mysql.GetAuthorityItem(p.AuthorityId)
	if err != nil {
		return response.ErrorBuyVip
	}
	println(user.Authority.AuthorityWeight)
	println(authority.AuthorityWeight)
	if user.Authority.AuthorityWeight <= authority.AuthorityWeight {
		return response.ErrorUserLevel
	}
	if err = mysql.UpdateUserCash(-(p.AuthorityParamValue * p.AuthorityNum), uuid); err != nil {
		return response.ErrorCashTakeOut
	}
	if err = mysql.UpdateUserAuthority(p.AuthorityId, uuid); err != nil {
		return response.ErrorBuyVip
	}
	// 设置定时器
	err, task := mysql.GetTask(&model.XdTask{
		Slug: p.AuthorityParamId,
		Type: "vip",
		Uuid: uuid,
	})
	if err != nil {
		mm, _ := time.ParseDuration(strconv.Itoa(p.AuthorityParamDay*24) + "h")
		xdTime := global.XdTime{Time: time.Now().Add(mm)}
		if err = mysql.CreateTask(&model.XdTask{
			Name:    "用户会员过期",
			Slug:    p.AuthorityParamId,
			Type:    "vip",
			Uuid:    uuid,
			TimeOut: xdTime,
		}); err != nil {
			return response.ErrorBuyVip
		}
	} else {
		mm, _ := time.ParseDuration(strconv.Itoa(p.AuthorityParamDay*24) + "h")
		xdTime := global.XdTime{Time: task.TimeOut.Add(mm)}
		if err := mysql.DeleteTask(&task); err != nil {
			return response.ErrorBuyVip
		}
		if err = mysql.CreateTask(&model.XdTask{
			Name:    "用户会员过期",
			Slug:    p.AuthorityParamId,
			Type:    "vip",
			Uuid:    uuid,
			TimeOut: xdTime,
		}); err != nil {
			return response.ErrorBuyVip
		}
	}
	return err
}

func UpdateUserDetail(nickName string, desc string, uuid uuid.UUID) (err error) {
	err = mysql.UpdateUserDetail(nickName, desc, uuid)
	return err
}

func UpdateUserEmail(email string, emailCaptcha string, uuid uuid.UUID) (err error) {
	if result, err := global.XD_REDIS.Get(context.Background(), redis.EmailPreKey+"changePwd_"+email).Result(); err != nil {
		return response.ErrorValidateCode
	} else {
		if result == emailCaptcha {
			if err = mysql.UpdateUserEmail(email, uuid); err != nil {
				return err
			}
			// 初始化积分通知
			err, userReward := mysql.GetSettingItem(model.XdSetting{
				Slug: "user_reward",
			})
			var reward response.ResUserReward
			if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
				return response.ErrorUserUpdate
			}
			if reward.UpdateEmailType == "cash" {
				err = mysql.UpdateUserCredit(reward.UpdateEmailCredit, uuid)
				if err == nil {
					CreateNotification(uuid, "", "", 0, reward.UpdateEmailCredit, 0, "updateEmail", "")
				}
			} else {
				err = mysql.UpdateUserCash(reward.UpdateEmailCredit, uuid)
				if err == nil {
					CreateNotification(uuid, "", "", reward.UpdateEmailCredit, 0, 0, "updateEmail", "")
				}
			}
		}
	}
	return err
}

func UpdateUserPwd(password string, oldPassword string, uuid uuid.UUID) (err error) {
	if err = mysql.UpdateUserPwd(password, oldPassword, uuid); err != nil {
		return err
	}
	// 初始化积分通知
	err, userReward := mysql.GetSettingItem(model.XdSetting{
		Slug: "user_reward",
	})
	var reward response.ResUserReward
	if err = json.Unmarshal([]byte(userReward.Value), &reward); err != nil {
		return response.ErrorUserUpdate
	}
	if reward.UpdatePwdType == "cash" {
		err = mysql.UpdateUserCredit(reward.UpdatePwdCredit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", 0, reward.UpdatePwdCredit, 0, "updatePwd", "")
		}
	} else {
		err = mysql.UpdateUserCash(reward.UpdatePwdCredit, uuid)
		if err == nil {
			CreateNotification(uuid, "", "", reward.UpdatePwdCredit, 0, 0, "updatePwd", "")
		}
	}
	return err
}

func CreateUser(p model.XdUser) (user model.XdUser, err error) {
	var authorities []model.XdAuthority
	authorities = append(authorities, model.XdAuthority{
		AuthorityId: p.AuthorityId,
	})
	p.Authorities = authorities
	if err, user = mysql.Reg(p); err != nil {
		return
	}
	return
}

func UpdateUser(p request.ReqUser) (err error) {
	// 密码加密 注册
	password, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return response.ErrorUserUpdate
	}
	var authorities []model.XdAuthority
	authorities = append(authorities, model.XdAuthority{
		AuthorityId: p.AuthorityId,
	})
	uuid, _ := uuid.FromString(p.UUID)
	u := model.XdUser{
		UUID:        uuid,
		Username:    p.Username,
		Password:    p.Password,
		NickName:    p.NickName,
		Avatar:      p.Avatar,
		HeaderImg:   p.HeaderImg,
		Sex:         p.Sex,
		Cash:        p.Cash,
		Credit:      p.Credit,
		Exp:         p.Exp,
		Phone:       p.Phone,
		Email:       p.Email,
		AuthorityId: p.AuthorityId,
		Authority:   p.Authority,
		Authorities: authorities,
	}
	p.Password = string(password)
	err = mysql.UpdateUser(u)
	return err
}

func DeleteUser(userId uuid.UUID, p request.GetByUUId) (err error) {
	err = mysql.DeleteUser(p.UUID, userId)
	return
}
