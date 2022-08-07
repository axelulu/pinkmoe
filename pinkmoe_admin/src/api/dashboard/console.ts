/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:12
 * @FilePath: /pinkmoe_admin/src/api/dashboard/console.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

//获取主控台信息
export function getConsoleInfo() {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Dashboard/Console',
      method: 'get',
    },
    {
      isTransformResponse: false,
    }
  );
}
