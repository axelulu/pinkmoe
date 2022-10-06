/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 15:45:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-17 21:52:48
 * @FilePath: /pinkmoe_admin/src/views/shop/goods/columns.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
 */
import { h } from "vue";
import { NAvatar, NButton, NSwitch, NTag, useMessage } from "naive-ui";
import { ResultEnum } from "@/enums/httpEnum";
import { updateGoodsStatus } from "@/api/shop/goods";

export const columns = [
  {
    title: "商品ID",
    key: "goodsId"
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
          default: () => (row.type === "physical" ? "实物" : "虚拟")
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
          row.sizeRelation = null
          row.sizeValRelation = null
          const { code, message: msg } = await updateGoodsStatus(row);
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
