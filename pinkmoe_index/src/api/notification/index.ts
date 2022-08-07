/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 22:34:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:54
 * @FilePath: /pinkmoe_index/src/api/notification/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios';
import { ResPage } from '/@/api/common/types';
import { ResTopicPost } from '/@/api/topic/types';
// import axios from 'axios';
enum URL {
  list = '/api/Notification/NotificationList',
}

const getNotificationList = async (params) =>
  get<ResPage<ResTopicPost>>({
    url: URL.list,
    params,
  });
export { getNotificationList };
