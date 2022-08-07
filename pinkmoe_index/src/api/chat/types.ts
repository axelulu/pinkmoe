/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-30 16:26:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:30
 * @FilePath: /pinkmoe_index/src/api/chat/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ResAuthor } from '../home/types';
import { BaseResponseType, ReqPage } from '/@/api/common/types';

export interface ReqChat extends ReqPage {
  sendId?: string;
}

export interface ResChat extends BaseResponseType {
  ID?: number;
  amount?: number;
  cmd?: number;
  content?: string;
  media?: number;
  memo?: string;
  pic?: string;
  sendId?: string;
  url?: string;
  userId?: string;
  sendIdRelation?: ResAuthor;
  userIdRelation?: ResAuthor;
}

export interface ResChatRelation extends BaseResponseType {
  ID?: number;
  amount?: number;
  cmd?: number;
  content?: string;
  media?: number;
  memo?: string;
  pic?: string;
  sendId?: string;
  url?: string;
  userId?: string;
  sendIdRelation?: ResAuthor;
  userIdRelation?: ResAuthor;
}
