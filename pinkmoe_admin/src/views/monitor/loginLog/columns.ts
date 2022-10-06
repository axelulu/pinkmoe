/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:56:23
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-11 09:30:32
 * @FilePath: /pinkmoe_admin/src/views/monitor/loginLog/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { h } from 'vue';
import { NTag } from 'naive-ui';

export const columns = [
  {
    title: '登陆用户',
    key: 'userName',
    width: 150,
  },
  {
    title: '登陆IP',
    key: 'ip',
    width: 150,
  },
  {
    title: '登陆地点',
    key: 'loginLocation',
    width: 200,
  },
  {
    title: '浏览器',
    key: 'explorer',
    width: 200,
  },
  {
    title: '操作系统',
    key: 'os',
    width: 300,
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
    width: 100,
  },
  {
    title: '更新时间',
    key: 'UpdatedAt',
    width: 300,
  },
];
