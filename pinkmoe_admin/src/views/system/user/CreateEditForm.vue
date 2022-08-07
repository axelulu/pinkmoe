<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-01 20:32:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:26:04
 * @FilePath: /pinkmoe_admin/src/views/system/user/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-modal
    v-model:show="showForm"
    :show-icon="false"
    style="width: 600px"
    preset="dialog"
    :title="formType === 'create' ? '新建' : '更新'"
  >
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="80"
      class="py-4"
    >
      <n-form-item label="名称" path="userName">
        <n-input placeholder="请输入名称" v-model:value="formParams.userName" />
      </n-form-item>
      <n-form-item label="昵称" path="nick_name">
        <n-input placeholder="请输入昵称" v-model:value="formParams.nickName" />
      </n-form-item>
      <n-form-item label="性别" path="sex">
        <n-radio-group v-model:value="formParams.sex" name="sex">
          <n-radio value="3"> 男</n-radio>
          <n-radio value="2"> 女</n-radio>
          <n-radio value="1"> 保密</n-radio>
        </n-radio-group>
      </n-form-item>
      <n-form-item label="电话号码" path="phone">
        <n-input-group>
          <n-button>+86</n-button>
          <n-input placeholder="请输入电话号码" v-model:value="formParams.phone" />
        </n-input-group>
      </n-form-item>
      <n-form-item label="邮箱" path="email">
        <n-auto-complete
          v-model:value="formParams.email"
          :input-props="{
            autocomplete: 'disabled',
          }"
          :options="emailOptions"
          placeholder="请输入邮箱"
        />
      </n-form-item>
      <n-form-item label="头像" path="avatar">
        <BasicUpload
          :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
          :headers="uploadHeaders"
          :data="{'post_id': null, 'uuid': userMeta.uuid, 'type': 'user'}"
          v-model:value="avatar"
          max-size="1"
          max-number="1"
          @uploadChange="uploadAvatarChange"
        />
      </n-form-item>
      <n-form-item label="背景" path="header_img">
        <BasicUpload
          :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
          :headers="uploadHeaders"
          :data="{'post_id': null, 'uuid': userMeta.uuid, 'type': 'user'}"
          v-model:value="headerImg"
          max-size="1"
          max-number="1"
          @uploadChange="uploadHeaderImgChange"
        />
      </n-form-item>
      <n-form-item label="现金" path="cash">
        <n-input-number
          min="0"
          max="100000000"
          placeholder="请输入现金"
          v-model:value="formParams.cash"
        >
          <template #prefix>
            ￥
          </template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="积分" path="credit">
        <n-input-number min="0" placeholder="请输入积分" v-model:value="formParams.credit" />
      </n-form-item>
      <n-form-item label="经验" path="exp">
        <n-input-number
          min="0"
          max="100000000"
          placeholder="请输入经验"
          v-model:value="formParams.exp"
        />
      </n-form-item>
      <n-form-item label="角色" path="authorityId">
        <n-tree-select
          v-model:value="formParams.authorityId"
          key-field="authorityId"
          label-field="authorityName"
          :options="authorityData"
        />
      </n-form-item>
      <n-form-item label="密码" path="password">
        <n-input
          show-password-on="click"
          type="password"
          placeholder="请输入密码"
          v-model:value="formParams.password"
        />
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
import { ACCESS_TOKEN, CURRENT_USER } from "@/store/mutation-types";
import { createUser, updateUser } from "@/api/system/user";
import { ResultEnum } from "@/enums/httpEnum";
import { useGlobSetting } from "@/hooks/setting";
import { BasicUpload } from "@/components/Upload";
import { createStorage } from "@/utils/Storage";
import { getRoleList } from "@/api/system/authority";

const rules = {
  userName: {
    required: true,
    message: "4到16位（字母，数字，下划线，减号）",
    trigger: "blur",
    pattern: /^[a-zA-Z0-9_-]{4,16}$/ //验证用户名 4到16位（字母，数字，下划线，减号）
  },
  password: {
    required: true,
    trigger: "blur",
    message: "6-20位，包含大小写字母和数字",
    pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[\s\S]{6,20}$/ //验证密码 6-20位，包含大小写字母和数字
  },
  nickName: {
    required: false,
    message: "请输入2到16位昵称（字母,汉字,数字,下划线）",
    trigger: "blur",
    pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,16}$/ //验证昵称 2到16位（字母,汉字,数字,下划线）
  },
  avatar: {
    required: false,
    trigger: "blur",
    message: "请选择正确的用户头像",
    pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/
  },
  headerImg: {
    required: false,
    trigger: "blur",
    message: "请选择正确的背景图片",
    pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/
  },
  cash: {
    type: "number",
    required: true,
    trigger: "blur",
    message: "请填写用户现金",
    pattern: /^[0-9]{11}$/ //验证用户现金
  },
  credit: {
    type: "number",
    required: true,
    trigger: "blur",
    message: "请填写用户积分",
    pattern: /^[0-9]{11}$/ //验证用户积分
  },
  exp: {
    type: "number",
    required: true,
    trigger: "blur",
    message: "请填写用户经验",
    pattern: /^[0-9]{11}$/ //验证用户经验
  },
  phone: {
    required: false,
    message: "请输入正确的手机号",
    trigger: "blur",
    pattern: /^[0-9]{11}$/ //验证手机号 11位手机号
  },
  email: {
    type: "email",
    required: true,
    trigger: "blur",
    message: "请填写正确邮箱"
  },
  authorityId: {
    required: true,
    trigger: "blur",
    message: "请选择角色"
  },
  sex: {
    type: "enum",
    enum: ["1", "2", "3"],
    required: true,
    trigger: "blur",
    message: "请选择性别"
  }
};
const Storage = createStorage({ storage: localStorage });
const globSetting = useGlobSetting();

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
    const { uploadUrl } = globSetting;
    const formBtnLoading = ref(false);
    const formRef: any = ref(null);
    const authorityData = ref([]);
    const params = ref({
      pageSize: 999
    });
    const defaultValueRef = () => ({
      userName: "",
      password: "",
      nickName: "",
      avatar: "",
      headerImg: "",
      cash: 0,
      credit: 0,
      exp: 0,
      phone: "",
      email: "",
      authorityId: "",
      sex: "0"
    });

    const state = reactive({
      showForm: false,
      avatar: [] as string[],
      headerImg: [] as string[],
      formParams: defaultValueRef()
    });

    watch(
      () => props.formType,
      () => {
        if (props.formType == "edit") {
          if (state.formParams.avatar) {
            state.avatar[0] = state.formParams.avatar;
          }
          if (state.formParams.headerImg) {
            state.headerImg[0] = state.formParams.headerImg;
          }
        } else {
          state.avatar = [];
          state.headerImg = [];
        }
      }
    );

    async function getRoleLists() {
      const { code, message: msg, result } = await getRoleList(params.value);
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        authorityData.value = result.list;
      } else {
        message.error(msg || "获取失败");
      }
    }

    getRoleLists();

    function uploadHeaderImgChange(list: string[]) {
      state.headerImg = list;
    }

    function uploadAvatarChange(list: string[]) {
      state.avatar = list;
    }

    const uploadHeaders = reactive({
      platform: "miniPrograms",
      timestamp: new Date().getTime(),
      Authorization: "Bearer " + Storage.get(ACCESS_TOKEN, "")
    });

    function confirmForm(e) {
      e.preventDefault();
      formBtnLoading.value = true;
      formRef.value.validate(async (errors) => {
        if (!errors) {
          state.formParams.avatar = state.avatar[0];
          state.formParams.headerImg = state.headerImg[0];
          const { code, message: msg } =
            props.formType === "create"
              ? await createUser(state.formParams)
              : await updateUser(state.formParams);
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

    function openForm(record) {
      if (record !== null) {
        state.formParams = record;
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
      uploadUrl,
      authorityData,
      formBtnLoading,
      uploadHeaders,
      openForm,
      closeForm,
      uploadAvatarChange,
      uploadHeaderImgChange,
      confirmForm,
      userMeta,
      emailOptions: computed(() => {
        return ["@gmail.com", "@163.com", "@qq.com"].map((suffix) => {
          const prefix = state.formParams.email.split("@")[0];
          return {
            label: prefix + suffix,
            value: prefix + suffix
          };
        });
      })
    };
  }
});
</script>
