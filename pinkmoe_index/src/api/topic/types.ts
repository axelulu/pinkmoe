/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 18:20:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:21
 * @FilePath: /pinkmoe_index/src/api/topic/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { BaseResponseType, ReqPage } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';

export interface ReqTopicPost extends ReqPage {
  value?: string | string[];
}

export interface ResTopicPost extends BaseResponseType {
  value?: string;
  sort?: number;
  label?: string;
  icon?: string;
  ID?: number;
  post?: Array<ResPost>;
}
