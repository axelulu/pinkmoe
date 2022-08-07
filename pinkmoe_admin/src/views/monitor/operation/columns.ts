/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:34:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:40
 * @FilePath: /pinkmoe_admin/src/views/monitor/operation/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { h } from 'vue';
import { NTag, NCode } from 'naive-ui';

export const columns = [
  {
    title: '操作用户',
    key: 'userId',
  },
  {
    title: '请求IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '方法',
    key: 'method',
    width: 100,
    render(row) {
      return h(
        NTag,
        {},
        {
          default: () => (row.method ? row.method : '暂无'),
        }
      );
    },
  },
  {
    title: '状态码',
    key: 'status',
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          type: row.status == 200 ? 'success' : 'error',
        },
        {
          default: () => (row.status ? row.status : '暂无'),
        }
      );
    },
  },
  {
    title: 'api路径',
    key: 'path',
  },
  {
    title: '请求',
    key: 'body',
    width: 200,
    render(row) {
      return h(NCode, {
        language: 'json',
        code: row.body,
        wordWrap: true,
      });
    },
  },
  {
    title: '响应',
    key: 'resp',
    render(row) {
      return h(NCode, {
        language: 'json',
        code: row.resp,
      });
    },
  },
];
