<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:57:34
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:42
 * @FilePath: /pinkmoe_admin/src/views/monitor/operation/operation.vue
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
      :scroll-x="1090"
    >
    </BasicTable>
  </n-card>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useMessage } from "naive-ui";
import { BasicTable } from "@/components/Table";
import { BasicForm, useForm } from "@/components/Form";
import { columns } from "./columns";
import { PlusOutlined } from "@vicons/antd";
import { getOperationList } from "@/api/monitor/operation";
import { ResultEnum } from "@/enums/httpEnum";

const schemas = [
  {
    field: "method",
    component: "NSelect",
    label: "方法",
    componentProps: {
      placeholder: "请选择方法",
      options: [
        {
          label: "GET",
          value: "GET"
        },
        {
          label: "POST",
          value: "POST"
        }
      ],
      onUpdateValue: (e: any) => {
        formParams.method = e;
      }
    }
  },
  {
    field: "status",
    component: "NSelect",
    label: "状态码",
    componentProps: {
      placeholder: "请选择状态码",
      options: [
        {
          label: "200",
          value: 200
        },
        {
          label: "-1",
          value: -1
        }
      ],
      onUpdateValue: (e: any) => {
        formParams.status = e;
      }
    }
  },
  {
    field: "userId",
    component: "NInput",
    label: "用户ID",
    componentProps: {
      placeholder: "请输入用户ID",
      onInput: (e: any) => {
        formParams.userId = e;
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
    userId: "",
    method: "GET",
    ip: "",
    status: ""
  };
};

const formParams = reactive(original());

const params = ref({
  pageSize: 10
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
  } = await getOperationList({ ...formParams, ...params.value, ...res });
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  return result;
};

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
