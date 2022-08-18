/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-19 17:39:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 09:41:49
 * @FilePath: /pinkmoe_index/src/api/category/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios';
import {
  ReqCategoryPost,
  ResCategory,
  ResCategoryGoods,
  ResCategoryPost,
} from '/@/api/category/types';
import { ResPage } from '/@/api/common/types';
// import axios from 'axios';
enum URL {
  list = '/api/Category/CategoryList',
  post = '/api/Post/CategoryPostList',
  shopList = '/api/GoodsCategory/GoodsCategoryList',
  shopPost = '/api/Goods/CategoryGoodsList',
}

const getCategoryList = async () =>
  get<ResPage<Array<ResCategory>>>({
    url: URL.list,
    params: {
      page: 1,
      pageSize: 999,
    },
  });

const getCategoryPostList = async (params: ReqCategoryPost) =>
  get<ResPage<ResCategoryPost>>({
    url: URL.post,
    params,
  });

const getShopCategoryList = async () =>
  get<ResPage<Array<ResCategory>>>({
    url: URL.shopList,
    params: {
      page: 1,
      pageSize: 999,
    },
  });

const getCategoryGoodsList = async (params: ReqCategoryPost) =>
  get<ResPage<ResCategoryGoods>>({
    url: URL.shopPost,
    params,
  });
export { getCategoryList, getCategoryPostList, getCategoryGoodsList, getShopCategoryList };
