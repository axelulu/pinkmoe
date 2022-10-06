<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 16:19:43
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-17 21:05:35
 * @FilePath: /pinkmoe_admin/src/views/setting/shopMod/shopMod.vue
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
            maxWidth: '1040px',
          }"
        >
          <n-form-item path="popular" label="首页幻灯片商品">
            <n-select
              v-model:value="formParams.popular"
              :render-label="renderPostLabel"
              :render-tag="renderPostSelectTag"
              :options="postList"
              clearable
              filterable
              multiple
              tag
            />
          </n-form-item>
          <n-form-item path="desc" label="首页cms">
            <n-dynamic-input v-model:value="formParams.cms" show-sort-button
                             :on-create="onCreate">
              <template #create-button-default> 请选择cms内容</template>
              <template #default="{ value }">
                <div
                  style="display: flex; flex-direction: column; align-items: center; width: 100%">
                  <div style="display: flex; align-items: center; width: 100%">
                    <n-cascader
                      :virtual-scroll="true"
                      cascade="false"
                      label-field="name"
                      value-field="slug"
                      v-model:value="value.category"
                      :options="category"
                    />
                  </div>
                  <div style="display: flex;flex-direction: row; align-items: center; width: 100%">
                    <div @click="value.style = 1"
                         :style="value.style === 1 ? 'box-shadow: 0 0 0 2px #ff6699;' : ''"
                         style="margin-left: 2%; margin-right: 2%; margin-top: 2%; width: 100%;cursor: pointer">
                      <n-skeleton :animated="false" text width="100%" height="150px" :repeat="1"
                                  style="width: 100%;" />
                      <n-skeleton :animated="false" text width="100%" height="20px" :repeat="1"
                                  style="width: 100%;" />
                      <n-skeleton :animated="false" text width="31%" height="20px" :repeat="3"
                                  style="width: 31%;" />
                    </div>
                    <div @click="value.style = 2"
                         :style="value.style === 2 ? 'box-shadow: 0 0 0 2px #ff6699;' : ''"
                         style="margin: 2% 2% 2% 10px;width: 100%;cursor: pointer;">
                      <n-skeleton :animated="false" text width="100%" height="80px" :repeat="1"
                                  style="width: 100%;" />
                      <n-skeleton :animated="false" text width="100%" height="20px" :repeat="1"
                                  style="width: 100%;" />
                      <n-skeleton :animated="false" text width="31%" height="20px" :repeat="3"
                                  style="width: 31%;" />
                    </div>
                  </div>
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
import {
  NAvatar,
  NText,
  SelectRenderLabel,
  SelectRenderTag,
  useMessage
} from "naive-ui";
import { h, onMounted, ref } from "vue";
import { ResultEnum } from "@/enums/httpEnum";
import { getSetting, updateSetting } from "@/api/setting/setting";
import { getGoodsList } from "@/api/shop/goods";
import { getShopCategoryList } from "@/api/shop/category";

const postList = ref([]);

const rules = {
  recommend: {
    type: "array",
    required: true,
    trigger: "blur",
    message: "请输入2到20位（字母,汉字,数字,下划线）"
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/ //验证菜单名称 2到20位（字母,汉字,数字,下划线）
  },
  popular: {
    type: "array",
    required: true,
    trigger: "blur",
    message: "请输入2到50位（字母,汉字,数字,下划线）"
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,50}$/ //验证菜单名称 2到50位（字母,汉字,数字,下划线）
  },
  cms: {
    type: "array",
    required: true,
    trigger: "blur",
    message: "请输入2到100位（字母,汉字,数字,下划线）"
    // pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,100}$/ //验证菜单名称 2到100位（字母,汉字,数字,下划线）
  }
};

let formParams = ref({
  popular: [] as string[],
  recommend: [] as string[],
  cms: [] as string[]
});

const message = useMessage();
const category = ref([]);

const renderPostSelectTag: SelectRenderTag = ({ option }) => {
  return h(
    "div",
    {
      style: {
        display: "flex",
        alignItems: "center"
      }
    },
    [
      h(NAvatar, {
        src: option.cover as string,
        round: true,
        objectFit: "cover",
        size: 24,
        style: {
          marginRight: "12px"
        }
      }),
      option.title as string
    ]
  );
};

const renderPostLabel: SelectRenderLabel = (option) => {
  return h(
    "div",
    {
      style: {
        display: "flex",
        alignItems: "center"
      }
    },
    [
      h(NAvatar, {
        src: option.cover as string,
        round: true,
        objectFit: "cover",
        size: "small"
      }),
      h(
        "div",
        {
          style: {
            marginLeft: "12px",
            padding: "4px 0"
          }
        },
        [
          h("div", null, [option.title as string]),
          h(
            NText,
            { depth: 3, tag: "div" },
            {
              default: () => option.exerpt as string
            }
          )
        ]
      )
    ]
  );
};

function onCreate() {
  return {
    category: "",
    style: ""
  };
}

async function getGoodsCategoryLists() {
  const {
    code,
    message: msg,
    result
  } = await getShopCategoryList({
    pageSize: 999,
    page: 1
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    category.value = result.list;
  } else {
    message.error(msg || "获取失败");
  }
}

async function getGoodsLists() {
  const {
    code,
    message: msg,
    result
  } = await getGoodsList({
    pageSize: 999,
    page: 1
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    result.list.forEach((item) => {
      item.label = item.title;
      item.value = item.goodsId;
    });
    postList.value = result.list;
    console.log(postList.value);
  } else {
    message.error(msg || "获取失败");
  }
}

async function handleSubmit() {
  const v = JSON.stringify(formParams.value);
  const { code, message: msg } = await updateSetting({
    slug: "website_shop",
    value: v
  });
  message.destroyAll();
  if (code == ResultEnum.SUCCESS) {
    message.success("更新成功");
  } else {
    message.error(msg || "更新失败");
  }
}

async function getSettings() {
  const {
    code,
    message: msg,
    result
  } = await getSetting({
    slug: "website_shop"
  });
  message.destroyAll();
  if (code != ResultEnum.SUCCESS) {
    message.error(msg || "获取失败");
  }
  formParams.value = JSON.parse(result.value);
}

onMounted(() => {
  getSettings();
  getGoodsLists();
  getGoodsCategoryLists();
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
