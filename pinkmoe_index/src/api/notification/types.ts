/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 22:34:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:32
 * @FilePath: /pinkmoe_vitesse/src/api/notification/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { BaseResponseType, ReqPage } from '/@/api/common/types'
import type { ResPost } from '/@/api/home/types'

export interface ReqTopicPost extends ReqPage {
  value?: string | string[]
}

export interface ResTopicPost extends BaseResponseType {
  value?: string
  sort?: number
  label?: string
  icon?: string
  ID?: number
  post?: Array<ResPost>
}
