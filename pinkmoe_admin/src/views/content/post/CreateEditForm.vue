<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 15:46:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:23:00
 * @FilePath: /pinkmoe_admin/src/views/content/post/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <n-modal
    v-model:show="showForm"
    style="width: 1000px"
    :show-icon="false"
    preset="dialog"
    title="新建"
  >
    <n-form
      :model="formParams"
      :rules="rules"
      ref="formRef"
      label-placement="left"
      :label-width="80"
      class="py-4"
    >
      <n-form-item label="标题" path="title">
        <n-input placeholder="请输入标题" v-model:value="formParams.title" />
      </n-form-item>
      <n-form-item label="摘要" path="exerpt">
        <n-input type="textarea" placeholder="请输入摘要" v-model:value="formParams.exerpt" />
      </n-form-item>
      <n-form-item label="内容" path="content">
        <wangEditor ref="wangEditorRef" v-model:value="formParams.content" />
      </n-form-item>
      <n-form-item label="图片" path="cover">
        <PostImgUpload
          ref="postImgUpload"
          :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
          :headers="uploadHeaders"
          :data="{'postId': formParams.postId, 'uuid': userMeta.uuid, 'type': 'post'}"
          max-size="1"
          max-number="12"
          height="170"
          width="150"
          v-model:value="formParams.postImg"
          v-model:cover="formParams.cover"
          @setCover="setCover"
          @uploadChange="uploadChange"
          @insertPost="insertPostImg"
          @delete="deletePostImg"
        />
      </n-form-item>
      <n-form-item label="作者" path="author">
        <n-select
          v-model:value="formParams.author"
          :render-label="renderAuthorLabel"
          :render-tag="renderAuthorSelectTag"
          :options="userList"
          clearable
          remote
        />
      </n-form-item>
      <n-form-item label="话题" path="topic">
        <n-select
          v-model:value="topic"
          :render-label="renderTopicLabel"
          :render-tag="renderTopicSelectTag"
          :options="topicOptions"
          filterable
          multiple
          tag
          placeholder="输入，按回车确认"
        />
      </n-form-item>
      <n-form-item label="查看数" path="view">
        <n-input-number
          min="0"
          max="100000000"
          placeholder="请输入查看数"
          v-model:value="formParams.view"
        >
        </n-input-number>
      </n-form-item>
      <n-form-item label="下载" path="download">
        <n-dynamic-input v-model:value="formParams.downloadRelation" show-sort-button
                         :on-create="onCreateMusic">
          <template #create-button-default> 请填写下载信息</template>
          <template #default="{ value }">
            <div style="display: flex; align-items: center; width: 100%">
              <div class="w-2/12 pr-1">
                <n-input
                  maxlength="30"
                  show-count
                  placeholder="请输入下载名称"
                  v-model:value="value.name"
                  type="text"
                />
              </div>
              <div class="w-3/12 pr-1">
                <n-input
                  placeholder="请输入下载链接"
                  :input-props="{ type: 'link' }"
                  v-model:value="value.url"
                  type="text"
                >
                  <template #suffix>
                    <n-icon :component="GlobeOutline" />
                  </template>
                </n-input>
              </div>
              <div class="w-2/12 pr-1">
                <n-input
                  maxlength="30"
                  show-count
                  placeholder="请输入提取密码"
                  v-model:value="value.extract_pwd"
                  type="text"
                />
              </div>
              <div class="w-2/12 pr-1">
                <n-input
                  maxlength="30"
                  show-count
                  placeholder="请输入解压密码"
                  v-model:value="value.unpack_pwd"
                  type="text"
                />
              </div>
              <div class="w-1/12 pr-1">
                <n-select v-model:value="value.price_type" :options="[
                {
                  'label': '现金',
                  'value': 'cash',
                },{
                  'label': '积分',
                  'value': 'credit',
                }
              ]" />
              </div>
              <div class="w-2/12">
                <n-input-number
                  max="10000"
                  clearable
                  placeholder="请输入价格"
                  v-model:value="value.price"
                />
              </div>
            </div>
          </template>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item label="类型" path="type">
        <n-radio-group v-model:value="formParams.type" name="type">
          <n-radio value="post"> 文章</n-radio>
          <n-radio value="music"> 音乐</n-radio>
          <n-radio value="video"> 视频</n-radio>
        </n-radio-group>
      </n-form-item>
      <n-form-item v-if="formParams.type === 'music'" label="音乐" path="music">
        <n-dynamic-input v-model:value="formParams.musicRelation" show-sort-button
                         :on-create="onCreateMusic">
          <template #create-button-default> 请填写音乐信息</template>
          <template #default="{ value }">
            <div style="display: flex; align-items: center; width: 100%">
              <n-input
                maxlength="30"
                show-count
                placeholder="请输入音乐名称"
                style="margin-right: 12px"
                v-model:value="value.name"
                type="text"
              />
              <n-input
                placeholder="请输入音乐链接"
                :input-props="{ type: 'url' }"
                v-model:value="value.url"
                type="text"
              >
                <template #suffix>
                  <n-icon :component="GlobeOutline" />
                </template>
              </n-input>
            </div>
          </template>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item v-if="formParams.type === 'video'" label="视频" path="video">
        <n-dynamic-input v-model:value="formParams.videoRelation" show-sort-button
                         :on-create="onCreateVideo">
          <template #create-button-default> 请填写视频信息</template>
          <template #default="{ value }">
            <div style="display: flex; align-items: center; width: 100%">
              <n-input
                maxlength="30"
                show-count
                placeholder="请输入视频名称"
                style="margin-right: 12px"
                v-model:value="value.name"
                type="text"
              />
              <n-input
                placeholder="请输入视频链接"
                style="margin-right: 12px"
                :input-props="{ type: 'url' }"
                v-model:value="value.url"
                type="text"
              >
                <template #suffix>
                  <n-icon :component="GlobeOutline" />
                </template>
              </n-input>
              <n-input
                placeholder="请输入视频字幕链接"
                :input-props="{ type: 'url' }"
                v-model:value="value.subtitles"
                type="text"
              >
                <template #suffix>
                  <n-icon :component="GlobeOutline" />
                </template>
              </n-input>
            </div>
          </template>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item label="分类" path="category">
        <n-cascader
          :virtual-scroll="true"
          cascade="false"
          label-field="name"
          value-field="slug"
          v-model:value="formParams.category"
          :options="options"
        />
      </n-form-item>
      <n-form-item label="来源" path="type">
        <n-radio-group v-model:value="formParams.from" name="type">
          <n-radio value="original"> 原创</n-radio>
          <n-radio value="reprinted"> 转载</n-radio>
        </n-radio-group>
      </n-form-item>
      <n-form-item label="状态" path="status">
        <n-radio-group v-model:value="formParams.status" name="status">
          <n-radio value="published"> 已发布</n-radio>
          <n-radio value="pending"> 待审核</n-radio>
          <n-radio value="draft"> 草稿</n-radio>
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
import { computed, defineComponent, h, reactive, ref, toRefs, watch } from "vue";
import {
  useMessage,
  SelectOption,
  NText,
  SelectRenderLabel,
  NAvatar,
  SelectRenderTag
} from "naive-ui";
import { GlobeOutline } from "@vicons/ionicons5";
import { ResultEnum } from "@/enums/httpEnum";
import { createPost, updatePost } from "@/api/content/post";
import { getCategoryList } from "@/api/content/category";
import { useGlobSetting } from "@/hooks/setting";
import { ACCESS_TOKEN, CURRENT_USER } from "@/store/mutation-types";
import { PostImgUpload } from "@/components/UploadPostImg";
import { createStorage } from "@/utils/Storage";
import { getUserList } from "@/api/system/user";
import wangEditor from "./wangEditor.vue";
import "@vueup/vue-quill/dist/vue-quill.snow.css";
import { getTopicList } from "@/api/content/topic";
import { deleteImgFile } from "@/api/upload/upload";

const globSetting = useGlobSetting();
const { uploadUrl } = globSetting;
const Storage = createStorage({ storage: localStorage });

const rules = {
  title: {
    required: true,
    trigger: "blur",
    message: "请输入2到50位标题内容"
    // pattern: /^[\u4E00-\u9FA5A-Za-z0-9_]{2,50}$/ //验证菜单路径 2到50位（字母）
  },
  exerpt: {
    required: true,
    trigger: "blur",
    message: "请输入2到50位摘要内容"
    // pattern: /^[\u4E00-\u9FA5A-Za-z0-9_]{2,100}$/ //验证菜单路径 2到100位（字母）
  },
  content: {
    required: true,
    trigger: "blur",
    message: "请输入文章内容"
  },
  cover: {
    required: false,
    trigger: "blur",
    message: "请输入文章封面"
  },
  view: {
    required: true,
    trigger: "blur",
    type: "number",
    message: "请填写用户现金",
    pattern: /^[0-9]{11}$/ //验证用户现金
  },
  author: {
    required: true,
    trigger: "blur",
    message: "请输入文章作者",
    pattern: /^[a-zA-Z0-9_-]{2,40}$/ //验证文章作者
  },
  // key: {
  //   required: false,
  //   trigger: "blur",
  //   message: "请输入文章话题",
  // },
  type: {
    required: true,
    trigger: "blur",
    type: "enum",
    enum: ["post", "music", "video"],
    message: "请输入文章类型"
  },
  // music: {
  //   required: true,
  //   trigger: "blur",
  //   message: "请输入音乐"
  // },
  // video: {
  //   required: true,
  //   trigger: "blur",
  //   message: "请输入视频"
  // },
  category: {
    required: true,
    trigger: "blur",
    message: "请输文章分类",
    pattern: /^[a-zA-Z0-9_-]{2,40}$/ //验证文章分类
  },
  status: {
    required: true,
    trigger: "blur",
    message: "请选择文章状态",
    type: "enum",
    enum: ["published", "pending", "draft"]
  }
};
export default defineComponent({
  name: "CreateEditForm",
  components: {
    PostImgUpload,
    wangEditor
  },
  props: {
    formType: {
      type: String,
      default: ""
    }
  },
  setup(props, context) {
    const wangEditorRef = ref();
    const message = useMessage();
    const formBtnLoading = ref(false);
    const formRef: any = ref(null);
    const options = ref([]);
    const topicOptions = ref([]);
    const userList = ref<SelectOption[]>([]);
    const quillEditor = ref();
    const defaultValueRef = () => ({
      postId: "",
      title: "",
      exerpt: "",
      content: "hello",
      category: "",
      author: null,
      cover: "",
      postImg: [],
      type: "post",
      view: 0,
      videoRelation: [],
      musicRelation: [],
      from: "original",
      downloadRelation: [],
      status: "draft",
      topic: [] as string[],
      commentStatus: true
    });
    const renderAuthorSelectTag: SelectRenderTag = ({ option }) => {
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
            size: 24,
            style: {
              marginRight: "12px"
            }
          }),
          option.label as string
        ]
      );
    };
    const renderAuthorLabel: SelectRenderLabel = (option) => {
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
    const renderTopicSelectTag: SelectRenderTag = ({ option }) => {
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
            src: option.icon as string,
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
    const renderTopicLabel: SelectRenderLabel = (option) => {
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
            src: option.icon as string,
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
                  default: () => option.value as string
                }
              )
            ]
          )
        ]
      );
    };

    const uploadHeaders = reactive({
      platform: "miniPrograms",
      timestamp: new Date().getTime(),
      Authorization: "Bearer " + Storage.get(ACCESS_TOKEN, "")
    });

    async function getCategoryLists() {
      const {
        code,
        message: msg,
        result
      } = await getCategoryList({
        pageSize: 999,
        page: 1
      });
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        if (result.list) {
          options.value = result.list;
        }
      } else {
        message.error(msg || "获取失败");
      }
    }

    async function getTopicLists() {
      const {
        code,
        message: msg,
        result
      } = await getTopicList({
        pageSize: 999,
        page: 1
      });
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        topicOptions.value = result.list;
      } else {
        message.error(msg || "获取失败");
      }
    }

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

    const state = reactive({
      showForm: false,
      topic: [] as string[],
      formParams: defaultValueRef(),
      formType: ""
    });
    const postImgUpload = ref();

    function confirmForm(e) {
      e.preventDefault();
      formBtnLoading.value = true;
      formRef.value.validate(async (errors) => {
        if (!errors) {
          state.formParams.topic = state.topic;
          const { code, message: msg } =
            props.formType === "create"
              ? await createPost(state.formParams)
              : await updatePost(state.formParams);
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

    function onCreateMusic() {
      return {
        postId: state.formParams.postId,
        name: "",
        url: ""
      };
    }

    function onCreateVideo() {
      return {
        postId: state.formParams.postId,
        name: "",
        subtitles: "",
        url: ""
      };
    }

    function setCover(img: any) {
      state.formParams.cover = img.url;
    }

    function insertPostImg(img: any) {
      wangEditorRef.value.insertCover(img.url);
    }

    async function deletePostImg(file: any) {
      const { code, message: msg } = await deleteImgFile({
        "url": file.img.url
      });
      message.destroyAll();
      if (code == ResultEnum.SUCCESS) {
        message.success("删除成功");
        if (file.img.url == state.formParams.cover) {
          state.formParams.cover = state.formParams.postImg[0].url;
        }
        postImgUpload.value.deleteImg(file.index);
      } else {
        message.error(msg || "删除失败");
      }
    }

    function uploadChange(postImgs: []) {
      state.formParams.postImg.push(...postImgs);
    }

    watch(
      () => props.formType,
      () => {
        if (props.formType == "create") {
          state.topic = [];
        }
      }
    );

    function openForm(record) {
      if (record !== null) {
        state.formParams = record;
      }
      getCategoryLists();
      getTopicLists();
      getUserLists();
      state.topic = state.formParams.topic.map((item) => {
        return item.value;
      });
      state.showForm = true;
    }

    function closeForm() {
      Object.assign(state.formParams, defaultValueRef());
      state.showForm = false;
    }

    const userMeta = computed(() => Storage.get(CURRENT_USER));

    return {
      ...toRefs(state),
      formRef,
      rules,
      options,
      topicOptions,
      formBtnLoading,
      uploadUrl,
      uploadHeaders,
      userList,
      quillEditor,
      userMeta,
      GlobeOutline,
      onCreateMusic,
      onCreateVideo,
      renderAuthorLabel,
      renderAuthorSelectTag,
      renderTopicLabel,
      renderTopicSelectTag,
      setCover,
      insertPostImg,
      deletePostImg,
      openForm,
      closeForm,
      confirmForm,
      uploadChange,
      postImgUpload,
      wangEditorRef
    };
  }
});
</script>
