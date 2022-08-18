/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:38:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 09:44:50
 * @FilePath: /pinkmoe_index/src/api/category/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { BaseResponseType } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';
import { ReqPage } from '/@/api/common/types';
import { ResGoods } from '../shop/types';

export interface ReqCategoryPost extends ReqPage {
  category?: string | string[];
}

export interface ResCategory extends BaseResponseType {
  ID: number;
  name?: string;
  iconType?: string;
  icon?: string;
  slug?: string;
  sort?: string;
  parentId?: string;
  children?: Array<ResCategory>;
}

export interface ResCategoryPost {
  category?: Array<ResCategory>;
  post?: Array<ResPost>;
}

export interface ResCategoryGoods {
  category?: Array<ResCategory>;
  goods?: Array<ResGoods>;
}
