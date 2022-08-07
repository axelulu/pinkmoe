/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-02 22:53:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:23:08
 * @FilePath: /pinkmoe_admin/src/views/content/topic/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { h } from "vue";
import { NAvatar } from "naive-ui";

export const columns = [
  {
    title: '名称',
    key: 'label',
    width: 200,
  },
  {
    title: '标识',
    key: 'value',
    width: 200,
  },
  {
    title: '图标',
    key: 'icon',
    width: 200,
    render(row) {
      return h(NAvatar, {
        size: 48,
        objectFit: "cover",
        src: row.icon,
      });
    },
  },
];
