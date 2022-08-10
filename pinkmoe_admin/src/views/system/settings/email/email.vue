<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 22:25:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-10 09:13:37
 * @FilePath: /pinkmoe_admin/src/views/system/settings/email/email.vue
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
          <n-form-item path="user" label="用户名称">
            <n-input v-model:value="formParams.user" />
          </n-form-item>
          <n-form-item path="username" label="邮箱用户名">
            <n-input v-model:value="formParams.username" />
          </n-form-item>
          <n-form-item path="password" label="邮箱密码">
            <n-input v-model:value="formParams.password" />
          </n-form-item>
          <n-form-item path="host" label="邮箱host">
            <n-input v-model:value="formParams.host" />
          </n-form-item>
          <n-form-item path="port" label="邮箱端口">
            <n-input-number v-model:value="formParams.port" />
          </n-form-item>
          <n-form-item path="isSsl" label="是否ssl">
            <n-switch v-model:value="formParams.isSsl" />
          </n-form-item>
          <n-form-item path="replyEmail" label="回复邮箱地址">
            <n-input v-model:value="formParams.replyEmail" />
          </n-form-item>
          <n-row :gutter="[0, 24]">
            <n-col :span="24">
              <div style="display: flex; justify-content: flex-end">
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
  import { useMessage } from 'naive-ui';
  import { onMounted, ref } from 'vue';
  import { ResultEnum } from '@/enums/httpEnum';
  import { getSetting, updateSetting } from '@/api/setting/setting';

  const rules = {
    title: {
      required: true,
      trigger: 'blur',
      message: '请输入2到20位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/, //验证菜单名称 2到20位（字母,汉字,数字,下划线）
    },
  };

  let formParams = ref<any>({
    user: 'izhaicy@163.com',
    username: 'izhaicy@163.com',
    password: 'izhaicy@163.com',
    host: 'smtpdm.aliyun.com',
    port: 465,
    isSsl: true,
    replyEmail: 'izhaicy@163.com',
  });

  const message = useMessage();

  async function handleSubmit() {
    const v = JSON.stringify(formParams.value);
    const { code, message: msg } = await updateSetting({
      slug: 'system_email',
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
      slug: 'system_email',
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
