<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-30 12:02:13
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 09:21:33
 * @FilePath: /pinkmoe_admin/src/views/system/role/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-modal
    v-model:show="showForm"
    :show-icon="false"
    preset="dialog"
    style="width: 800px"
    :title="formType === 'create' ? '新建' : '更新'"
  >
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="80"
      class="py-4"
    >
      <n-form-item label="角色ID" path="authorityId">
        <n-input
          :disabled="formType !== 'create'"
          placeholder="请输入角色ID"
          v-model:value="formParams.authorityId"
        />
      </n-form-item>
      <n-form-item label="角色名称" path="authorityName">
        <n-input placeholder="请输入角色名称" v-model:value="formParams.authorityName" />
      </n-form-item>
      <n-form-item label="角色权重" path="authorityName">
        <n-input-number placeholder="请输入角色权重" v-model:value="formParams.authorityWeight" />
      </n-form-item>
      <n-form-item label="允许开通会员" path="authorityName">
        <n-switch placeholder="请选择是否允许开通会员" v-model:value="formParams.vipStart" />
      </n-form-item>
      <n-form-item label="角色颜色" path="authorityColor">
        <n-color-picker placeholder="请输入角色颜色" v-model:value="formParams.authorityColor" />
      </n-form-item>
      <n-form-item label="角色扩展参数" path="authorityParams">
        <n-dynamic-input v-model:value="formParams.authorityParams" :on-create="onCreate">
          <template #create-button-default>
            添加等级配置
          </template>
          <template #default="{ value }">
            <div style="display: flex; align-items: center; width: 100%">
              <n-input v-model:value="value.authorityParamId"
                       style="width: 140px;margin-right: 12px" type="text" placeholder="扩展参数ID" />
              <n-input v-model:value="value.authorityParamKey"
                       style="width: 130px;margin-right: 12px" type="text" placeholder="扩展参数名称" />
              <n-input-number max="100000000" min="0" v-model:value="value.authorityParamValue"
                       style="width: 140px;margin-right: 12px" type="text" placeholder="扩展参数值" />
              <n-input-number max="100000000" min="0" v-model:value="value.authorityParamDay"
                              style="width: 140px;margin-right: 12px" type="text" placeholder="扩展参数值2" />
            </div>
          </template>
        </n-dynamic-input>
      </n-form-item>
    </n-form>

    <template #action>
      <n-space>
        <n-button @click="() => (showForm = false)">取消</n-button>
        <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
      </n-space>
    </template>
  </n-modal>
</template>

<script lang="ts">
  import { defineComponent, reactive, ref, toRefs } from 'vue';
  import { useMessage } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { createAuthority, updateAuthority } from '@/api/system/authority';

  const rules = {
    authorityId: {
      required: true,
      message: '请输入4位数字',
      trigger: 'blur',
      pattern: /^[0-9]{4}$/, //验证角色ID 4位手机号
    },
    authorityName: {
      required: true,
      trigger: 'blur',
      message: '2到16位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,16}$/, //验证角色名称 2到16位（字母,汉字,数字,下划线）
    },
  };

  export default defineComponent({
    name: 'CreateEditForm',
    components: {},
    props: {
      formType: {
        type: String,
        default: 'create',
      },
    },
    setup(props, context) {
      const message = useMessage();
      const formBtnLoading = ref(false);
      const formRef: any = ref(null);
      const defaultValueRef = () => ({
        authorityName: '',
        authorityWeight: 10,
        vipStart: 1,
        authorityId: '',
        authorityColor: '',
        authorityParams: [
          {
            authorityParamKey: '',
            authorityParamValue: '',
            authorityParamDay: '',
            authorityParamId: '',
          }
        ],
      });

      function onCreate(){
        return {
          authorityParamKey: '',
          authorityParamValue: '',
          authorityParamDay: '',
          authorityParamId: '',
        }
      }

      const state = reactive({
        showForm: false,
        formParams: defaultValueRef(),
      });

      function confirmForm(e) {
        e.preventDefault();
        formBtnLoading.value = true;
        formRef.value.validate(async (errors) => {
          if (!errors) {
            state.formParams.vipStart = state.formParams.vipStart ? 1 : 0
            const { code, message: msg } =
              props.formType === 'create'
                ? await createAuthority(state.formParams)
                : await updateAuthority(state.formParams);
            message.destroyAll();
            if (code == ResultEnum.SUCCESS) {
              message.success(props.formType === 'create' ? '新建成功' : '更新成功');
              setTimeout(() => {
                state.showForm = false;
                closeForm();
                context.emit('reloadTable');
              });
            } else {
              message.error(msg || (props.formType === 'create' ? '新建失败' : '更新失败'));
            }
          } else {
            message.error('请填写完整信息');
          }
          formBtnLoading.value = false;
        });
      }

      function openForm(record) {
        if (record !== null) {
          record.vipStart = record.vipStart === 1 ? true : false
          state.formParams = record;
        }
        state.showForm = true;
      }

      function closeForm() {
        state.showForm = false;
      }

      return {
        ...toRefs(state),
        formRef,
        rules,
        onCreate,
        formBtnLoading,
        openForm,
        closeForm,
        confirmForm,
      };
    },
  });
</script>
