/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 08:41:08
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:27
 * @FilePath: /pinkmoe_index/src/api/upload/types.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ReqPage } from '/@/api/common/types';

export interface ReqFileList extends ReqPage {
  type?: string;
}

export interface ResFileList {
  file?: ResFile;
}

export interface ResFile {
  ID: number;
  key?: string;
  name?: string;
  postId?: string;
  tag?: string;
  type?: string;
  url?: string;
  uuid?: string;
}
