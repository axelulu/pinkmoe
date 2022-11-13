/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-22 11:13:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:01:31
 * @FilePath: /pinkmoe_server/dao/mysql/upload.go
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
package mysql

import (
	"mime/multipart"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/util/upload"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func Upload(file model.XdUploadFile) (err error) {
	if err = global.XD_DB.Create(&file).Error; err != nil {
		return response.ErrorFileUpload
	}
	return
}

func GetFileList(info request.XdUploadFileParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.XD_DB.Model(&model.XdUploadFile{})

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}

	if info.PostId != "" {
		db = db.Where("post_id = ?", info.PostId)
	}

	if info.Uuid.String() != "" {
		db = db.Where("uuid = ?", info.Uuid)
	}

	if err = db.Count(&total).Error; err != nil {
		return response.ErrorFileListGet, nil, 0
	}

	var file []model.XdUploadFile
	if err = db.Limit(limit).Offset(offset).Find(&file).Error; err != nil {
		return response.ErrorFileListGet, nil, 0
	}
	return err, file, total
}

func UploadFile(header *multipart.FileHeader, noSave string, postId string, goodsId string, uuids string, uploadType string) (err error, file model.XdUploadFile) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	println(filePath)
	uuidd, err := uuid.FromString(uuids)
	if err != nil {
		return err, model.XdUploadFile{}
	}
	if uploadErr != nil {
		return response.ErrorFileUpload, file
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		if len(postId) >= 19 {
			f := model.XdUploadFile{
				Url:    filePath,
				Name:   header.Filename,
				PostId: postId,
				Tag:    s[len(s)-1],
				Key:    key,
				Uuid:   uuidd,
				Type:   uploadType,
			}
			return Upload(f), f
		} else if len(goodsId) >= 19 {
			f := model.XdUploadFile{
				Url:     filePath,
				Name:    header.Filename,
				GoodsId: goodsId,
				Tag:     s[len(s)-1],
				Key:     key,
				Uuid:    uuidd,
				Type:    uploadType,
			}
			return Upload(f), f
		} else {
			f := model.XdUploadFile{
				Url:  filePath,
				Name: header.Filename,
				Tag:  s[len(s)-1],
				Key:  key,
				Uuid: uuidd,
				Type: uploadType,
			}
			return Upload(f), f
		}
	}
	return
}

func DeleteImgFile(p request.DeleteImgFileParams) (err error) {
	var File model.XdUploadFile
	if err = global.XD_DB.Where("url = ?", p.Url).Delete(&File).Error; err != nil {
		return response.ErrorFileDelete
	}
	return
}
