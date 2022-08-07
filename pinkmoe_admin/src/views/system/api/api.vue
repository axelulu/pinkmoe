<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 22:00:10
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:40
 * @FilePath: /pinkmoe_admin/src/views/system/api/api.vue
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
      :actionColumn="actionColumn"
      :scroll-x="1090"
    >
      <template #tableTitle>
        <n-button type="primary" @click="handleCreate">
          <template #icon>
            <n-icon>
              <PlusOutlined />
            </n-icon>
          </template>
          新建
        </n-button>
      </template>
    </BasicTable>
    <CreateEditForm @reloadTable="reloadResetTable" ref="createEditFormRef" :formType="formType" />
  </n-card>
</template>

<script lang="ts" setup>
import { h, reactive, ref } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { BasicTable, TableAction } from "@/components/Table";
import { BasicForm, useForm } from "@/components/Form/index";
import { columns } from "./columns";
import { PlusOutlined } from "@vicons/antd";
import { getApiList, deleteApi } from "@/api/system/api";
import CreateEditForm from "./CreateEditForm.vue";
import { ResultEnum } from "@/enums/httpEnum";

const dialog = useDialog();
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
    field: "path",
    component: "NInput",
    label: "路径",
    componentProps: {
      placeholder: "请输入路径",
      onInput: (e: any) => {
        formParams.path = e;
      }
    },
    rules: [{ required: false, message: "请输入路径", trigger: ["blur"] }]
  },
  {
    field: "apiGroup",
    component: "NInput",
    label: "api组",
    componentProps: {
      placeholder: "请输入api组",
      onInput: (e: any) => {
        formParams.apiGroup = e;
      }
    }
  }
];
const formType = ref("create");
const createEditFormRef = ref();

const message = useMessage();
const actionRef = ref();

const original = () => {
  return {
    ID: 0,
    path: "",
    method: "",
    apiGroup: "",
    description: ""
  };
};

let formParams = reactive(original());

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
          label: "编辑",
          onClick: handleEdit.bind(null, record),
          ifShow: () => {
            return true;
          }
          // auth: ['basic_list'],
        },
        {
          label: "删除",
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
  } = await getApiList({ ...formParams, ...params.value, ...res });
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  return result;
};

function reloadResetTable() {
  Object.assign(formParams, original());
  actionRef.value.reload();
}

function reloadTable() {
  actionRef.value.reload();
}

function openCreateEditForm(record) {
  const { openForm } = createEditFormRef.value;
  openForm(record);
}

function handleCreate() {
  formType.value = "create";
  openCreateEditForm(formParams);
}

function handleEdit(record: Recordable) {
  formType.value = "edit";
  openCreateEditForm(record);
}

function handleDelete(record: Recordable) {
  dialog.info({
    title: "提示",
    content: `您想删除角色: ${record.path}`,
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const { code, message: msg } = await deleteApi(record);
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("删除成功");
        reloadTable();
      } else {
        message.error(msg || "删除失败");
      }
    },
    onNegativeClick: () => {
    }
  });
}

function handleSubmit() {
  reloadTable();
}

function handleReset() {
  Object.assign(formParams, original());
}
</script>

<style lang="less" scoped></style>
