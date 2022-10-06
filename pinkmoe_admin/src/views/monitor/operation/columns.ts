/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:34:36
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-11 09:31:44
 * @FilePath: /pinkmoe_admin/src/views/monitor/operation/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { h } from 'vue';
import { NTag, NCode } from 'naive-ui';

export const columns = [
  {
    title: '操作用户',
    key: 'userId',
    width: 350,
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
    width: 300,
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
