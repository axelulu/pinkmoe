/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 08:41:19
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:24
 * @FilePath: /pinkmoe_index/src/api/upload/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { ReqFileList, ResFileList } from '/@/api/upload/types'
import type { IResponse } from '/@/utils/http/axios/type'
// import axios from 'axios';
enum URL {
  upload = '/Cms/Upload/FileUpload',
  list = '/Cms/Upload/FileList',
  delete = '/Cms/Upload/ImgFileDelete',
}

const fileList = async (params: ReqFileList) =>
  get<any>({
    url: URL.list,
    params,
  })

const upload = async (fd: FormData) =>
  post<IResponse<ResFileList>>(
    {
      url: URL.upload,
      data: fd,
    },
    true,
  )

const deleteFile = async data =>
  post<IResponse>(
    {
      url: URL.delete,
      data,
    },
    true,
  )

export { upload, fileList, deleteFile }
