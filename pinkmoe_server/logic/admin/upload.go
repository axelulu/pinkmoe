/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 08:58:52
 * @FilePath: /pinkmoe_server/logic/admin/upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
package adminLogic

import (
	"mime/multipart"
	"server/dao/mysql"
	"server/model"
	"server/model/request"
)

func UploadFile(header *multipart.FileHeader, noSave string, postId string, goodsId string, uuid string, uploadType string) (err error, file model.XdUploadFile) {
	err, file = mysql.UploadFile(header, noSave, postId, goodsId, uuid, uploadType)
	return
}

func GetFileList(p request.XdUploadFileParams) (err error, list interface{}, total int64) {
	err, list, total = mysql.GetFileList(p)
	return
}

func DeleteImgFile(p request.DeleteImgFileParams) (err error) {
	err = mysql.DeleteImgFile(p)
	return
}
