<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-10 20:11:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-10 19:07:21
 * @FilePath: /pinkmoe_admin/src/views/system/role/role.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <div>
    <n-card :bordered="false" class="mt-4 proCard">
      <BasicTable
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
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
    </n-card>

    <n-modal
      v-model:show="showMenuAuthModal"
      :show-icon="false"
      preset="dialog"
      :title="editRoleTitle"
    >
      <div class="py-3 menu-list">
        <n-tree
          block-line
          cascade
          checkable
          check-strategy="child"
          :virtual-scroll="true"
          :data="treeMenuData"
          :expandedKeys="expandedMenuKeys"
          v-model:checked-keys="checkedMenuKeys"
          style="max-height: 950px; overflow: hidden"
          @update:checked-keys="checkedMenuTree"
          @update:expanded-keys="onExpandedMenuKeys"
        />
      </div>
      <template #action>
        <n-space>
          <n-button type="info" ghost icon-placement="left" @click="packMenuHandle">
            全部{{ expandedMenuKeys.length ? "收起" : "展开" }}
          </n-button>
          <n-button type="info" ghost icon-placement="left" @click="checkedAllMenuHandle">
            全部{{ checkedAllMenu ? "取消" : "选择" }}
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmMenuForm"
          >提交
          </n-button
          >
        </n-space>
      </template>
    </n-modal>

    <n-modal
      v-model:show="showApiAuthModal"
      :show-icon="false"
      preset="dialog"
      :title="editRoleTitle"
    >
      <div class="py-3 menu-list">
        <n-tree
          block-line
          cascade
          checkable
          check-strategy="child"
          :virtual-scroll="true"
          :data="treeApiData"
          :expandedKeys="expandedApiKeys"
          v-model:checked-keys="checkedApiKeys"
          style="max-height: 950px; overflow: hidden"
          @update:checked-keys="checkedApiTree"
          @update:expanded-keys="onExpandedApiKeys"
        />
      </div>
      <template #action>
        <n-space>
          <n-button type="info" ghost icon-placement="left" @click="packApiHandle">
            全部{{ expandedApiKeys.length ? "收起" : "展开" }}
          </n-button>
          <n-button type="info" ghost icon-placement="left" @click="checkedAllApiHandle">
            全部{{ checkedAllApi ? "取消" : "选择" }}
          </n-button>
          <n-button type="primary" :loading="formBtnLoading" @click="confirmApiForm">提交</n-button>
        </n-space>
      </template>
    </n-modal>

    <CreateEditForm @reloadTable="reloadResetTable" ref="createEditFormRef" :formType="formType" />
  </div>
</template>

<script lang="ts" setup>
import { h, onMounted, reactive, ref, unref } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { BasicTable, TableAction } from "@/components/Table";
import { getRoleList } from "@/api/system/authority";
import { getAllApi, getApiAuthority, updateApiAuthority } from "@/api/system/api";
import { getAllMenu, getMenuAuthority } from "@/api/system/menu";
import { updateMenuAuthority, deleteAuthority } from "@/api/system/authority";
import { updateMenuTreeNode, updateMenuTreeNodeResult } from "@/utils";
import { columns } from "./columns";
import { PlusOutlined } from "@vicons/antd";
import { ResultEnum } from "@/enums/httpEnum";

import CreateEditForm from "./CreateEditForm.vue";

const dialog = useDialog();

const message = useMessage();
const actionRef = ref();

const showMenuAuthModal = ref(false);
const showApiAuthModal = ref(false);
const formBtnLoading = ref(false);
const checkedAllMenu = ref(false);
const checkedAllApi = ref(false);
const editRoleTitle = ref("");
const treeMenuData = ref([]);
const treeApiData = ref([]);
const treeMenuDataUser = ref([]);
const expandedApiKeys = ref([]);
const checkedApiKeys = ref([""]);
const expandedMenuKeys = ref([]);
const checkedMenuKeys = ref([""]);
const checkedMenuValues = ref([""]);
const checkedApiValues = ref([""]);
const checkedId = ref("");
const createEditFormRef = ref();
const formType = ref("create");

const original = () => {
  return {
    authorityName: "",
    authorityWeight: 10,
    vipStart: 1,
    authorityId: "",
    authorityColor: '',
    authorityParams: [
      {
        authorityParamKey: '',
        authorityParamValue: '',
        authorityParamDay: '',
        authorityParamId: '',
      }
    ],
  };
};

const formParams = reactive(original());

const params = ref({
  pageSize: 5
});

const actionColumn = reactive({
  width: 350,
  title: "操作",
  key: "action",
  fixed: "right",
  render(record) {
    return h(TableAction, {
      style: "button",
      actions: [
        {
          label: "菜单权限",
          onClick: handleMenuAuth.bind(null, record),
          // 根据业务控制是否显示 isShow 和 auth 是并且关系
          ifShow: () => {
            return true;
          }
          // 根据权限控制是否显示: 有权限，会显示，支持多个
          // auth: ['basic_list'],
        },
        {
          label: "api权限",
          onClick: handleApiAuth.bind(null, record),
          // 根据业务控制是否显示 isShow 和 auth 是并且关系
          ifShow: () => {
            return true;
          }
          // 根据权限控制是否显示: 有权限，会显示，支持多个
          // auth: ['basic_list'],
        },
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

const loadDataTable = async (res: any) => {
  const {
    code,
    message: msg,
    result
  } = await getRoleList({ ...formParams, ...params.value, ...res });
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

async function confirmMenuForm(e: any) {
  e.preventDefault();
  formBtnLoading.value = true;
  const { code, message: msg } = await updateMenuAuthority({
    authorityId: checkedId.value,
    menus: checkedMenuValues.value
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    message.success("修改成功");
    setTimeout(() => {
      showMenuAuthModal.value = false;
      reloadTable();
    });
  } else {
    message.error(msg || "修改失败");
  }
  formBtnLoading.value = false;
}

async function confirmApiForm(e: any) {
  e.preventDefault();
  formBtnLoading.value = true;
  const { code, message: msg } = await updateApiAuthority({
    authorityId: checkedId.value,
    casbinInfos: checkedApiValues.value
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    message.success("修改成功");
    setTimeout(() => {
      showApiAuthModal.value = false;
      reloadTable();
    });
  } else {
    message.error(msg || "修改失败");
  }
  formBtnLoading.value = false;
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
    content: `您想删除角色: ${record.authorityName}`,
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const { code, message: msg } = await deleteAuthority(record);
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

async function handleMenuAuth(record: Recordable) {
  editRoleTitle.value = `分配 ${record.authorityName} 的菜单权限`;
  checkedMenuKeys.value = record.menu_keys;
  checkedId.value = record.authorityId;
  const {
    code,
    message: msg,
    result
  } = await getMenuAuthority({
    authorityId: record.authorityId
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    treeMenuDataUser.value = result;
    checkedMenuKeys.value = [];
    checkedMenuValues.value = [];
    let tree = treeMenuData.value.concat();
    updateMenuTreeNodeResult(treeMenuData.value, tree, null, (option, item) => {
      updateMenuTreeNode(option, item.parentId, option, (parentId, item) => {
        if (item.ID === parentId) {
          treeMenuDataUser.value = treeMenuDataUser.value.filter((item) => item.ID != parentId);
          return;
        }
      });
    });
    getMenuTreeTitle(treeMenuDataUser.value);
    showMenuAuthModal.value = true;
    initTree(checkedMenuValues.value);
  } else {
    message.error(msg || "获取失败");
  }
}

async function handleApiAuth(record: Recordable) {
  editRoleTitle.value = `分配 ${record.authorityName} 的api权限`;
  checkedApiKeys.value = record.menu_keys;
  checkedId.value = record.authorityId;
  const {
    code,
    message: msg,
    result
  } = await getApiAuthority({
    authorityId: record.authorityId
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    checkedApiKeys.value = [];
    checkedApiValues.value = [];
    getApiTreeTitle(result);
    showApiAuthModal.value = true;
  } else {
    message.error(msg || "获取失败");
  }
}

function initCheckedTree(option) {
  let tree = treeMenuData.value.concat();
  updateMenuTreeNodeResult(treeMenuData.value, option, tree, (option, item) => {
    if (option.includes(item)) {
      updateMenuTreeNode(tree, item.parentId, option, (parentId, item, option) => {
        if (item.ID === parentId && !option.includes(item)) {
          option.push(item);
          return;
        }
      });
    }
  });
}

function initTree(option) {
  let tree = treeMenuData.value.concat();
  updateMenuTreeNodeResult(treeMenuData.value, option, tree, (option, item) => {
    updateMenuTreeNode(tree, item.parentId, option, (parentId, item, option) => {
      if (item.ID === parentId && !option.includes(item)) {
        option.push(item);
        return;
      }
    });
  });
}

function checkedMenuTree(keys, option) {
  initCheckedTree(option);
  checkedMenuKeys.value = [...keys];
  checkedMenuValues.value = [...option];
}

function checkedApiTree(keys, option) {
  checkedApiKeys.value = [...keys];
  checkedApiValues.value = [...option];
}

function onExpandedMenuKeys(keys) {
  expandedMenuKeys.value = keys;
}

function onExpandedApiKeys(keys) {
  expandedApiKeys.value = keys;
}

function packMenuHandle() {
  if (expandedMenuKeys.value.length) {
    expandedMenuKeys.value = [];
  } else {
    expandedMenuKeys.value = treeMenuData.value.map((item: any) => item.key) as [];
  }
}

function packApiHandle() {
  if (expandedApiKeys.value.length) {
    expandedApiKeys.value = [];
  } else {
    expandedApiKeys.value = treeApiData.value.map((item: any) => item.key) as [];
  }
}

function checkedAllMenuHandle() {
  if (!checkedAllMenu.value) {
    checkedMenuKeys.value = [];
    checkedMenuValues.value = [];
    getMenuTreeTitle(treeMenuData.value);
    checkedAllMenu.value = true;
  } else {
    checkedMenuKeys.value = [];
    checkedMenuValues.value = [];
    checkedAllMenu.value = false;
  }
}

function checkedAllApiHandle() {
  if (!checkedAllApi.value) {
    checkedApiKeys.value = [];
    checkedApiValues.value = [];
    treeApiData.value.forEach((item) => {
      item.children.forEach((i) => {
        checkedApiKeys.value.push(i.path);
        checkedApiValues.value.push(i);
      });
    });
    checkedAllApi.value = true;
  } else {
    checkedApiKeys.value = [];
    checkedApiValues.value = [];
    checkedAllApi.value = false;
  }
}

function getMenuTreeTitle(treeMenuList) {
  treeMenuList.forEach((item) => {
    item.label = item.meta.title;
    item.key = item.name;
    checkedMenuKeys.value.push(item.name);
    checkedMenuValues.value.push(item);
    if (item.children !== null) {
      getMenuTreeTitle(item.children);
    }
  });
}

function getApiTreeTitle(treeApiList) {
  let apiList = [];
  if (treeApiList && treeApiList.length > 0) {
    treeApiList.forEach((item) => {
      item.label = item.description;
      item.key = item.path;
      checkedApiKeys.value.push(item.path);
      checkedApiValues.value.push(item);
      let apigroup: {}[];
      apigroup = treeApiList.filter((one) => one.apiGroup === item.apiGroup);
      let list = {
        key: item.apiGroup,
        label: item.apiGroup + "组",
        children: apigroup
      };
      if (apiList.filter((one) => one.key === item.apiGroup).length === 0) apiList.push(list);
    });
  }
  return apiList;
}

async function getMenuTree() {
  const { code, message: msg, result } = await getAllMenu();
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    let treeMenuList = result;
    getMenuTreeTitle(treeMenuList);
    expandedMenuKeys.value = treeMenuList.map((item) => item.key);
    treeMenuData.value = treeMenuList;
  } else {
    message.error(msg || "获取失败");
  }
}

async function getApiTree() {
  const { code, message: msg, result } = await getAllApi();
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    let treeApiList = result;
    treeApiList = getApiTreeTitle(treeApiList);
    expandedApiKeys.value = treeApiList.map((item) => item.key);
    treeApiData.value = treeApiList;
  } else {
    message.error(msg || "获取失败");
  }
}

onMounted(async () => {
  await getApiTree();
  await getMenuTree();
});
</script>

<style lang="less" scoped></style>
