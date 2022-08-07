/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:39
 * @FilePath: /pinkmoe_admin/src/api/userCenter/key.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取卡密列表
 * @param params
 */
export function getKeyList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Key/KeyList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新卡密
 * @param params
 */
export function updateKey(params) {
  return http.request(
    {
      url: `/Admin/Key/KeyUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建卡密
 * @param params
 */
export function createKey(params) {
  return http.request(
    {
      url: `/Admin/Key/KeyCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除卡密
 * @param params
 */
export function deleteKey(params) {
  return http.request(
    {
      url: `/Admin/Key/KeyDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
