<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-07 03:02:16
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 16:12:46
 * @FilePath: /pinkmoe_admin/src/views/setting/basic/basic.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <div>
    <n-card :bordered="false" class="mt-4 proCard">
      <div class="BasicForm">
        <n-form
          ref="formRef"
          :model="formParams"
          :rules="rules"
          label-placement="left"
          label-width="auto"
          require-mark-placement="right-hanging"
          size="medium"
          :style="{
            maxWidth: '640px',
          }"
        >
          <n-form-item path="url" label="网站url">
            <n-input v-model:value="formParams.url" />
          </n-form-item>
          <n-form-item path="title" label="网站标题">
            <n-input v-model:value="formParams.title" />
          </n-form-item>
          <n-form-item path="keyword" label="网站关键词">
            <n-input v-model:value="formParams.keyword" />
          </n-form-item>
          <n-form-item path="desc" label="网站描述">
            <n-input type="textarea" v-model:value="formParams.desc" />
          </n-form-item>
          <n-form-item path="color" label="网站默认颜色">
            <n-color-picker v-model:value="formParams.color" />
          </n-form-item>
          <n-form-item path="icon" label="网站icon">
            <BasicUpload
              :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
              :headers="uploadHeaders"
              :data="{'post_id': null, 'uuid': userMeta.uuid, 'type': 'settings'}"
              max-size="1"
              max-number="1"
              v-model:value="formParams.icon"
              @uploadChange="uploadIconChange"
            />
          </n-form-item>
          <n-form-item path="background" label="网站背景">
            <BasicUpload
              :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
              :headers="uploadHeaders"
              :data="{'post_id': null, 'uuid': userMeta.uuid, 'type': 'settings'}"
              max-size="1"
              max-number="1"
              v-model:value="formParams.background"
              @uploadChange="uploadBackgroundChange"
            />
          </n-form-item>
          <n-form-item path="logo" label="网站logo">
            <BasicUpload
              :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
              :headers="uploadHeaders"
              :data="{'post_id': null, 'uuid': userMeta.uuid, 'type': 'settings'}"
              max-size="1"
              max-number="1"
              v-model:value="formParams.logo"
              @uploadChange="uploadLogoChange"
            />
          </n-form-item>
          <n-form-item path="createTime" label="网站建立日期">
            <n-date-picker
              v-model:formatted-value="formParams.createTime"
              value-format="yyyy.MM.dd HH:mm:ss"
              type="datetime"
              clearable
            />
          </n-form-item>
          <n-row :gutter="[0, 24]">
            <n-col :span="24">
              <div style="display: flex; justify-content: flex-end">
                <n-button type="primary" @click="handleSubmit"> 提交</n-button>
              </div>
            </n-col>
          </n-row>
        </n-form>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import { NColorPicker, useMessage } from "naive-ui";
import { BasicUpload } from "@/components/Upload";
import { useGlobSetting } from "@/hooks/setting";
import { computed, onMounted, reactive, ref } from "vue";
import { ACCESS_TOKEN, CURRENT_USER } from "@/store/mutation-types";
import { createStorage } from "@/utils/Storage";
import { ResultEnum } from "@/enums/httpEnum";
import { getSetting, updateSetting } from "@/api/setting/setting";

const globSetting = useGlobSetting();
const { uploadUrl } = globSetting;
const Storage = createStorage({ storage: localStorage });
const uploadHeaders = reactive({
  platform: "miniPrograms",
  timestamp: new Date().getTime(),
  Authorization: "Bearer " + Storage.get(ACCESS_TOKEN, "")
});

const rules = {
  url: {
    required: false,
    trigger: "blur",
    message: "请输入2到20位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/ //验证菜单名称 2到20位（字母,汉字,数字,下划线）
  },
  title: {
    required: true,
    trigger: "blur",
    message: "请输入2到20位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/ //验证菜单名称 2到20位（字母,汉字,数字,下划线）
  },
  keyword: {
    required: true,
    trigger: "blur",
    message: "请输入2到50位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,50}$/ //验证菜单名称 2到50位（字母,汉字,数字,下划线）
  },
  desc: {
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,100}$/ //验证菜单名称 2到100位（字母,汉字,数字,下划线）
  },
  color: {
    required: true,
    trigger: "blur",
    message: "请输入颜色",
    pattern: /^[rgba0-9 ,()_]{2,30}$/ //验证颜色
  },
  icon: {
    required: true,
    trigger: "blur",
    message: "请选择正确的图片",
    pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/
  },
  logo: {
    required: true,
    trigger: "blur",
    message: "请选择正确的图片",
    pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/
  },
  createTime: {
    required: false,
    trigger: "blur",
    message: "请输入时间",
  }
};

let formParams = ref({
  url: "",
  title: "",
  keyword: "",
  desc: "",
  color: "",
  icon: [] as string[],
  background: [] as string[],
  logo: [] as string[],
  createTime: "2007.06.30 12:08:55"
});

const message = useMessage();

function uploadIconChange(list: string[]) {
  formParams.value.icon = list;
}

function uploadBackgroundChange(list: string[]) {
  formParams.value.background = list;
}

function uploadLogoChange(list: string[]) {
  formParams.value.logo = list;
}

const formRef: any = ref(null);

async function handleSubmit() {
  formRef.value.validate(async (errors) => {
    if (!errors) {
      const v = JSON.stringify(formParams.value);
      const { code, message: msg } = await updateSetting({
        slug: "website_basic",
        value: v
      });
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("更新成功");
      } else {
        message.error(msg || "更新失败");
      }
    } else {
      message.error("请填写完整信息");
    }
  });
}

async function getSettings() {
  const {
    code,
    message: msg,
    result
  } = await getSetting({
    slug: "website_basic"
  });
  message.destroyAll();
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  formParams.value = JSON.parse(result.value);
}

const userMeta = computed(() => Storage.get(CURRENT_USER));

onMounted(() => {
  getSettings();
});
</script>

<style lang="less" scoped>
.BasicForm {
  width: 950px;
  margin: 0 auto;
  overflow: hidden;
  padding-top: 20px;
}
</style>
