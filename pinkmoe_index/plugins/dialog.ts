/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 13:55:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 14:01:15
 * @FilePath: /pinkmoe_index/plugins/dialog.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
/*!
 * image lazyload
 */
import Search from '/@/components/Search/index'
import Login from '/@/components/Login/index'
import VipConfirm from '/@/components/Vipconfirm/index'
import ShopConfirm from '/@/components/Shopconfirm/index'
import GlobalNotice from '/@/components/GlobalNotice/index'
import Message from '/@/components/Message/index'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp
    .use(VipConfirm)
    .use(ShopConfirm)
    .use(GlobalNotice)
    .use(Search)
    .use(Login)
    .use(Message)
})
