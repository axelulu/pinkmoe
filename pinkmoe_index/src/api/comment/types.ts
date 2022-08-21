/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:57:54
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:41:12
 * @FilePath: /pinkmoe_vitesse/src/api/comment/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ReqPage } from '/@/api/common/types'

export interface ReqComment extends ReqPage {
  postId?: string | string[]
}

export interface ReqCommentPost extends ReqPage {
  formUid?: string | string[]
}
