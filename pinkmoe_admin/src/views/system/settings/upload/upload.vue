<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 22:25:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-15 15:37:27
 * @FilePath: /pinkmoe_admin/src/views/system/settings/upload/upload.vue
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
          <n-form-item path="ossType" label="文件上传类型">
            <n-select class="w-80" v-model:value="formParams.ossType" :options="options" />
          </n-form-item>
          <div v-if="formParams.ossType === 'local'">
            <n-form-item path="localPath" label="本地存储路径">
              <n-input v-model:value="formParams.localConfig.path" />
            </n-form-item>
          </div>
          <div v-if="formParams.ossType === 'aliyun'">
            <n-form-item path="aliyunEndpoint" label="Endpoint">
              <n-input v-model:value="formParams.aliyunOssConfig.endpoint" />
            </n-form-item>
            <n-form-item path="aliyunAccessKeyId" label="密钥ID">
              <n-input v-model:value="formParams.aliyunOssConfig.accessKeyId" />
            </n-form-item>
            <n-form-item path="aliyunAccessKeySecret" label="密钥secret">
              <n-input v-model:value="formParams.aliyunOssConfig.accessKeySecret" />
            </n-form-item>
            <n-form-item path="aliyunBucketName" label="储存桶">
              <n-input v-model:value="formParams.aliyunOssConfig.bucketName" />
            </n-form-item>
            <n-form-item path="aliyunBucketPoint" label="存储域名">
              <n-input v-model:value="formParams.aliyunOssConfig.bucketPoint" />
            </n-form-item>
          </div>
          <div v-if="formParams.ossType === 'tencent'">
            <n-form-item path="tencentBucket" label="储存桶">
              <n-input v-model:value="formParams.tencentCOSConfig.bucket" />
            </n-form-item>
            <n-form-item path="tencentRegion" label="区域">
              <n-input v-model:value="formParams.tencentCOSConfig.region" />
            </n-form-item>
            <n-form-item path="tencentSecretId" label="密钥ID">
              <n-input v-model:value="formParams.tencentCOSConfig.secretId" />
            </n-form-item>
            <n-form-item path="tencentSecretKey" label="密钥secret">
              <n-input v-model:value="formParams.tencentCOSConfig.secretKey" />
            </n-form-item>
            <n-form-item path="tencentBaseUrl" label="基本链接">
              <n-input v-model:value="formParams.tencentCOSConfig.baseUrl" />
            </n-form-item>
            <n-form-item path="tencentPathPrefix" label="路径prefix">
              <n-input v-model:value="formParams.tencentCOSConfig.pathPrefix" />
            </n-form-item>
          </div>
          <div v-if="formParams.ossType === 'qiniu'">
            <n-form-item path="qiniuZone" label="zone">
              <n-input v-model:value="formParams.qiniuConfig.zone" />
            </n-form-item>
            <n-form-item path="qiniuBucket" label="储存桶">
              <n-input v-model:value="formParams.qiniuConfig.bucket" />
            </n-form-item>
            <n-form-item path="qiniuImgPath" label="存储路径">
              <n-input v-model:value="formParams.qiniuConfig.imgPath" />
            </n-form-item>
            <n-form-item path="qiniuUseHttps" label="使用https">
              <n-input v-model:value="formParams.qiniuConfig.useHttps" />
            </n-form-item>
            <n-form-item path="qiniuAccessKey" label="访问key">
              <n-input v-model:value="formParams.qiniuConfig.accessKey" />
            </n-form-item>
            <n-form-item path="qiniuSecretKey" label="密钥secret">
              <n-input v-model:value="formParams.qiniuConfig.secretKey" />
            </n-form-item>
            <n-form-item path="qiniuUseCdnDomains" label="使用cdn域名">
              <n-input v-model:value="formParams.qiniuConfig.useCdnDomains" />
            </n-form-item>
          </div>
          <div v-if="formParams.ossType === 'huawei'">
            <n-form-item path="huaweiPath" label="路径">
              <n-input v-model:value="formParams.huaWeiObsConfig.path" />
            </n-form-item>
            <n-form-item path="huaweiBucket" label="储存桶">
              <n-input v-model:value="formParams.huaWeiObsConfig.bucket" />
            </n-form-item>
            <n-form-item path="huaweiEndpoint" label="域名">
              <n-input v-model:value="formParams.huaWeiObsConfig.endpoint" />
            </n-form-item>
            <n-form-item path="huaweiAccessKey" label="密钥key">
              <n-input v-model:value="formParams.huaWeiObsConfig.accessKey" />
            </n-form-item>
            <n-form-item path="huaweiSecretKey" label="密钥secret">
              <n-input v-model:value="formParams.huaWeiObsConfig.secretKey" />
            </n-form-item>
          </div>
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

  const options = ref([
    {
      label: "本地",
      value: "local"
    },
    {
      label: "阿里云",
      value: "aliyun"
    },
    {
      label: "腾讯云",
      value: "tencent"
    },
    {
      label: "七牛云",
      value: "qiniu"
    },
    {
      label: "华为云",
      value: "huawei"
    }
  ])

  const rules = {
    title: {
      required: true,
      trigger: 'blur',
      message: '请输入2到20位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,20}$/, //验证菜单名称 2到20位（字母,汉字,数字,下划线）
    },
  };

  let formParams = ref<any>({
    ossType: 'local',
    localConfig: {
      path: 'uploads/file'
    },
    aliyunOssConfig: {
      endpoint: '',
      accessKeyId: '',
      accessKeySecret: '',
      bucketName: 'pinkmoe',
      bucketPoint: '',
    },
    tencentCOSConfig: {
      bucket: '',
      region: '',
      secretId: '',
      secretKey: '',
      baseUrl: '',
      pathPrefix: '',
    },
    qiniuConfig: {
      zone: '',
      bucket: '',
      imgPath: '',
      useHttps: false,
      accessKey: '',
      secretKey: '',
      useCdnDomains: false,
    },
    huaWeiObsConfig: {
      Path: '',
      bucket: '',
      endpoint: '',
      accessKey: '',
      secretKey: '',
    }
  });

  const message = useMessage();

  async function handleSubmit() {
    const v = JSON.stringify(formParams.value);
    const { code, message: msg } = await updateSetting({
      slug: 'system_oss',
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
      slug: 'system_oss',
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
