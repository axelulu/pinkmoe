/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-28 15:36:26
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:14:42
 * @FilePath: /pinkmoe_admin/src/api/common.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
export interface BasicResponseModel<T = any> {
  code: number;
  message: string;
  result: T;
}
