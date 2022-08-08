/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 15:45:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 11:02:19
 * @FilePath: /pinkmoe_admin/src/views/content/post/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { h } from "vue";
import { NAvatar, NButton, NSwitch, NTag, useMessage } from "naive-ui";
import {  updatePostStatus } from "@/api/content/post";
import { ResultEnum } from "@/enums/httpEnum";

export const columns = [
  {
    title: "文章ID",
    key: "postId"
  },
  {
    title: "封面",
    key: "cover",
    render(row) {
      return h(NAvatar, {
        size: 48,
        src: row.cover,
        objectFit: "cover",
      });
    }
  },
  {
    title: "标题",
    key: "title"
  },
  {
    title: "类型",
    key: "type",
    render(row) {
      return h(
        NTag,
        {
          round: true,
          tertiary: true
        },
        {
          default: () => (row.type === "post" ? "文章" : row.type === "music" ? "音乐" : row.type === "active" ? "动态" : "视频")
        }
      );
    }
  },
  {
    title: "作者",
    key: "author",
    render(row) {
      return h(
        NButton,
        {
          round: true,
          tertiary: true
        },
        {
          default: () => (row.AuthorRelation.nickName)
        }
      );
    }
  },
  {
    title: "分类",
    key: "category",
    render(row) {
      return h(
        NButton,
        {
          round: true,
          tertiary: true
        },
        {
          default: () => (row.CategoryRelation.name)
        }
      );
    }
  },
  {
    title: "状态",
    key: "status",
    render(row) {
      return h(
        NTag,
        {
          type: row.status === "published" ? "success" : row.status === "pending" ? "info" : "default"
        },
        {
          default: () => (row.status === "published" ? "已发布" : row.status === "pending" ? "待审核" : "草稿")
        }
      );
    }
  },
  {
    title: "发布",
    key: "status",
    render(row) {
      const message = useMessage();
      return h(NSwitch, {
        value: row.status === "published",
        onUpdateValue: async (e) => {
          if (e) {
            row.status = "published";
          } else {
            row.status = "pending";
          }
          row.topic = row.topic.map((item) => {
            return item.value;
          });
          const { code, message: msg } = await updatePostStatus(row);
          if (code == ResultEnum.SUCCESS) {
            message.success("更新成功");
          } else {
            message.error(msg || "更新失败");
          }
        }
      });
    }
  }
];
