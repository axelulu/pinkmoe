/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-02 22:53:47
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:26:02
 * @FilePath: /pinkmoe_admin/src/views/system/user/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { h } from "vue";
import { NAvatar, NTag } from "naive-ui";

export const columns = [
  {
    title: "用户ID",
    key: "uuid"
  },
  {
    title: "用户头像",
    key: "avatar",
    width: 100,
    render(row) {
      return h(NAvatar, {
        size: 48,
        objectFit: "cover",
        src: row.avatar
      });
    }
  },
  {
    title: "用户手机号",
    key: "phone",
    render(row) {
      return h(
        NTag,
        {},
        {
          default: () => (row.phone ? row.phone : "暂无")
        }
      );
    }
  },
  {
    title: "用户邮箱",
    key: "email"
  },
  {
    title: "创建时间",
    key: "UpdatedAt"
  },
  {
    title: "用户状态",
    key: "DeletedAt",
    render(row) {
      return h(
        NTag,
        {
          type: !row.deleted_at ? "success" : "error"
        },
        {
          default: () => (!row.deleted_at ? "启用" : "删除")
        }
      );
    }
  }
];
