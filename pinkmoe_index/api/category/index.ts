/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-19 17:39:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 09:41:49
 * @FilePath: /pinkmoe_index/src/api/category/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type {
  ReqCategoryPost,
  ResCategory,
  ResCategoryGoods,
  ResCategoryPost,
} from '/@/api/category/types'
import type { ResPage } from '/@/api/common/types'
// import axios from 'axios';
enum URL {
  list = '/Cms/Category/CategoryList',
  post = '/Cms/Post/CategoryPostList',
  shopList = '/Cms/GoodsCategory/GoodsCategoryList',
  shopPost = '/Cms/Goods/CategoryGoodsList',
}

const getCategoryList = async () =>
  get<ResPage<Array<ResCategory>>>({
    url: URL.list,
    params: {
      page: 1,
      pageSize: 999,
    },
  })

const getCategoryPostList = async (params: ReqCategoryPost) =>
  get<ResPage<ResCategoryPost>>({
    url: URL.post,
    params,
  })

const getShopCategoryList = async () =>
  get<ResPage<Array<ResCategory>>>({
    url: URL.shopList,
    params: {
      page: 1,
      pageSize: 999,
    },
  })

const getCategoryGoodsList = async (params: ReqCategoryPost) =>
  get<ResPage<ResCategoryGoods>>({
    url: URL.shopPost,
    params,
  })
export { getCategoryList, getCategoryPostList, getCategoryGoodsList, getShopCategoryList }
