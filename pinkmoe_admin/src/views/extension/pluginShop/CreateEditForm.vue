<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 22:24:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:00
 * @FilePath: /pinkmoe_admin/src/views/extension/pluginShop/CreateEditForm.vue
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
      <n-form-item label="名称" path="label">
        <n-input placeholder="请输入名称" v-model:value="formParams.label" />
      </n-form-item>
      <n-form-item label="标识" path="value">
        <n-input placeholder="请输入标识" v-model:value="formParams.value" />
      </n-form-item>
      <n-form-item label="图标" path="icon">
        <BasicUpload
          :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
          :headers="uploadHeaders"
          max-size="1"
          max-number="1"
          v-model:value="icon"
          @uploadChange="uploadTopicChange"
        />
      </n-form-item>
      <n-form-item label="排序" path="sort">
        <n-input-number
          placeholder="请输入排序"
          max="100"
          min="0"
          v-model:value="formParams.sort"
        />
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
  import { defineComponent, reactive, ref, toRefs, watch } from 'vue';
  import { useMessage } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { createTopic, updateTopic } from '@/api/content/topic';
  import { useGlobSetting } from '@/hooks/setting';
  import { createStorage } from '@/utils/Storage';
  import { ACCESS_TOKEN } from '@/store/mutation-types';
  import { BasicUpload } from '@/components/Upload';

  const globSetting = useGlobSetting();
  const { uploadUrl } = globSetting;
  const Storage = createStorage({ storage: localStorage });

  const rules = {
    label: {
      required: true,
      trigger: 'blur',
      message: '请输入2到16位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,30}$/, //验证菜单名称 2到16位（字母,汉字,数字,下划线）
    },
    value: {
      required: true,
      trigger: 'blur',
      message: '请输入2到100位（字母）',
      pattern: /^[0-9a-zA-Z/_-]{2,100}$/, //验证菜单路径 2到100位（字母）
    },
    icon: {
      required: true,
      trigger: 'blur',
      message: '请选择正确的图标',
      pattern: /\.(png|jpg|jpeg|gif|svg)(\?.*)?/,
    },
    sort: {
      required: true,
      trigger: 'blur',
      message: '请输入1-2位数字',
      pattern: /^[0-9]{1,2}$/, //验证排序内容 1-2位数字
    },
  };

  export default defineComponent({
    name: 'CreateEditForm',
    components: {
      BasicUpload,
    },
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
      const options = [];
      const defaultValueRef = () => ({
        ID: 0,
        label: '',
        value: '',
        icon: '',
        sort: '0',
      });

      const uploadHeaders = reactive({
        platform: 'miniPrograms',
        timestamp: new Date().getTime(),
        Authorization: 'Bearer ' + Storage.get(ACCESS_TOKEN, ''),
      });

      const state = reactive({
        showForm: false,
        icon: [] as string[],
        formParams: defaultValueRef(),
      });

      function uploadTopicChange(list: string[]) {
        state.icon = list;
      }

      function confirmForm(e) {
        e.preventDefault();
        formBtnLoading.value = true;
        formRef.value.validate(async (errors) => {
          if (!errors) {
            state.formParams.icon = state.icon[0];
            const { code, message: msg } =
              props.formType === 'create'
                ? await createTopic(state.formParams)
                : await updateTopic(state.formParams);
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

      watch(
        () => props.formType,
        () => {
          if (props.formType == 'create') {
            state.icon = [];
          }
        }
      );

      function openForm(record) {
        if (record !== null) {
          state.formParams = record;
        }
        if (state.formParams.icon) {
          state.icon[0] = state.formParams.icon;
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
        uploadHeaders,
        uploadUrl,
        uploadTopicChange,
        openForm,
        closeForm,
        confirmForm,
      };
    },
  });
</script>
