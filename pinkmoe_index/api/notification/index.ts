/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 22:34:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-15 22:33:04
 * @FilePath: /pinkmoe_index/src/api/notification/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ResPage } from '/@/api/common/types'
import type { ResComment } from '../post/types'
// import axios from 'axios';
enum URL {
  list = '/Cms/Notification/NotificationList',
}

const getNotificationList = async params =>
  get<ResPage<Array<ResComment>>>({
    url: URL.list,
    params,
  })
export { getNotificationList }
