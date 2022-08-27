/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:22:29
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 10:08:47
 * @FilePath: /pinkmoe_index/src/api/goods/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ReqGoodsItem, ResGoodsItem } from './types'
// import axios from 'axios';
enum URL {
  item = '/Cms/Goods/GoodsItem',
}

const getGoodsItem = async (params: ReqGoodsItem) =>
  get<ResGoodsItem>({
    url: URL.item,
    params,
  })

export { getGoodsItem }
