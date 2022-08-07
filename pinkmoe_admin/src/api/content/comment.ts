/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 14:55:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:04
 * @FilePath: /pinkmoe_admin/src/api/content/comment.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取评论列表
 * @param params
 */
export function getCommentList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Comment/CommentList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新评论
 * @param params
 */
export function updateComment(params) {
  return http.request(
    {
      url: `/Admin/Comment/CommentUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建评论
 * @param params
 */
export function createComment(params) {
  return http.request(
    {
      url: `/Admin/Comment/CommentCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除评论
 * @param params
 */
export function deleteComment(params) {
  return http.request(
    {
      url: `/Admin/Comment/CommentDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
