/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:17
 * @FilePath: /pinkmoe_admin/src/api/monitor/operation.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取操作列表
 * @param params
 */
export function getOperationList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Operation/OperationList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新操作
 * @param params
 */
export function updateOperation(params) {
  return http.request(
    {
      url: `/Admin/Operation/OperationUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建操作
 * @param params
 */
export function createOperation(params) {
  return http.request(
    {
      url: `/Admin/Operation/OperationCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除操作
 * @param params
 */
export function deleteOperation(params) {
  return http.request(
    {
      url: `/Admin/Operation/OperationDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
