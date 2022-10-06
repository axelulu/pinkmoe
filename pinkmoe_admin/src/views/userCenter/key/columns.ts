/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-02 22:53:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:26:10
 * @FilePath: /pinkmoe_admin/src/views/userCenter/key/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { h } from "vue";
import { NTag } from "naive-ui";

export const columns = [
  {
    title: '卡密',
    key: 'key',
    width: 1000,
  },
  {
    title: '面额',
    key: 'value',
    width: 400,
  },
  {
    title: '类型',
    key: 'type',
    render(row) {
      return h(
        NTag,
        {
          round: true,
          tertiary: true
        },
        {
          default: () => (row.type === "cash" ? "现金" : "积分")
        }
      );
    }
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      return h(
        NTag,
        {
          type: row.status === "1" ? "success" : "default"
        },
        {
          default: () => (row.status === "1" ? "未使用" : "已使用")
        }
      );
    }
  },
];
