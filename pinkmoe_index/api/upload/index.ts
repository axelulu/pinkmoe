/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 08:41:19
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:24
 * @FilePath: /pinkmoe_index/src/api/upload/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { ReqFileList, ResFileList } from '/@/api/upload/types'
import type { IResponse } from '/@/utils/http/axios/type'
// import axios from 'axios';
enum URL {
  upload = '/Api/Upload/FileUpload',
  list = '/Api/Upload/FileList',
  delete = '/Api/Upload/ImgFileDelete',
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
