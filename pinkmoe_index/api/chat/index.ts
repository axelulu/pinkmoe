/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 16:26:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:33
 * @FilePath: /pinkmoe_index/src/api/chat/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
