/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-08 01:11:49
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-16 22:15:50
 * @FilePath: /pinkmoe_admin/src/api/shop/goods.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取商品列表
 * @param params
 */
export function getGoodsList(params?) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/Goods/GoodsList',
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
 * @description: 更新商品
 * @param params
 */
export function updateGoods(params) {
  return http.request(
    {
      url: `/Admin/Goods/GoodsUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新商品
 * @param params
 */
export function updateGoodsStatus(params) {
  return http.request(
    {
      url: `/Admin/Goods/GoodsUpdateStatus`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建商品
 * @param params
 */
export function createGoods(params) {
  return http.request(
    {
      url: `/Admin/Goods/GoodsCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除商品
 * @param params
 */
export function deleteGoods(params) {
  return http.request(
    {
      url: `/Admin/Goods/GoodsDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
