/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:25
 * @FilePath: /pinkmoe_admin/src/api/system/api.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取api列表
 * @param params
 */
export function getApiList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Api/ApiList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 获取所有api
 * @param params
 */
export function getAllApi(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Api/AllApi',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 获取角色api
 * @param params
 */
export function getApiAuthority(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Api/ApiAuthority',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新api
 * @param params
 */
export function updateApi(params) {
  return http.request(
    {
      url: `/Admin/Api/ApiUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新角色api
 * @param params
 */
export function updateApiAuthority(params) {
  return http.request(
    {
      url: `/Admin/Api/ApiAuthorityUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建api
 * @param params
 */
export function createApi(params) {
  return http.request(
    {
      url: `/Admin/Api/ApiCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除api
 * @param params
 */
export function deleteApi(params) {
  return http.request(
    {
      url: `/Admin/Api/ApiDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
