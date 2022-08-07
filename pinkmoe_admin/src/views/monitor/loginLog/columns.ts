/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:56:23
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:34
 * @FilePath: /pinkmoe_admin/src/views/monitor/loginLog/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: '登陆用户',
    key: 'userName',
  },
  {
    title: '登陆IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '登陆地点',
    key: 'loginLocation',
    width: 100,
  },
  {
    title: '浏览器',
    key: 'explorer',
    width: 100,
  },
  {
    title: '操作系统',
    key: 'os',
    width: 100,
  },
  {
    title: '状态',
    key: 'status',
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          type: row.status ? 'success' : 'error',
        },
        {
          default: () => (row.status ? '成功' : '失败'),
        }
      );
    },
  },
  {
    title: '操作信息',
    key: 'msg',
  },
  {
    title: '更新时间',
    key: 'UpdatedAt',
    width: 100,
  },
];
