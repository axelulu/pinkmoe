/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-02 22:53:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:26:10
 * @FilePath: /pinkmoe_admin/src/views/userCenter/key/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
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
