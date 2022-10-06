<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-01 16:04:09
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:25:46
 * @FilePath: /pinkmoe_admin/src/views/system/api/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <n-modal v-model:show="showForm" :show-icon="false" preset="dialog" title="新建">
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="80"
      class="py-4"
    >
      <n-form-item label="路径" path="path">
        <n-input placeholder="请输入路径" v-model:value="formParams.path" />
      </n-form-item>
      <n-form-item label="请求" path="method">
        <n-select v-model:value="formParams.method" :options="options" />
      </n-form-item>
      <n-form-item label="分组" path="apiGroup">
        <n-input placeholder="请输入分组" v-model:value="formParams.apiGroup" />
      </n-form-item>
      <n-form-item label="简介" path="description">
        <n-input type="textarea" placeholder="请输入简介" v-model:value="formParams.description" />
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
  import { createApi, updateApi } from '@/api/system/api';

  const rules = {
    path: {
      required: true,
      trigger: 'blur',
      message: '请输入2到100位（字母）',
      pattern: /^[0-9a-zA-Z/_-]{2,100}$/, //验证菜单路径 2到100位（字母）
    },
    method: {
      type: 'enum',
      enum: ['GET', 'POST', 'PUT', 'DELETE'],
      required: true,
      trigger: 'blur',
      message: '请输入地址',
    },
    apiGroup: {
      required: true,
      trigger: 'blur',
      message: '请输入2到30位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,30}$/, //验证菜单名称 2到30位（字母,汉字,数字,下划线）
    },
    description: {
      required: true,
      trigger: 'blur',
      message: '请输入2到30位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,30}$/, //验证菜单名称 2到30位（字母,汉字,数字,下划线）
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
      const options = [
        {
          label: 'GET',
          value: 'GET',
        },
        {
          label: 'POST',
          value: 'POST',
        },
        {
          label: 'PUT',
          value: 'PUT',
        },
        {
          label: 'DELETE',
          value: 'DELETE',
        },
      ];
      const defaultValueRef = () => ({
        ID: 0,
        path: '',
        method: 'GET',
        apiGroup: '',
        description: '',
      });

      const state = reactive({
        showForm: false,
        formParams: defaultValueRef(),
      });

      function confirmForm(e) {
        e.preventDefault();
        formBtnLoading.value = true;
        formRef.value.validate(async (errors) => {
          if (!errors) {
            const { code, message: msg } =
              props.formType === 'create'
                ? await createApi(state.formParams)
                : await updateApi(state.formParams);
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
        options,
        formBtnLoading,
        openForm,
        closeForm,
        confirmForm,
      };
    },
  });
</script>
