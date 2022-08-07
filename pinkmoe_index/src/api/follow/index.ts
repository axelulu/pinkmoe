/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 18:39:28
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:43
 * @FilePath: /pinkmoe_index/src/api/follow/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios';
import { ReqFollowStatus } from '/@/api/follow/types';
import { IResponse } from '/@/utils/http/axios/type';
// import axios from 'axios';
enum URL {
  create = '/api/Follow/CreateFollow',
  delete = '/api/Follow/DeleteFollow',
  status = '/api/Follow/FollowStatus',
}

const createFollow = async (data: ReqFollowStatus) =>
  post<IResponse<any>>(
    {
      url: URL.create,
      data,
    },
    true,
  );

const deleteFollow = async (data: ReqFollowStatus) =>
  post<IResponse<any>>(
    {
      url: URL.delete,
      data,
    },
    true,
  );

const followStatus = async (params: ReqFollowStatus) =>
  get<boolean>({
    url: URL.status,
    params,
  });

export { createFollow, deleteFollow, followStatus };
