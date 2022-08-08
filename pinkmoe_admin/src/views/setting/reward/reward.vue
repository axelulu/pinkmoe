<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 22:25:04
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-08 16:12:25
 * @FilePath: /pinkmoe_admin/src/views/setting/reward/reward.vue
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
          <n-form-item path="publishPost" label="投稿奖励">
            <n-select class="w-20 mr-2" v-model:value="formParams.publishPostType" :options="options" />
            <n-input-number v-model:value="formParams.publishPostCredit" />
          </n-form-item>
          <n-form-item path="reg" label="注册奖励">
            <n-select class="w-20 mr-2" v-model:value="formParams.regType" :options="options" />
            <n-input-number v-model:value="formParams.regCredit" />
          </n-form-item>
          <n-form-item path="forgetPwd" label="忘记密码扣除">
            <n-select class="w-20 mr-2" v-model:value="formParams.forgetPwdType" :options="options" />
            <n-input-number type="textarea" v-model:value="formParams.forgetPwdCredit" />
          </n-form-item>
          <n-form-item path="checkIn" label="签到奖励区间">
            <n-select class="w-20 mr-2" v-model:value="formParams.checkInType" :options="options" />
            <n-input-number class="mr-2" v-model:value="formParams.checkInCreditHead" />
            <n-input-number v-model:value="formParams.checkInCreditFoot" />
          </n-form-item>
          <n-form-item path="updateAvatar" label="修改头像扣除">
            <n-select class="w-20 mr-2" v-model:value="formParams.updateAvatarType" :options="options" />
            <n-input-number v-model:value="formParams.updateAvatarCredit" />
          </n-form-item>
          <n-form-item path="updateEmail" label="修改邮箱扣除">
            <n-select class="w-20 mr-2" v-model:value="formParams.updateEmailType" :options="options" />
            <n-input-number v-model:value="formParams.updateEmailCredit" />
          </n-form-item>
          <n-form-item path="updatePwd" label="修改密码扣除">
            <n-select class="w-20 mr-2" v-model:value="formParams.updatePwdType" :options="options" />
            <n-input-number v-model:value="formParams.updatePwdCredit" />
          </n-form-item>
          <n-form-item path="comment" label="评论奖励">
            <n-select class="w-20 mr-2" v-model:value="formParams.commentType" :options="options" />
            <n-input-number v-model:value="formParams.commentCredit" />
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

  const options = ref([
    {
      label: "积分",
      value: "credit"
    },
    {
      label: "现金",
      value: "cash"
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
    publishPostType: 'credit',
    publishPostCredit: 0,
    commentType: 'credit',
    commentCredit: 0,
    regType: 'credit',
    regCredit: 0,
    updatePwdType: 'credit',
    updatePwdCredit: 0,
    updateEmailType: 'credit',
    updateEmailCredit: 0,
    updateAvatarType: 'credit',
    updateAvatarCredit: 0,
    checkInType: 'credit',
    checkInCreditHead: 0,
    checkInCreditFoot: 0,
    forgetPwdType: 'credit',
    forgetPwdCredit: 0,
  });

  const message = useMessage();

  async function handleSubmit() {
    const v = JSON.stringify(formParams.value);
    const { code, message: msg } = await updateSetting({
      slug: 'user_reward',
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
      slug: 'user_reward',
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
