<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-05 10:47:13
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:22:53
 * @FilePath: /pinkmoe_admin/src/views/content/comment/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
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
      <n-form-item label="类型" path="type">
        <n-select placeholder="请输入评论类型" v-model:value="formParams.type" :options="options" />
      </n-form-item>
      <n-form-item v-if="formParams.type === 'post'" label="评论文章" path="postId">
        <n-select
          v-model:value="formParams.postId"
          :render-label="renderPostLabel"
          :render-tag="renderPostSelectTag"
          :options="postList"
          clearable
          remote
        />
      </n-form-item>
      <n-form-item v-if="formParams.type === 'comment'" label="父评论" path="parentId">
        <n-cascader
          :virtual-scroll="true"
          cascade="false"
          label-field="content"
          value-field="ID"
          v-model:value="formParams.parentId"
          :render-label="renderCommentLabel"
          :options="commentOptions"
        />
      </n-form-item>
      <n-form-item label="来自用户" path="formUid">
        <n-select
          v-model:value="formParams.formUid"
          :render-label="renderUserLabel"
          :render-tag="renderUserSelectTag"
          :options="userList"
          clearable
          remote
        />
      </n-form-item>
      <n-form-item v-if="formParams.type === 'comment'" label="评论用户" path="toUid">
        <n-select
          v-model:value="formParams.toUid"
          :render-label="renderUserLabel"
          :render-tag="renderUserSelectTag"
          :options="userList"
          clearable
          remote
        />
      </n-form-item>
      <n-form-item label="内容" path="content">
        <n-input type="textarea" placeholder="请输入评论内容" v-model:value="formParams.content" />
      </n-form-item>
      <n-form-item label="状态" path="status">
        <n-radio-group v-model:value="formParams.status" name="status">
          <n-radio value="pending"> 待审核</n-radio>
          <n-radio value="published"> 已发布</n-radio>
        </n-radio-group>
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
import { defineComponent, h, reactive, ref, toRefs, watch } from "vue";
import {
  NAvatar,
  NText,
  SelectOption,
  SelectRenderLabel,
  SelectRenderTag,
  useMessage
} from "naive-ui";
import { ResultEnum } from "@/enums/httpEnum";
import { createComment, getCommentList, updateComment } from "@/api/content/comment";
import { getUserList } from "@/api/system/user";
import { getPostList } from "@/api/content/post";

const rules = {
  type: {
    required: true,
    type: "enum",
    enum: ["post", "comment"],
    trigger: "blur",
    message: "请输入2到100位（字母）"
  },
  postId: {
    required: true,
    trigger: "blur",
    message: "请输入文章ID",
    pattern: /^[0-9]{2,30}$/ //验证文章ID
  },
  parentId: {
    required: true,
    trigger: "blur",
    message: "请输入评论ID",
    pattern: /^[0-9]{2,30}$/ //验证评论ID
  },
  formUid: {
    required: true,
    trigger: "blur",
    message: "请输入用户ID",
    pattern: /^[0-9a-zA-Z_-]{2,40}$/ //验证用户ID
  },
  toUid: {
    required: true,
    trigger: "blur",
    message: "请输入用户ID",
    pattern: /^[0-9_-]{2,30}$/ //验证用户ID
  },
  content: {
    required: true,
    trigger: "blur",
    message: "请输入文章内容",
  },
  status: {
    required: true,
    trigger: "blur",
    type: "enum",
    enum: ["published", "pending"],
    message: "请输入评论状态",
    pattern: /^[0-9_-]{2,30}$/ //验证评论状态
  }
};

export default defineComponent({
  name: "CreateEditForm",
  components: {},
  props: {
    formType: {
      type: String,
      default: "create"
    }
  },
  setup(props, context) {
    const message = useMessage();
    const formBtnLoading = ref(false);
    const formRef: any = ref(null);
    const userList = ref<SelectOption[]>([]);
    const postList = ref<SelectOption[]>([]);
    const commentOptions = ref<SelectOption[]>([]);
    const options = [
      {
        label: "文章",
        value: "post"
      },
      {
        label: "评论",
        value: "comment"
      }
    ];
    const defaultValueRef = () => ({
      ID: 0,
      postId: null,
      parentId: 0,
      formUid: null,
      toUid: null,
      content: "",
      type: "post",
      status: "published",
      children: null
    });

    const state = reactive({
      showForm: false,
      formParams: defaultValueRef()
    });

    const renderUserSelectTag: SelectRenderTag = ({ option }) => {
      return h(
        "div",
        {
          style: {
            display: "flex",
            alignItems: "center"
          }
        },
        [
          h(NAvatar, {
            src: option.avatar as string,
            round: true,
            size: 24,
            objectFit: "cover",
            style: {
              marginRight: "12px"
            }
          }),
          option.label as string
        ]
      );
    };

    const renderUserLabel: SelectRenderLabel = (option) => {
      return h(
        "div",
        {
          style: {
            display: "flex",
            alignItems: "center"
          }
        },
        [
          h(NAvatar, {
            src: option.avatar as string,
            round: true,
            objectFit: "cover",
            size: "small"
          }),
          h(
            "div",
            {
              style: {
                marginLeft: "12px",
                padding: "4px 0"
              }
            },
            [
              h("div", null, [option.label as string]),
              h(
                NText,
                { depth: 3, tag: "div" },
                {
                  default: () => option.nickName as string
                }
              )
            ]
          )
        ]
      );
    };

    const renderPostSelectTag: SelectRenderTag = ({ option }) => {
      return h(
        "div",
        {
          style: {
            display: "flex",
            alignItems: "center"
          }
        },
        [
          h(NAvatar, {
            src: option.cover as string,
            round: true,
            objectFit: "cover",
            size: 24,
            style: {
              marginRight: "12px"
            }
          }),
          option.label as string
        ]
      );
    };

    const renderPostLabel: SelectRenderLabel = (option) => {
      return h(
        "div",
        {
          style: {
            display: "flex",
            alignItems: "center"
          }
        },
        [
          h(NAvatar, {
            src: option.cover as string,
            round: true,
            objectFit: "cover",
            size: "small"
          }),
          h(
            "div",
            {
              style: {
                marginLeft: "12px",
                padding: "4px 0"
              }
            },
            [
              h("div", null, [option.title as string]),
              h(
                NText,
                { depth: 3, tag: "div" },
                {
                  default: () => option.exerpt as string
                }
              )
            ]
          )
        ]
      );
    };

    async function getUserLists() {
      const {
        code,
        message: msg,
        result
      } = await getUserList({
        pageSize: 999,
        page: 1
      });
      if (code != ResultEnum.SUCCESS) {
        message.error(msg || "获取失败");
      }
      result.list.forEach((item) => {
        item.label = item.userName;
        item.value = item.uuid;
      });
      userList.value = result.list;
    }

    async function getPostLists() {
      const {
        code,
        message: msg,
        result
      } = await getPostList({
        pageSize: 999,
        page: 1
      });
      if (code != ResultEnum.SUCCESS) {
        message.error(msg || "获取失败");
      }
      result.list.forEach((item) => {
        item.label = item.title;
        item.value = item.postId;
      });
      postList.value = result.list;
    }

    async function getCommentLists() {
      const {
        code,
        message: msg,
        result
      } = await getCommentList({
        pageSize: 999,
        page: 1
      });
      if (code != ResultEnum.SUCCESS) {
        message.error(msg || "获取失败");
      }
      commentOptions.value = result.list;
    }

    function confirmForm(e) {
      e.preventDefault();
      formBtnLoading.value = true;
      formRef.value.validate(async (errors) => {
        if (!errors) {
          const { code, message: msg } =
            props.formType === "create"
              ? await createComment(state.formParams)
              : await updateComment(state.formParams);
          message.destroyAll();
          if (code == ResultEnum.SUCCESS) {
            message.success(props.formType === "create" ? "新建成功" : "更新成功");
            setTimeout(() => {
              state.showForm = false;
              closeForm();
              context.emit("reloadTable");
            });
          } else {
            message.error(msg || (props.formType === "create" ? "新建失败" : "更新失败"));
          }
        } else {
          message.error("请填写完整信息");
        }
        formBtnLoading.value = false;
      });
    }

    watch(
      () => props.formType,
      () => {
        if (props.formType == "create") {
          Object.assign(state.formParams, defaultValueRef());
        }
      }
    );

    function openForm(record) {
      if (record !== null) {
        state.formParams = record;
      }
      getUserLists();
      getPostLists();
      getCommentLists();
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
      userList,
      postList,
      commentOptions,
      renderPostLabel,
      renderPostSelectTag,
      renderUserLabel,
      renderUserSelectTag,
      openForm,
      closeForm,
      confirmForm,
      renderCommentLabel(option: { value?: string | number; content?: string }) {
        return `评论: ${option.content}`;
      }
    };
  }
});
</script>
