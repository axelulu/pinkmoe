<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 15:46:56
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 12:28:13
 * @FilePath: /pinkmoe_admin/src/views/shop/goods/CreateEditForm.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
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
      <n-form-item label="内容" path="content">
        <wangEditor ref="wangEditorRef" v-model:value="formParams.content" />
      </n-form-item>
      <n-form-item label="图片" path="cover">
        <PostImgUpload
          ref="postImgUpload"
          :action="`${uploadUrl}/api/Admin/Upload/FileUpload`"
          :headers="uploadHeaders"
          :data="{'goods_id': formParams.goodsId, 'uuid': userMeta.uuid, 'type': 'goods'}"
          max-size="1"
          max-number="12"
          height="170"
          width="150"
          v-model:value="formParams.goodsImg"
          v-model:cover="formParams.cover"
          @setCover="setCover"
          @uploadChange="uploadChange"
          @insertPost="insertPostImg"
          @delete="deletePostImg"
        />
      </n-form-item>
      <n-form-item label="类型" path="type">
        <n-radio-group v-model:value="formParams.type" name="type">
          <n-radio value="physical"> 实物</n-radio>
          <n-radio value="virtual"> 虚拟</n-radio>
        </n-radio-group>
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
      <n-form-item label="商品规格" path="view">
        <n-dynamic-input v-model:value="formParams.sizeRelation" :on-create="onCreateSize">
          <template #create-button-default>
            添加商品规格
          </template>
          <template #default="{ value, index }">
            <div style="display: flex; align-items: center; width: 100%">
              <div class="w-1/6 mr-4">
                <n-input v-model:value="value.key" type="text" />
              </div>
              <n-dynamic-tags @update:value="(value: DynamicTagsOption[]) => updateSize(value, index)" v-model:value="value.value" />
            </div>
          </template>
        </n-dynamic-input>
      </n-form-item>
      <n-form-item label="规格对应值" path="view">
        <n-dynamic-input v-model:value="formParams.sizeValRelation" :on-create="onCreateSizeVal">
          <template #create-button-default>
            添加商品规格
          </template>
          <template #default="{ value }">
            <div style="display: flex; align-items: center; width: 100%">
              <n-dynamic-tags disabled="true" class="mr-4" v-model:value="value.value" />
              <div class="w-1/6 mr-4">
                <n-input-number v-model:value="value.saleAmount" type="text" placeholder="售价" />
              </div>
              <div class="w-1/6 mr-4">
                <n-input-number v-model:value="value.stock" type="text" placeholder="销量" />
              </div>
            </div>
          </template>
        </n-dynamic-input>
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
import { createGoods, updateGoods } from "@/api/shop/goods";
import { getShopCategoryList } from "@/api/shop/category";
import { useGlobSetting } from "@/hooks/setting";
import { ACCESS_TOKEN, CURRENT_USER } from "@/store/mutation-types";
import { PostImgUpload } from "@/components/UploadPostImg";
import { createStorage } from "@/utils/Storage";
import { getUserList } from "@/api/system/user";
import wangEditor from "./wangEditor.vue";
import "@vueup/vue-quill/dist/vue-quill.snow.css";
import { deleteImgFile } from "@/api/upload/upload";
import { DynamicTagsOption } from "naive-ui/lib/dynamic-tags/src/interface";

const globSetting = useGlobSetting();
const { uploadUrl } = globSetting;
const Storage = createStorage({ storage: localStorage });
const onCreateSize = () => {
  return {
    key: '',
    value: []
  }
}
const onCreateSizeVal = () => {
  return {
    stock: '',
    saleAmount: '',
    value: []
  }
}
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
    enum: ["physical", "virtual"],
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
    const userList = ref<SelectOption[]>([]);
    const quillEditor = ref();
    const defaultValueRef = () => ({
      goodsId: "",
      title: "",
      content: "hello",
      category: "",
      author: null,
      cover: "",
      goodsImg: [],
      type: "physical",
      view: 0,
      sizeRelation: [],
      sizeValRelation: [],
      status: "draft",
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
      } = await getShopCategoryList({
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
          const { code, message: msg } =
            props.formType === "create"
              ? await createGoods(state.formParams)
              : await updateGoods(state.formParams);
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
        postId: state.formParams.goodsId,
        name: "",
        url: ""
      };
    }

    function onCreateVideo() {
      return {
        postId: state.formParams.goodsId,
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
          state.formParams.cover = state.formParams.goodsImg[0].url;
        }
        postImgUpload.value.deleteImg(file.index);
      } else {
        message.error(msg || "删除失败");
      }
    }

    function uploadChange(postImgs: []) {
      console.log(postImgs)
      state.formParams.goodsImg.push(...postImgs);
      console.log(state.formParams.goodsImg)
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
      getUserLists();
      state.formParams.sizeRelation.forEach((e)=>{
        e.value = JSON.parse(e.value)
      })
      state.formParams.sizeValRelation.forEach((e)=>{
        e.value = JSON.parse(e.value)
      })
      state.showForm = true;
    }

    function closeForm() {
      Object.assign(state.formParams, defaultValueRef());
      state.showForm = false;
    }

    const userMeta = computed(() => Storage.get(CURRENT_USER));

    const getDimComb = (doubleList = []) => {
      if (doubleList.length == 0) return [];//先return掉空的
      
      const result = [];//最终组合集
      /**
      * doubleList 二维数组 Array
      * index 当前所在步骤
      * currentList 当前累积维度的数组
      */
      const _back = (doubleList, index, currentList) => {
        //判断是否到了最底层
        if (index >= doubleList.length) {
          result.push([...currentList]);
        } else {
          //循环递归
          if (doubleList[index].value) {
            doubleList[index].value.forEach(item => {
              //累加当前数组
              currentList.push(item);
              //递归
              _back(doubleList, index + 1, currentList)
              currentList.pop();
            });
          }
        }
      }
      _back(doubleList, 0, []);
      return result;
    }

    const updateSize = (value, index) => {
      state.formParams.sizeRelation[index].value = value;
      const result = getDimComb(state.formParams.sizeRelation)
      let size = [];
      result.forEach(element => {
          size.push({
            stock: 0,
            saleAmount: 0,
            value: element,
          })
      });
      state.formParams.sizeValRelation = size;
    }

    return {
      ...toRefs(state),
      formRef,
      rules,
      options,
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
      onCreateSize,
      onCreateSizeVal,
      updateSize,
      postImgUpload,
      wangEditorRef
    };
  }
});
</script>
