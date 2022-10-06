<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:57:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:47
 * @FilePath: /pinkmoe_admin/src/views/monitor/userOnline/userOnline.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <n-card :bordered="false" class="proCard">
    <BasicForm @register="register" @submit="handleSubmit" @reset="handleReset">
      <template #statusSlot="{ model, field }">
        <n-input v-model:value="model[field]" />
      </template>
    </BasicForm>

    <BasicTable
      :columns="columns"
      :request="loadDataTable"
      :row-key="(row) => row.id"
      ref="actionRef"
      :actionColumn="actionColumn"
      :scroll-x="1090"
    />
  </n-card>
</template>

<script lang="ts" setup>
import { h, reactive, ref } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { BasicTable, TableAction } from "@/components/Table";
import { BasicForm, useForm } from "@/components/Form";
import { columns } from "./columns";
import { deleteUserOnline, getUserOnlineList } from "@/api/monitor/userOnline";
import { ResultEnum } from "@/enums/httpEnum";

const dialog = useDialog();

const schemas = [
  {
    field: "userName",
    component: "NInput",
    label: "用户名",
    componentProps: {
      placeholder: "请输入用户名",
      onInput: (e: any) => {
        formParams.userName = e;
      }
    }
  },
  {
    field: "explorer",
    component: "NInput",
    label: "浏览器",
    componentProps: {
      placeholder: "请输入浏览器",
      onInput: (e: any) => {
        formParams.explorer = e;
      }
    }
  },
  {
    field: "ip",
    component: "NInput",
    label: "ip",
    componentProps: {
      placeholder: "请输入ip",
      onInput: (e: any) => {
        formParams.ip = e;
      }
    }
  }
];
const message = useMessage();
const actionRef = ref();

const original = () => {
  return {
    ID: 0,
    userName: "",
    explorer: "",
    ip: "",
    os: ""
  };
};

const formParams = reactive(original());

const params = ref({
  pageSize: 10
});

const actionColumn = reactive({
  width: 150,
  title: "操作",
  key: "action",
  fixed: "right",
  render(record) {
    return h(TableAction as any, {
      style: "button",
      actions: [
        {
          label: "强制下线",
          icon: "ic:outline-delete-outline",
          onClick: handleDelete.bind(null, record),
          // 根据业务控制是否显示 isShow 和 auth 是并且关系
          ifShow: () => {
            return true;
          }
          // 根据权限控制是否显示: 有权限，会显示，支持多个
          // auth: ['basic_list'],
        }
      ]
    });
  }
});

const [register, {}] = useForm({
  gridProps: { cols: "1 s:1 m:2 l:3 xl:4 2xl:4" },
  labelWidth: 80,
  schemas
});

const loadDataTable = async (res) => {
  const {
    code,
    message: msg,
    result
  } = await getUserOnlineList({ ...formParams, ...params.value, ...res });
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  return result;
};

function handleDelete(record: Recordable) {
  dialog.info({
    title: "提示",
    content: `您想强制下线用户: ${record.userName}`,
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const { code, message: msg } = await deleteUserOnline(record);
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("下线成功");
        reloadTable();
      } else {
        message.error(msg || "下线失败");
      }
    },
    onNegativeClick: () => {
    }
  });
}

function reloadTable() {
  actionRef.value.reload();
}

function handleSubmit() {
  reloadTable();
}

function handleReset() {
  Object.assign(formParams, original());
}
</script>

<style lang="less" scoped></style>
