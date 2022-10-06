/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:54:46
 * @FilePath: /pinkmoe_server/model/response/xd_setting.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
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
