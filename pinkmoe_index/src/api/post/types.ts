/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:57:54
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:04
 * @FilePath: /pinkmoe_index/src/api/post/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ResAuthor, ResPost } from '/@/api/home/types';
import { BaseResponseType, ReqPage } from '/@/api/common/types';
import { ResFile } from '../upload/types';

export interface ReqPostItem {
  postId?: string | string[];
}

export interface ReqBbsPost extends ReqPage {
  type?: string;
  category?: string;
  author?: string; //作者筛选
  userId?: string; //收藏查询所需
}

export interface ReqPostDownloadPost {
  name?: string;
  url?: string;
  unpackPwd?: string;
  extractPwd?: string;
  price?: number;
  priceType?: string;
}

export interface ReqPostVideoPost {
  name?: string;
  url?: string;
  subtitles?: string;
  price?: number;
  priceType?: string;
}

export interface ReqPostMusicPost {
  name?: string;
  url?: string;
  price?: number;
  priceType?: string;
}

export interface ReqPublishPost extends BaseResponseType {
  postId?: string;
  title?: string;
  exerpt?: string;
  content?: string;
  category?: string;
  author?: string;
  cover?: string;
  postImg?: Array<ResFile>;
  type?: string;
  view?: number;
  from?: string;
  videoRelation?: Array<ReqPostVideoPost>;
  musicRelation?: Array<ReqPostMusicPost>;
  downloadRelation?: Array<ReqPostDownloadPost>;
  status?: string;
  topic?: Array<string>;
  commentStatus?: string;
}

export interface ReqBbsActive extends ReqPage {
  content?: string;
  title?: string;
  type?: string;
  author?: string;
}

export interface ResPostItem {
  authorPosts?: Array<ResPost>;
  comments?: Array<ResComment>;
  post?: ResPost;
  users?: Array<ResAuthor>;
  commentCount?: number;
  fansCount?: number;
  followCount?: number;
  followStatus?: boolean;
  postCount?: number;
}

export interface ResComment extends BaseResponseType {
  ID?: number;
  children?: Array<ResComment>;
  content?: string;
  formUid?: string;
  parentId?: number;
  postId?: string;
  status?: string;
  toUid?: string;
  type?: string;
  postRelation?: ResPost;
  ToUidRelation?: ResAuthor;
  FormUidRelation?: ResAuthor;
}
