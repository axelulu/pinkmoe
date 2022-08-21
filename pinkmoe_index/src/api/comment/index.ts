/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:22:29
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:36
 * @FilePath: /pinkmoe_index/src/api/comment/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { ReqComment } from '/@/api/comment/types'
// import axios from 'axios';
enum URL {
  create = '/api/Comment/CommentCreate',
  list = '/api/Comment/CommentTreeList',
  userList = '/api/Comment/CommentList',
}

const getCommentList = async (params: ReqComment) =>
  get<any>({
    url: URL.list,
    params,
  })

const getUserComment = async params =>
  get<any>({
    url: URL.userList,
    params,
  })

const createComment = async data =>
  post<any>(
    {
      url: URL.create,
      data,
    },
    true,
  )

export { getCommentList, getUserComment, createComment }
