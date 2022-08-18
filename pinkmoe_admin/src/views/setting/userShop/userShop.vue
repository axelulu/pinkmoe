<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 07:40:19
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:25
 * @FilePath: /pinkmoe_admin/src/views/setting/shop/search.vue
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
        >
          <n-dynamic-input style="width: 100%" v-model:value="formParams" :on-create="onCreate">
            <template #create-button-default>
              添加用户商城配置
            </template>
            <template #default="{ value }">
              <div style="display: flex; align-items: start; width: 100%">
                <div class="w-1/2 pr-4">
                  <n-select v-model:value="value.shopType" :options="[
                {
                  'label': '购买积分',
                  'value': 'credit',
                },
                {
                  'label': '充值现金',
                  'value': 'cash',
                }
              ]" />
                </div>
                <n-input v-model:value="value.shopName"
                         class="mr-4 w-1/2" type="text" placeholder="商品名称" />
                <n-input-number v-model:value="value.shopValue"
                         class="mr-4 w-1/2" type="text" placeholder="商品值" />
                <n-input v-model:value="value.shopDesc"
                         class="mr-4 w-full" type="textarea" placeholder="商品描述" />
                <div class="w-1/2 pr-4">
                  <n-select v-model:value="value.priceType" :options="[
                {
                  'label': '现金',
                  'value': 'cash',
                },{
                  'label': '积分',
                  'value': 'credit',
                },{
                  'label': '卡密',
                  'value': 'key',
                }
              ]" />
                </div>
                <n-input-number
                  v-if="value.priceType === 'cash' || value.priceType === 'credit'"
                  v-model:value="value.shopCredit"
                  class="w-1/2"
                  placeholder="商品价格"
                />
              </div>
            </template>
          </n-dynamic-input>
          <n-row :gutter="[0, 24]">
            <n-col :span="24">
              <div style="display: flex; justify-content: flex-end;margin-top: 20px">
                <n-button type="primary" @click="handleSubmit"> 提交 </n-button>
              </div>
            </n-col>
          </n-row>
        </n-form>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { NColorPicker, useMessage } from 'naive-ui';
  import { BasicUpload } from '@/components/Upload';
  import { useGlobSetting } from '@/hooks/setting';
  import { onMounted, reactive, ref } from 'vue';
  import { ACCESS_TOKEN } from '@/store/mutation-types';
  import { createStorage } from '@/utils/Storage';
  import { ResultEnum } from '@/enums/httpEnum';
  import { getSetting, updateSetting } from '@/api/setting/setting';

  const globSetting = useGlobSetting();
  const { uploadUrl } = globSetting;
  const Storage = createStorage({ storage: localStorage });
  const uploadHeaders = reactive({
    platform: 'miniPrograms',
    timestamp: new Date().getTime(),
    Authorization: 'Bearer ' + Storage.get(ACCESS_TOKEN, ''),
  });
  function onCreate(){
    return {
      shopCredit: '',
      shopDesc: '',
      shopName: '',
    }
  }

  const rules = {
    title: {
      required: true,
      trigger: 'blur',
      message: '请输入2到20位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/, //验证菜单名称 2到20位（字母,汉字,数字,下划线）
    },
  };

  let formParams = ref<any>([]);

  const message = useMessage();

  async function handleSubmit() {
    const v = JSON.stringify(formParams.value);
    const { code, message: msg } = await updateSetting({
      slug: 'user_shop',
      value: v,
    });
    message.destroyAll();
    if (code == ResultEnum.SUCCESS) {
      message.success('更新成功');
    } else {
      message.error(msg || '更新失败');
    }
  }

  async function getSettings() {
    const {
      code,
      message: msg,
      result,
    } = await getSetting({
      slug: 'user_shop',
    });
    message.destroyAll();
    if (code != ResultEnum.SUCCESS) {
      message.error(msg || '获取失败');
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
