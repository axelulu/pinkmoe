/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:22:29
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:06
 * @FilePath: /pinkmoe_index/src/api/search/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ResPage } from '/@/api/common/types'
import type { ResPost } from '/@/api/home/types'
import type { ReqSearchPost } from '/@/api/search/types'
// import axios from 'axios';
enum URL {
  list = '/Api/Post/SearchPostList',
}

const getSearchPostList = async (params: ReqSearchPost) =>
  get<ResPage<Array<ResPost>>>({
    url: URL.list,
    params,
  })

export { getSearchPostList }
