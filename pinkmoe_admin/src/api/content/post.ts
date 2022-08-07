/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-08 01:11:49
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:07
 * @FilePath: /pinkmoe_admin/src/api/content/post.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取文章列表
 * @param params
 */
export function getPostList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Post/PostList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 上传文件
 * @param params
 */
export function upload(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Upload/FileUpload',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新文章
 * @param params
 */
export function updatePost(params) {
  return http.request(
    {
      url: `/Admin/Post/PostUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新文章
 * @param params
 */
export function updatePostStatus(params) {
  return http.request(
    {
      url: `/Admin/Post/PostUpdateStatus`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建文章
 * @param params
 */
export function createPost(params) {
  return http.request(
    {
      url: `/Admin/Post/PostCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除文章
 * @param params
 */
export function deletePost(params) {
  return http.request(
    {
      url: `/Admin/Post/PostDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
