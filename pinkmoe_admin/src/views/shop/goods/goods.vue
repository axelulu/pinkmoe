<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 15:47:44
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-17 14:29:38
 * @FilePath: /pinkmoe_admin/src/views/shop/goods/goods.vue
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
    <CreateEditForm @reloadTable="reloadPostTable" ref="createEditFormRef" :formType="formType" />
  </n-card>
</template>

<script lang="ts" setup>
import { h, reactive, ref } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { BasicTable, TableAction } from "@/components/Table";
import { BasicForm, useForm } from "@/components/Form/index";
import { columns } from "./columns";
import { PlusOutlined } from "@vicons/antd";
import { getGoodsList, deleteGoods } from "@/api/shop/goods";
import CreateEditForm from "./CreateEditForm.vue";
import { ResultEnum } from "@/enums/httpEnum";

const dialog = useDialog();
const schemas = [
  {
    field: "title",
    component: "NInput",
    label: "标题",
    componentProps: {
      placeholder: "请输入标题",
      onInput: (e: any) => {
        formParams.title = e;
      }
    },
    rules: [{ required: false, message: "请输入标题", trigger: ["blur"] }]
  },
  {
    field: "type",
    component: "NSelect",
    label: "类型",
    componentProps: {
      placeholder: "请选择类型",
      options: [
        {
          label: "全部",
          value: ""
        },
        {
          label: "文章",
          value: "post"
        },
        {
          label: "音乐",
          value: "music"
        },
        {
          label: "视频",
          value: "video"
        }
      ],
      onUpdateValue: (e: any) => {
        formParams.type = e;
      }
    }
  },
  {
    field: "status",
    component: "NSelect",
    label: "文章状态",
    componentProps: {
      placeholder: "请选择状态",
      options: [
        {
          label: "全部",
          value: ""
        },
        {
          label: "草稿",
          value: "draft"
        },
        {
          label: "待审核",
          value: "pending"
        },
        {
          label: "已发布",
          value: "published"
        }
      ],
      onUpdateValue: (e: any) => {
        formParams.status = e;
      }
    }
  },
  {
    field: "content",
    component: "NInput",
    label: "内容",
    componentProps: {
      placeholder: "请输入内容",
      onInput: (e: any) => {
        formParams.content = e;
      }
    },
    rules: [{ required: false, message: "请输入内容", trigger: ["blur"] }]
  }
];
const formType = ref("");
const createEditFormRef = ref();

const message = useMessage();
const actionRef = ref();

let formParams = reactive({
  goodsId: "",
  title: "",
  content: "",
  category: "",
  author: null,
  cover: "",
  goodsImg: [],
  type: "",
  view: 0,
  sizeRelation: [],
  sizeValRelation: [],
  status: "",
});

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
  } = await getGoodsList({ ...formParams, ...params.value, ...res });
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  return result;
};

function reloadTable() {
  actionRef.value.reload();
}

function reloadPostTable() {
  formParams = {
    goodsId: "",
    title: "",
    content: "",
    category: "",
    author: null,
    cover: "",
    goodsImg: [],
    type: "",
    view: 0,
    sizeRelation: [],
    sizeValRelation: [],
    status: "",
  };
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
    content: `您想删除文章: ${record.title}`,
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const { code, message: msg } = await deleteGoods(record);
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

function handleSubmit(values: Recordable) {
  console.log(values);
  reloadTable();
}

function handleReset(values: Recordable) {
  formParams.title = "";
  formParams.content = "";
  formParams.status = "";
  formParams.type = "";
}
</script>

<style lang="less" scoped></style>
