/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:06:33
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-11 09:32:37
 * @FilePath: /pinkmoe_admin/src/views/monitor/userOnline/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
export const columns = [
  {
    title: '用户名',
    key: 'userName',
    width: 150,
  },
  {
    title: '请求IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '浏览器',
    key: 'explorer',
    width: 100,
  },
  {
    title: '操作系统',
    key: 'os',
  },
  {
    title: '更新时间',
    key: 'UpdatedAt',
    render(row) {
      return row.UpdatedAt;
    },
  },
];
