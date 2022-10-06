/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 16:26:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:33
 * @FilePath: /pinkmoe_index/src/api/chat/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { IResponse } from '/@/utils/http/axios/type'
// import axios from 'axios';
enum URL {
  list = '/Cms/Chat/ChatList',
  listRelation = '/Cms/Chat/ChatRelationList',
  createRelation = '/Cms/Chat/CreateChatRelation',
  deleteRelation = '/Cms/Chat/DeleteChatRelation',
}

const getChatList = async (params: any) =>
  get<any>({
    url: URL.list,
    params,
  })

const getChatRelationList = async (params: any) =>
  get<any>({
    url: URL.listRelation,
    params,
  })

const createChatRelationList = async (data: any) =>
  post<IResponse<any>>(
    {
      url: URL.createRelation,
      data,
    },
    true,
  )

const deleteChatRelationList = async (data: any) =>
  post<IResponse<any>>(
    {
      url: URL.deleteRelation,
      data,
    },
    true,
  )

export { getChatList, deleteChatRelationList, getChatRelationList, createChatRelationList }
