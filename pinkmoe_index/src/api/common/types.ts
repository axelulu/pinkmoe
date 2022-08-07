/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 20:58:42
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:40
 * @FilePath: /pinkmoe_index/src/api/common/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
export interface BaseResponseType {
  CreatedAt?: Date;
  UpdatedAt?: Date;
}

export interface ReqPage {
  page?: number;
  pageSize?: number;
  orderKey?: string;
  desc?: boolean;
}

export interface ResResultData<T = any> {
  code?: number;
  result?: T;
  message: string;
}

export interface ResPage<T = any> {
  list?: T;
  page?: number;
  pageCount?: number;
  pageSize: number;
}
