/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 15:50:43
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 18:29:57
 * @FilePath: /pinkmoe_vitesse/src/modules/photoPreview.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import vue3PhotoPreview from 'vue3-photo-preview'
import 'vue3-photo-preview/dist/index.css'
import type { UserModule } from '../types'

export const install: UserModule = ({ app }) => {
  app
    .use(vue3PhotoPreview)
}
