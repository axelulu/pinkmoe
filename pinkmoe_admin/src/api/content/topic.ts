/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:09
 * @FilePath: /pinkmoe_admin/src/api/content/topic.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取话题列表
 * @param params
 */
export function getTopicList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Topic/TopicList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新话题
 * @param params
 */
export function updateTopic(params) {
  return http.request(
    {
      url: `/Admin/Topic/TopicUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建话题
 * @param params
 */
export function createTopic(params) {
  return http.request(
    {
      url: `/Admin/Topic/TopicCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除话题
 * @param params
 */
export function deleteTopic(params) {
  return http.request(
    {
      url: `/Admin/Topic/TopicDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
