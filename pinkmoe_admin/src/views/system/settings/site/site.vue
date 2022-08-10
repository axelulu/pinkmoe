<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 22:25:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-10 09:40:07
 * @FilePath: /pinkmoe_admin/src/views/system/settings/site/site.vue
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
          <n-form-item path="name" label="server的名称">
            <n-input v-model:value="formParams.name" />
          </n-form-item>
          <n-form-item path="version" label="server的版本">
            <n-input v-model:value="formParams.version" />
          </n-form-item>
          <n-form-item path="mode" label="server启动模式">
            <n-input v-model:value="formParams.mode" />
          </n-form-item>
          <n-form-item path="fileVideoSize" label="server的api端口">
            <n-input-number v-model:value="formParams.port" />
          </n-form-item>
          <n-form-item path="fileVideoSize" label="文件上传大小">
            <n-input-number v-model:value="formParams.videoSize" />
          </n-form-item>
          <n-form-item path="filePicSize" label="文件上传数量">
            <n-input-number v-model:value="formParams.picSize" />
          </n-form-item>
          <n-form-item path="snowStartTime" label="(生成用户ID)雪花算法起始时间">
            <n-input v-model:value="formParams.startTime" />
          </n-form-item>
          <n-form-item path="snowMachineId" label="雪花算法机器ID">
            <n-input-number v-model:value="formParams.machineId" />
          </n-form-item>
          <n-form-item path="rateLimitTime" label="后台流量限制次数">
            <n-input-number v-model:value="formParams.rateLimitTime" />
          </n-form-item>
          <n-form-item path="rateLimitNum" label="流量限制每次数量">
            <n-input-number v-model:value="formParams.rateLimitNum" />
          </n-form-item>
          <n-form-item path="captchaHeight" label="图片验证码长度">
            <n-input-number v-model:value="formParams.captcha.imgHeight" />
          </n-form-item>
          <n-form-item path="captchaWidth" label="图片验证码宽度">
            <n-input-number v-model:value="formParams.captcha.imgWidth" />
          </n-form-item>
          <n-form-item path="captchaLong" label="图片验证码位数">
            <n-input-number v-model:value="formParams.captcha.keyLong" />
          </n-form-item>
          <n-form-item path="rbac" label="rbac权限控制配置文件">
            <n-input v-model:value="formParams.casbin.modelPath" />
          </n-form-item>
          <n-form-item path="jwtExpire" label="登陆token过期时间(小时)">
            <n-input-number v-model:value="formParams.authConfig.jwtExpire" />
          </n-form-item>
          <n-form-item path="jwtIssuer" label="token的issuer">
            <n-input v-model:value="formParams.authConfig.issuer" />
          </n-form-item>
          <n-form-item path="jwtAuthSecret" label="token的secret">
            <n-input v-model:value="formParams.authConfig.authSecret" />
          </n-form-item>
          <n-form-item path="logLevel" label="日志等级">
            <n-input v-model:value="formParams.logConfig.level" />
          </n-form-item>
          <n-form-item path="logFilename" label="日志存储路径">
            <n-input v-model:value="formParams.logConfig.filename" />
          </n-form-item>
          <n-form-item path="logMaxSize" label="日志最大大小(MB)">
            <n-input-number v-model:value="formParams.logConfig.maxSize" />
          </n-form-item>
          <n-form-item path="logMaxAge" label="日志最大存储时间(天)">
            <n-input-number v-model:value="formParams.logConfig.maxAge" />
          </n-form-item>
          <n-form-item path="logMaxBackups" label="日志分桶数量">
            <n-input-number v-model:value="formParams.logConfig.maxBackups" />
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
    name: 'pinkMoe',
    mode: 'dev',
    version: 'v1.0.1',
    port: 9527,
    videoSize: 100,
    picSize: 1,
    startTime: '2006-01-02',
    machineId: 1,
    rateLimitTime: 1,
    rateLimitNum: 10000,
    captcha: {
      imgHeight: 50,
      imgWidth: 200,
      keyLong: 6,
    },
    casbin: {
      modelPath: './util/rbac_model.conf',
    },
    authConfig: {
      jwtExpire: 168,
      issuer: 'PinkMoe',
      authSecret: 'PinkMoe',
    },
    logConfig: {
      level: 'debug',
      filename: 'logs/pinkMoe.log',
      maxSize: 100,
      maxAge: 30,
      maxBackups: 10,
    },
  });

  const message = useMessage();

  async function handleSubmit() {
    const v = JSON.stringify(formParams.value);
    const { code, message: msg } = await updateSetting({
      slug: 'system_site',
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
      slug: 'system_site',
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
