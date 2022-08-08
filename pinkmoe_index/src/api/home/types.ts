/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:59:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 10:06:48
 * @FilePath: /pinkmoe_index/src/api/home/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { BaseResponseType } from '/@/api/common/types';
import { ResCategory } from '/@/api/category/types';
import { ResFile } from '../upload/types';
import { ReqPostDownloadPost, ReqPostMusicPost, ReqPostVideoPost } from '../post/types';

export interface ResHomeList {
  content?: Array<ResContent>;
  popular?: Array<ResPost>;
  recommend?: Array<ResPost>;
}

export interface ResContent {
  icon?: string;
  iconType?: string;
  name?: string;
  slug?: string;
  sort?: number;
  style?: number;
  topic?: Array<any>;
  post?: Array<ResPost>;
}

export interface ResPost extends BaseResponseType {
  author?: string;
  category?: string;
  commentStatus?: string;
  content?: string;
  cover?: string;
  exerpt?: string;
  from?: string;
  postId?: string;
  postImg?: string;
  status?: string;
  title?: string;
  type?: string;
  view?: number;
  topic?: Array<any>;
  AuthorRelation?: ResAuthor;
  CategoryRelation?: ResCategory;
  collectRelation?: Array<any>;
  starRelation?: Array<any>;
  musicRelation?: Array<ReqPostMusicPost>;
  videoRelation?: Array<ReqPostVideoPost>;
  downloadRelation?: Array<ReqPostDownloadPost>;
  fileRelation?: Array<ResFile>;
}

export interface ResAuthor extends BaseResponseType {
  authorities?: ResAuthority;
  authority?: string;
  authorityId?: string;
  avatar?: string;
  cash?: number;
  credit?: number;
  desc?: string;
  email?: string;
  exp?: number;
  headerImg?: string;
  nickName?: string;
  phone?: string;
  sex?: string;
  userName?: string;
  uuid?: string;
}

export interface ResAuthority extends BaseResponseType {
  authorityId?: string;
  authorityName?: string;
  dataAuthorityId?: string;
  menus?: string;
}
