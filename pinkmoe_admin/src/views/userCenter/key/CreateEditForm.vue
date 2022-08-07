<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 16:16:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:26:13
 * @FilePath: /pinkmoe_admin/src/views/userCenter/key/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-modal v-model:show="showForm" :show-icon="false" preset="dialog" title="新建">
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="80"
      class="py-4"
    >
      <n-form-item label="面额" path="value">
        <n-input-number placeholder="请输入面额" v-model:value="formParams.value" />
      </n-form-item>
      <n-form-item label="类型" path="type">
        <n-select v-model:value="formParams.type" :options="keyTypes" />
      </n-form-item>
      <n-form-item label="生成数量" path="type">
        <n-input-number placeholder="请输入生成数量" v-model:value="formParams.num" />
      </n-form-item>
    </n-form>

    <template #action>
      <n-space>
        <n-button @click="() => (showForm = false)">取消</n-button>
        <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
      </n-space>
    </template>
  </n-modal>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, toRefs, watch } from "vue";
import { useMessage } from "naive-ui";
import { ResultEnum } from "@/enums/httpEnum";
import { createKey, updateKey } from "@/api/userCenter/key";
import { useGlobSetting } from "@/hooks/setting";
import { createStorage } from "@/utils/Storage";
import { ACCESS_TOKEN, CURRENT_USER } from "@/store/mutation-types";
import { BasicUpload } from "@/components/Upload";

const globSetting = useGlobSetting();
const { uploadUrl } = globSetting;
const Storage = createStorage({ storage: localStorage });

const rules = {
  label: {
    required: true,
    trigger: "blur",
    message: "请输入2到16位（字母,汉字,数字,下划线）",
    pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,30}$/ //验证菜单名称 2到16位（字母,汉字,数字,下划线）
  },
  value: {
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母）",
    pattern: /^[0-9a-zA-Z/_-]{2,100}$/ //验证菜单路径 2到100位（字母）
  },
  icon: {
    required: false,
    trigger: "blur",
    message: "请选择正确的图标",
    pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/
  },
  sort: {
    required: true,
    trigger: "blur",
    message: "请输入1-2位数字",
    pattern: /^[0-9]{1,2}$/ //验证排序内容 1-2位数字
  }
};

export default defineComponent({
  name: "CreateEditForm",
  components: {
    BasicUpload
  },
  props: {
    formType: {
      type: String,
      default: "create"
    }
  },
  setup(props, context) {
    const message = useMessage();
    const formBtnLoading = ref(false);
    const formRef: any = ref(null);
    const options = [];
    const defaultValueRef = () => ({
      ID: 0,
      key: "",
      value: 10,
      num: 10,
      type: "credit",
      status: "1"
    });
    const keyTypes = [
      {
        label: "积分",
        value: 'credit',
      },
      {
        label: "现金",
        value: 'cash',
      },
    ]

    const uploadHeaders = reactive({
      platform: "miniPrograms",
      timestamp: new Date().getTime(),
      Authorization: "Bearer " + Storage.get(ACCESS_TOKEN, "")
    });

    const state = reactive({
      showForm: false,
      icon: [] as string[],
      formParams: defaultValueRef()
    });

    function confirmForm(e) {
      e.preventDefault();
      formBtnLoading.value = true;
      formRef.value.validate(async (errors) => {
        if (!errors) {
          state.formParams.icon = state.icon[0];
          const { code, message: msg } =
            props.formType === "create"
              ? await createKey(state.formParams)
              : await updateKey(state.formParams);
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            message.success(props.formType === "create" ? "新建成功" : "更新成功");
            setTimeout(() => {
              state.showForm = false;
              closeForm();
              context.emit("reloadTable");
            });
          } else {
            message.error(msg || (props.formType === "create" ? "新建失败" : "更新失败"));
          }
        } else {
          message.error("请填写完整信息");
        }
        formBtnLoading.value = false;
      });
    }

    watch(
      () => props.formType,
      () => {
        if (props.formType == "create") {
          state.icon = [];
        }
      }
    );

    function openForm(record) {
      if (record !== null) {
        state.formParams = record;
      }
      if (state.formParams.icon) {
        state.icon[0] = state.formParams.icon;
      }
      state.showForm = true;
    }

    function closeForm() {
      state.showForm = false;
    }

    const userMeta = computed(() => Storage.get(CURRENT_USER));

    return {
      ...toRefs(state),
      formRef,
      rules,
      options,
      formBtnLoading,
      uploadHeaders,
      uploadUrl,
      keyTypes,
      openForm,
      closeForm,
      confirmForm,
      userMeta
    };
  }
});
</script>
