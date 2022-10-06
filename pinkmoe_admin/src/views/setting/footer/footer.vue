<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 16:20:06
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-21 17:08:27
 * @FilePath: /pinkmoe_admin/src/views/setting/footer/footer.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
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
          <n-form-item label="链接导航">
            <n-dynamic-input v-model:value="formParams.links" show-sort-button
                             :on-create="onCreate">
              <template #create-button-default> 请填写音乐信息</template>
              <template #default="{ value }">
                <div style="display: flex; align-items: center; width: 100%">
                  <n-input
                    maxlength="30"
                    show-count
                    placeholder="请输入链接名称"
                    style="margin-right: 12px"
                    v-model:value="value.name"
                    type="text"
                  />
                  <n-input
                    placeholder="请输入链接"
                    :input-props="{ type: 'url' }"
                    v-model:value="value.url"
                    type="text"
                  >
                    <template #suffix>
                      <n-icon :component="GlobeOutline" />
                    </template>
                  </n-input>
                </div>
              </template>
            </n-dynamic-input>
          </n-form-item>
          <n-form-item path="notice" label="网站申明">
            <n-input type="textarea" v-model:value="formParams.notice" />
          </n-form-item>
          <n-form-item path="contact" label="联系我们">
            <n-input type="textarea" v-model:value="formParams.contact" />
          </n-form-item>
          <n-form-item path="about" label="关于小站">
            <n-input type="textarea" v-model:value="formParams.about" />
          </n-form-item>
          <n-form-item label="链接导航">
            <n-dynamic-input v-model:value="formParams.friendsLinks" show-sort-button
                             :on-create="onCreate">
              <template #create-button-default> 请填写音乐信息</template>
              <template #default="{ value }">
                <div style="display: flex; align-items: center; width: 100%">
                  <n-input
                    maxlength="30"
                    show-count
                    placeholder="请输入链接名称"
                    style="margin-right: 12px"
                    v-model:value="value.name"
                    type="text"
                  />
                  <n-input
                    placeholder="请输入链接"
                    :input-props="{ type: 'url' }"
                    v-model:value="value.url"
                    type="text"
                  >
                    <template #suffix>
                      <n-icon :component="GlobeOutline" />
                    </template>
                  </n-input>
                </div>
              </template>
            </n-dynamic-input>
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
import { useMessage } from "naive-ui";
import { onMounted, ref } from "vue";
import { ResultEnum } from "@/enums/httpEnum";
import { getSetting, updateSetting } from "@/api/setting/setting";
import { GlobeOutline } from "@vicons/ionicons5";
import { array } from "vue-types";

const rules = {
  links: {
    required: true,
    type: array,
    trigger: "blur",
    message: "请输入正确的值"
  },
  notice: {
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9.\-,:。 \n，!：；、_\u4e00-\u9fa5]{2,100}$/ //验证菜单名称 2到100位（字母,汉字,数字,下划线）
  },
  contact: {
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9.\-,:。 \n，!：；、_\u4e00-\u9fa5]{2,100}$/ //验证菜单名称 2到100位（字母,汉字,数字,下划线）
  },
  about: {
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母,汉字,数字,下划线）",
    // pattern: /^[a-zA-Z0-9.\-,:。 \n，!：；、_\u4e00-\u9fa5]{2,100}$/ //验证菜单名称 2到100位（字母,汉字,数字,下划线）
  },
  friendsLinks: {
    required: true,
    type: array,
    trigger: "blur",
    message: "请输入正确的值"
  }
};

let formParams = ref({
  links: {
    name: "",
    url: ""
  },
  notice: "",
  contact: "",
  about: "",
  friendsLinks: {
    name: "",
    url: ""
  }
});

const message = useMessage();

const formRef: any = ref(null);

async function handleSubmit() {
  formRef.value.validate(async (errors) => {
    if (!errors) {
      const v = JSON.stringify(formParams.value);
      const { code, message: msg } = await updateSetting({
        slug: "website_footer",
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

function onCreate() {
  return {
    name: "",
    url: ""
  };
}

async function getSettings() {
  const {
    code,
    message: msg,
    result
  } = await getSetting({
    slug: "website_footer"
  });
  message.destroyAll();
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  formParams.value = JSON.parse(result.value);
}

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
