/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 21:16:47
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-17 21:21:22
 * @FilePath: /pinkmoe_index/src/api/shop/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ResShopList } from '/@/api/shop/types'
// import axios from 'axios';
enum URL {
  list = '/Api/Shop/ShopList',
}

const getShopList = async () =>
  get<ResShopList>({
    url: URL.list,
    params: {
      page: 1,
      pageSize: 999,
    },
  })
export { getShopList }
