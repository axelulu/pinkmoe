<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:15
 * @FilePath: /pinkmoe_admin/src/views/form/stepForm/Step2.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <n-form
    :label-width="90"
    :model="formValue"
    :rules="rules"
    label-placement="left"
    ref="form2Ref"
    style="max-width: 500px; margin: 40px auto 0 80px"
  >
    <n-form-item label="付款账户" path="myAccount">
      <span>NaiveUiAdmin@163.com</span>
    </n-form-item>
    <n-form-item label="收款账户" path="account">
      <span>NaiveUiAdmin@qq.com</span>
    </n-form-item>
    <n-form-item label="收款人姓名" path="name">
      <span>Ah jung</span>
    </n-form-item>
    <n-form-item label="转账金额" path="money">
      <span>￥1980</span>
    </n-form-item>
    <n-divider />
    <n-form-item label="支付密码" path="password">
      <n-input type="password" v-model:value="formValue.password" />
    </n-form-item>
    <div style="margin-left: 80px">
      <n-space>
        <n-button type="primary" :loading="loading" @click="formSubmit">提交</n-button>
        <n-button @click="prevStep">上一步</n-button>
      </n-space>
    </div>
  </n-form>
</template>

<script lang="ts" setup>
  import { ref, defineEmits } from 'vue';
  import { useMessage } from 'naive-ui';
  const form2Ref: any = ref(null);
  const message = useMessage();
  const loading = ref(false);

  const formValue = ref({
    password: '086611',
  });

  const rules = {
    password: {
      required: true,
      message: '请输入支付密码',
      trigger: 'blur',
    },
  };

  const emit = defineEmits(['prevStep', 'nextStep']);

  function prevStep() {
    emit('prevStep');
  }

  function formSubmit() {
    loading.value = true;
    form2Ref.value.validate((errors) => {
      if (!errors) {
        setTimeout(() => {
          emit('nextStep');
        }, 1500);
      } else {
        message.error('验证失败，请填写完整信息');
      }
    });
  }
</script>
