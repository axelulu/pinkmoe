/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:46
 * @FilePath: /pinkmoe_server/model/response/xd_setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package response

type ResUserReward struct {
	PublishPostType    string `json:"publishPostType" form:"publishPostType"`       // 话题标识
	PublishPostCredit  int    `json:"publishPostCredit" form:"publishPostCredit"`   // 话题标识
	PublishGoodsType   string `json:"publishGoodsType" form:"publishGoodsType"`     // 话题标识
	PublishGoodsCredit int    `json:"publishGoodsCredit" form:"publishGoodsCredit"` // 话题标识
	CommentType        string `json:"commentType" form:"commentType"`               // 话题标识
	CommentCredit      int    `json:"commentCredit" form:"commentCredit"`           // 话题标识
	RegType            string `json:"regType" form:"regType"`                       // 话题标识
	RegCredit          int    `json:"regCredit" form:"regCredit"`                   // 话题标识
	UpdatePwdType      string `json:"updatePwdType" form:"updatePwdType"`           // 话题标识
	UpdatePwdCredit    int    `json:"updatePwdCredit" form:"updatePwdCredit"`       // 话题标识
	UpdateEmailType    string `json:"updateEmailType" form:"updateEmailType"`       // 话题标识
	UpdateEmailCredit  int    `json:"updateEmailCredit" form:"updateEmailCredit"`   // 话题标识
	UpdateAvatarType   string `json:"updateAvatarType" form:"updateAvatarType"`     // 话题标识
	UpdateAvatarCredit int    `json:"updateAvatarCredit" form:"updateAvatarCredit"` // 话题标识
	CheckInType        string `json:"checkInType" form:"checkInType"`               // 话题标识
	CheckInCreditHead  int    `json:"checkInCreditHead" form:"checkInCreditHead"`   // 话题标识
	CheckInCreditFoot  int    `json:"checkInCreditFoot" form:"checkInCreditFoot"`   // 话题标识
	ForgetPwdType      string `json:"forgetPwdType" form:"forgetPwdType"`           // 话题标识
	ForgetPwdCredit    int    `json:"forgetPwdCredit" form:"forgetPwdCredit"`       // 话题标识
}
