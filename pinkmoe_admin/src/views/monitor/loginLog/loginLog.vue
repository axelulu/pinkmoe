<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 21:57:21
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:36
 * @FilePath: /pinkmoe_admin/src/views/monitor/loginLog/loginLog.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
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
    />
  </n-card>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useMessage } from "naive-ui";
import { BasicTable } from "@/components/Table";
import { BasicForm, useForm } from "@/components/Form";
import { columns } from "./columns";
import { getLoginLogList } from "@/api/monitor/loginLog";
import { ResultEnum } from "@/enums/httpEnum";

const schemas = [
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
  },
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
  }
];
const message = useMessage();
const actionRef = ref();

const original = () => {
  return {
    ID: 0,
    explorer: "",
    userName: "",
    ip: "",
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
  } = await getLoginLogList({ ...formParams, ...params.value, ...res });
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
