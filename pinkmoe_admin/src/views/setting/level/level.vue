<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-28 19:02:50
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:20
 * @FilePath: /pinkmoe_admin/src/views/setting/level/level.vue
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
          <n-dynamic-input v-model:value="formParams" :on-create="onCreate">
            <template #create-button-default>
              添加等级配置
            </template>
            <template #default="{ value }">
              <n-input v-model:value="value.levelName"
                       style="width: 120px" type="text" placeholder="等级名称" />
              <div style="display: flex; align-items: center; width: 100%">
                <n-input-number
                  v-model:value="value.headExp"
                  style="margin: 0 12px; width: 120px"
                  placeholder="等级起始经验"
                />
                <n-input-number
                  v-model:value="value.footExp"
                  style="width: 120px;margin-right: 12px"
                  placeholder="等级结束经验"
                />
                <n-color-picker
                  v-model:value="value.color"
                  style="width: 120px" />
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
      levelName: '',
      color: '',
      headExp: 0,
      footExp: 0
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
      slug: 'user_level',
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
      slug: 'user_level',
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
    width: 550px;
    margin: 0 auto;
    overflow: hidden;
    padding-top: 20px;
  }
</style>
