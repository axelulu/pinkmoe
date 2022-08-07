/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:23
 * @FilePath: /pinkmoe_admin/src/api/setting/setting.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取设置
 * @param params
 */
export function getSetting(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Setting/SettingItem',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新设置
 * @param params
 */
export function updateSetting(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Setting/SettingUpdate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建设置
 * @param params
 */
export function createSetting(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Setting/SettingCreate',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
