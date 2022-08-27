/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 21:16:47
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:13
 * @FilePath: /pinkmoe_index/src/api/setting/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ResSiteSetting } from '/@/api/setting/types'
// import axios from 'axios';
enum URL {
  list = '/Cms/Setting/SiteSetting',
}

const getSiteSetting = async () =>
  get<ResSiteSetting>({
    url: URL.list,
  })
export { getSiteSetting }
