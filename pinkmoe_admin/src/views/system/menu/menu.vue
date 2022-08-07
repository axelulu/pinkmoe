<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 22:11:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:51
 * @FilePath: /pinkmoe_admin/src/views/system/menu/menu.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <div>
    <n-grid class="mt-4" cols="1 s:1 m:1 l:3 xl:3 2xl:3" responsive="screen" :x-gap="12">
      <n-gi span="1">
        <n-card :segmented="{ content: 'hard' }" :bordered="false" size="small">
          <template #header>
            <n-space>
              <n-dropdown trigger="hover" @select="selectAddMenu" :options="addMenuOptions">
                <n-button type="info" ghost icon-placement="right">
                  添加菜单
                  <template #icon>
                    <div class="flex items-center">
                      <n-icon size="14">
                        <DownOutlined />
                      </n-icon>
                    </div>
                  </template>
                </n-button>
              </n-dropdown>
              <n-button type="info" ghost icon-placement="left" @click="packHandle">
                全部{{ expandedKeys.length ? "收起" : "展开" }}
                <template #icon>
                  <div class="flex items-center">
                    <n-icon size="14">
                      <AlignLeftOutlined />
                    </n-icon>
                  </div>
                </template>
              </n-button>
            </n-space>
          </template>
          <div class="w-full menu">
            <n-input type="input" v-model:value="pattern" placeholder="输入菜单名称搜索">
              <template #suffix>
                <n-icon size="18" class="cursor-pointer">
                  <SearchOutlined />
                </n-icon>
              </template>
            </n-input>
            <div class="py-3 menu-list">
              <template v-if="loading">
                <div class="flex items-center justify-center py-4">
                  <n-spin size="medium" />
                </div>
              </template>
              <template v-else>
                <n-tree
                  block-line
                  cascade
                  checkable
                  :virtual-scroll="true"
                  :pattern="pattern"
                  :data="treeData"
                  :expandedKeys="expandedKeys"
                  style="max-height: 650px; overflow: hidden"
                  @update:selected-keys="selectedTree"
                  @update:expanded-keys="onExpandedKeys"
                />
              </template>
            </div>
          </div>
        </n-card>
      </n-gi>
      <n-gi span="2">
        <n-card :segmented="{ content: 'hard' }" :bordered="false" size="small">
          <template #header>
            <n-space>
              <n-icon size="18">
                <FormOutlined />
              </n-icon>
              <span>编辑菜单{{ treeItemTitle ? `：${treeItemTitle}` : "" }}</span>
            </n-space>
          </template>
          <n-alert type="info" closable> 从菜单列表选择一项后，进行编辑</n-alert>
          <n-form
            :model="formParams"
            :rules="rules"
            ref="formRef"
            label-placement="left"
            :label-width="100"
            v-if="isEditMenu"
            class="py-4"
          >
            <n-form-item label="类型" path="type">
              <span>{{ !formParams.hidden ? "侧边栏菜单" : "" }}</span>
            </n-form-item>
            <n-form-item label="标题" path="label">
              <n-input placeholder="请输入标题" v-model:value="formParams.meta.title" />
            </n-form-item>
            <n-form-item label="标识" path="name">
              <n-input placeholder="请输入标识" v-model:value="formParams.name" />
            </n-form-item>
            <n-form-item label="路径" path="path">
              <n-input placeholder="请输入路径" v-model:value="formParams.path" />
            </n-form-item>
            <n-form-item label="文件路径" path="component">
              <n-input placeholder="请输入文件路径" v-model:value="formParams.component" />
            </n-form-item>
            <n-form-item label="排序" path="sort">
              <n-input-number max="100" min="0" v-model:value="formParams.sort" />
            </n-form-item>
            <n-form-item label="图标" path="icons">
              <n-input placeholder="请输入图标" v-model:value="formParams.meta.icon" />
            </n-form-item>
            <n-form-item label="父节点" path="parentId">
              <n-cascader
                :virtual-scroll="true"
                cascade="false"
                label-field="label"
                value-field="ID"
                v-model:value="formParams.parentId"
                :options="menuTreeData"
              />
            </n-form-item>
            <n-form-item label="是否隐藏" path="hidden">
              <n-switch v-model:value="formParams.hidden" />
            </n-form-item>
            <n-form-item path="auth" style="margin-left: 100px">
              <n-space>
                <n-button type="primary" :loading="subLoading" @click="formSubmit"
                >保存修改
                </n-button
                >
                <n-button type="error" @click="handleDelete">删除</n-button>
              </n-space>
            </n-form-item>
          </n-form>
        </n-card>
      </n-gi>
    </n-grid>
    <CreateDrawer
      ref="createDrawerRef"
      @reloadForm="reloadForm"
      :title="drawerTitle"
      :menuTreeData="menuTreeData"
      :treeItemId="treeItemId"
      :treeItemK="treeItemK"
    />
  </div>
</template>
<script lang="ts" setup>
import { ref, unref, reactive, onMounted, computed } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { DownOutlined, AlignLeftOutlined, SearchOutlined, FormOutlined } from "@vicons/antd";
import { getMenuList, updateMenu, deleteMenu } from "@/api/system/menu";
import { getTreeItem, getMenuTreeTitle } from "@/utils";
import CreateDrawer from "./CreateDrawer.vue";
import { ResultEnum } from "@/enums/httpEnum";

const rules = {
  label: {
    required: true,
    message: "请输入2到16位（字母,汉字,数字,下划线）",
    trigger: "blur",
    pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,16}$/ //验证菜单名称 2到16位（字母,汉字,数字,下划线）
  },
  name: {
    required: true,
    message: "请输入2到30位（字母）",
    trigger: "blur",
    pattern: /^[a-zA-Z_-]{2,30}$/ //验证菜单标识 2到30位（字母）
  },
  path: {
    required: true,
    message: "请输入2到100位（字母）",
    trigger: "blur",
    pattern: /^[a-zA-Z/_-]{2,100}$/ //验证菜单路径 2到100位（字母）
  },
  component: {
    required: true,
    message: "请输入2到100位（字母）",
    trigger: "blur",
    pattern: /^[a-zA-Z/_-]{2,100}$/ //验证菜单文件 2到100位（字母）
  },
  sort: {
    required: false,
    message: "请输入1-2位数字",
    trigger: "blur",
    pattern: /^[0-9]{1,2}$/ //验证排序内容 1-2位数字
  },
  icon: {
    required: true,
    message: "请输入2到30位（字母）",
    trigger: "blur",
    pattern: /^[a-zA-Z_-]{2,30}$/ //验证图标 2到30位（字母）
  },
  parentId: {
    required: true,
    message: "请输入1-10位数字",
    trigger: "blur",
    pattern: /^[0-9]{1,10}$/ //验证父菜单 1-2位数字
  },
  hidden: {
    required: true,
    message: "请选择显示状态",
    trigger: "blur",
    pattern: /^[a-zA-Z]{4,5}$/ //验证显示隐藏 4到5位
  }
};

const formRef: any = ref(null);
const createDrawerRef = ref();
const message = useMessage();

let treeItemKey = ref([]);

let treeItemId = ref(0);
let treeItemK = ref("home");

let expandedKeys = ref([]);

const treeData = ref([]);
const menuTreeData = ref([]);
const dialog = useDialog();

const loading = ref(true);
const subLoading = ref(false);
const isEditMenu = ref(false);
const treeItemTitle = ref("");
const pattern = ref("");
const drawerTitle = ref("");

const isAddSon = computed(() => {
  return !treeItemKey.value.length;
});

const addMenuOptions = ref([
  {
    label: "添加顶级菜单",
    key: "home",
    disabled: false
  },
  {
    label: "添加子菜单",
    key: "son",
    disabled: isAddSon
  }
]);

const original = () => {
  return {
    ID: 0,
    authoritys: null,
    component: "LAYOUT",
    hidden: false,
    meta: {
      title: "Dashboard",
      icon: "DashboardOutlined"
    },
    name: "dashboard",
    parentId: 0,
    path: "/dashboard",
    sort: 0
  };
};

const formParams = reactive(original());

function selectAddMenu(key: string) {
  drawerTitle.value = key === "home" ? "添加顶栏菜单" : `添加子菜单：${treeItemTitle.value}`;
  treeItemK.value = key;
  openCreateDrawer();
}

function openCreateDrawer() {
  const { openDrawer } = createDrawerRef.value;
  openDrawer();
}

function disabled(menuTreeData, treeItemId) {
  menuTreeData.forEach((item) => {
    item.disabled = treeItemId == item.ID;
    if (item.children != null) {
      disabled(item.children, treeItemId);
    }
  });
}

function selectedTree(keys) {
  if (keys.length) {
    const treeItem = getTreeItem(unref(treeData), keys[0]);
    treeItemKey.value = keys;
    treeItemId.value = treeItem.ID;
    treeItemTitle.value = treeItem.meta.title;
    disabled(menuTreeData.value, treeItemId.value);
    Object.assign(formParams, treeItem);
    isEditMenu.value = true;
  } else {
    isEditMenu.value = false;
    treeItemKey.value = [];
    treeItemTitle.value = "";
  }
}

async function reloadForm() {
  await loadMenuList();
  isEditMenu.value = false;
}

function handleDelete() {
  const treeItem = getTreeItem(unref(treeData), treeItemKey.value[0]);
  dialog.info({
    title: "提示",
    content: `您想删除菜单: ${treeItem.name}`,
    positiveText: "确定",
    negativeText: "取消",
    onPositiveClick: async () => {
      const { code, message: msg } = await deleteMenu({
        id: treeItem.ID
      });
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("删除成功");
        await reloadForm();
      } else {
        message.error(msg || "删除失败");
      }
    },
    onNegativeClick: () => {
    }
  });
}

function formSubmit() {
  formRef.value.validate(async (errors: boolean) => {
    if (!errors) {
      const { code, message: msg } = await updateMenu(formParams);
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("修改成功");
      } else {
        message.error(msg || "修改失败");
      }
    } else {
      message.error("请填写完整信息");
    }
  });
}

function packHandle() {
  if (expandedKeys.value.length) {
    expandedKeys.value = [];
  } else {
    expandedKeys.value = unref(treeData).map((item: any) => item.key as string) as [];
  }
}

async function loadMenuList() {
  const { code, message: msg, result } = await getMenuList();
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    let treeMenuList = result;
    getMenuTreeTitle(treeMenuList.list);
    const keys = treeMenuList.list.map((item) => item.key);
    Object.assign(formParams, keys);
    treeData.value = treeMenuList.list;
    let tree = treeMenuList.list.concat();
    tree.unshift({
      label: "根目录",
      key: "home",
      ID: 0
    });
    menuTreeData.value = tree;
  } else {
    message.error(msg || "获取失败");
  }
  loading.value = false;
}

onMounted(async () => {
  await loadMenuList();
});

function onExpandedKeys(keys) {
  expandedKeys.value = keys;
}
</script>
