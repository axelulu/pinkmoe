<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-07 23:37:51
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-21 17:00:01
 * @FilePath: /pinkmoe_admin/src/views/shop/category/CreateEditForm.vue
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
      <n-form-item label="名称" path="name">
        <n-input placeholder="请输入名称" v-model:value="formParams.name" />
      </n-form-item>
      <n-form-item label="标识" path="slug">
        <n-input placeholder="请输入标识" v-model:value="formParams.slug" />
      </n-form-item>
      <n-form-item label="图标" path="icon">
        <n-input placeholder="请输入图标" v-model:value="formParams.icon" />
      </n-form-item>
      <n-form-item label="父分类" path="parentId">
        <n-cascader
          :virtual-scroll="true"
          cascade="false"
          label-field="name"
          value-field="ID"
          v-model:value="formParams.parentId"
          :options="options"
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
  import { defineComponent, reactive, ref, toRefs } from 'vue';
  import { useMessage } from 'naive-ui';
  import { ResultEnum } from '@/enums/httpEnum';
  import { createShopCategory, updateShopCategory, getShopCategoryList } from '@/api/shop/category';

  const rules = {
    name: {
      required: true,
      trigger: 'blur',
      message: '请输入2到16位（字母,汉字,数字,下划线）',
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]{2,30}$/, //验证菜单名称 2到16位（字母,汉字,数字,下划线）
    },
    slug: {
      required: true,
      trigger: 'blur',
      message: '请输入2到100位（字母）',
      pattern: /^[0-9a-zA-Z/_-]{2,100}$/, //验证菜单路径 2到100位（字母）
    },
    icon: {
      required: true,
      trigger: 'blur',
      message: '请输入2到80位（字母）',
      pattern: /^[a-zA-Z0-9:_-]{2,80}$/, //验证图标 2到30位（字母）
    },
    parentId: {
      required: true,
      trigger: 'blur',
      message: '请输入1-10位数字',
      pattern: /^[0-9]{1,10}$/, //验证父菜单 1-2位数字
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
      const options = ref([]);
      const defaultValueRef = () => ({
        ID: 0,
        name: '',
        slug: '',
        icon: '',
        parentId: 0,
        sort: 0,
      });

      function disabled(menuTreeData, treeItemId) {
        if (menuTreeData && menuTreeData.length > 0) {
          menuTreeData.forEach((item) => {
            item.disabled = treeItemId == item.ID;
            if (item.children !== null) {
              disabled(item.children, treeItemId);
            }
          });
        }
      }

      async function getCategoryLists() {
        const {
          code,
          message: msg,
          result,
        } = await getShopCategoryList({
          pageSize: 999,
          page: 1,
        });
        message.destroyAll();
        if (code == ResultEnum.SUCCESS) {
          if (result.list) {
            options.value = result.list;
          }
          disabled(options.value, state.formParams.ID);
          options.value.unshift({
            name: '根目录',
            ID: 0,
          });
        } else {
          message.error(msg || '获取失败');
        }
      }

      const state = reactive({
        showForm: false,
        subLoading: false,
        formParams: defaultValueRef(),
      });

      function confirmForm(e) {
        e.preventDefault();
        formBtnLoading.value = true;
        formRef.value.validate(async (errors) => {
          if (!errors) {
            const { code, message: msg } =
              props.formType === 'create'
                ? await createShopCategory(state.formParams)
                : await updateShopCategory(state.formParams);
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
        getCategoryLists();
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
