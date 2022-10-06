<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-07 22:20:13
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:49
 * @FilePath: /pinkmoe_admin/src/views/system/menu/CreateDrawer.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <n-drawer v-model:show="isDrawer" :width="width" :placement="placement">
    <n-drawer-content :title="title" closable>
      <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="100"
      >
        <n-form-item label="类型" path="type">
          <span>{{ !formParams.hidden ? "侧边栏菜单" : "" }}</span>
        </n-form-item>
        <n-form-item label="标题" path="title">
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
      </n-form>
      <template #footer>
        <n-space>
          <n-button type="primary" :loading="subLoading" @click="formSubmit">提交</n-button>
          <n-button @click="handleReset">重置</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, toRefs, watchEffect } from "vue";
import { useMessage } from "naive-ui";
import { ResultEnum } from "@/enums/httpEnum";
import { createMenu } from "@/api/system/menu";

const rules = {
  label: {
    required: true,
    message: "请输入标题",
    trigger: "blur",
    pattern: /^[a-zA-Z]{2,16}$/ //验证用户名 4到16位（字母）
  },
  name: {
    required: true,
    message: "请输入标识",
    trigger: "blur",
    pattern: /^[a-zA-Z_-]{2,16}$/ //验证用户名 4到16位（字母）
  },
  path: {
    required: true,
    message: "请输入路径",
    trigger: "blur",
    pattern: /^[a-zA-Z/]{2,16}$/ //验证用户名 4到16位（字母）
  },
  component: {
    required: true,
    message: "请输入文件路径",
    trigger: "blur",
    pattern: /^[a-zA-Z/]{2,100}$/ //验证用户名 4到16位（字母）
  },
  sort: {
    required: false,
    message: "请输入排序",
    trigger: "blur",
    pattern: /^[0-9]{1,2}$/ //验证验证码 1-2位数字
  },
  icon: {
    required: true,
    message: "请输入图标",
    trigger: "blur",
    pattern: /^[a-zA-Z]{2,16}$/ //验证用户名 4到16位（字母）
  },
  parentId: {
    required: true,
    message: "请输入父目录",
    trigger: "blur",
    pattern: /^[0-9]{1,2}$/ //验证验证码 1-2位数字
  },
  hidden: {
    required: true,
    message: "请选择显示状态",
    trigger: "blur",
    pattern: /^[a-zA-Z]{0,6}$/ //验证用户名 4到5位
  }
};
export default defineComponent({
  name: "CreateDrawer",
  components: {},
  props: {
    title: {
      type: String,
      default: "添加顶级菜单"
    },
    width: {
      type: Number,
      default: 450
    },
    menuTreeData: {
      type: Array,
      default: null
    },
    treeItemId: {
      type: Number,
      default: 0
    },
    treeItemK: {
      type: String
    }
  },
  setup(props, context) {
    const message = useMessage();
    const formRef: any = ref(null);
    const defaultValueRef = () => ({
      authoritys: null,
      component: "",
      hidden: false,
      meta: {
        title: "",
        icon: ""
      },
      name: "",
      parentId: 0,
      path: "",
      sort: 0
    });

    const state = reactive({
      isDrawer: false,
      formParams: defaultValueRef()
    });

    watchEffect(() => {
      state.formParams.parentId = props.treeItemK === "home" ? 0 : props.treeItemId;
    });

    function openDrawer() {
      state.isDrawer = true;
    }

    function closeDrawer() {
      state.isDrawer = false;
      context.emit("reloadForm");
    }

    function formSubmit() {
      formRef.value.validate(async (errors) => {
        if (!errors) {
          const { code, message: msg } = await createMenu(state.formParams);
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            message.success("添加成功");
          } else {
            message.error(msg || "添加失败");
          }
          handleReset();
          closeDrawer();
        } else {
          message.error("请填写完整信息");
        }
      });
    }

    function handleReset() {
      formRef.value.restoreValidation();
      state.formParams = Object.assign(state.formParams, defaultValueRef());
    }

    return {
      ...toRefs(state),
      formRef,
      rules,
      formSubmit,
      handleReset,
      openDrawer,
      closeDrawer
    };
  }
});
</script>
