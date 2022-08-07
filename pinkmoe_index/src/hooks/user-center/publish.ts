/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 14:33:05
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 11:13:17
 * @FilePath: /pinkmoe_index/src/hooks/user-center/publish.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useUserStore } from '/@/store';
import { getCategoryList } from '/@/api/category';
import { deleteFile, fileList, upload } from '/@/api/upload';
import { ReqFileList } from '/@/api/upload/types';
import { createPost } from '/@/api/post';
import { ReqPublishPost } from '/@/api/post/types';
import { useUtil } from '../util';

export const useUserCenterPublish = () => {
  const categoryList = ref<any>();
  const wangEditor = ref();
  const message = ref();
  const formPublish = ref();
  const publishCategory = ref();
  const { checkForm } = useUtil();
  const auth = useUserStore();
  const router = useRouter();
  const { proxy } = getCurrentInstance();
  const category = ref<any>({
    name: '',
    slug: '',
  });
  const formParams: ReqPublishPost = reactive({
    postId: '',
    title: '',
    exerpt: '',
    content: 'hello',
    category: '',
    author: auth.userInfo?.uuid,
    cover: '',
    postImg: <any>[],
    type: 'post',
    view: 0,
    videoRelation: [
      {
        name: '',
        url: '',
        subtitles: '',
        price: 0,
        priceType: 'credit',
      },
    ],
    musicRelation: [
      {
        name: '',
        url: '',
        price: 0,
        priceType: 'credit',
      },
    ],
    from: 'original',
    downloadRelation: [
      {
        name: '',
        url: '',
        unpackPwd: '',
        extractPwd: '',
        price: 0,
        priceType: 'credit',
      },
    ],
    status: 'pending',
    topic: [] as string[],
    commentStatus: 'true',
  });

  const fileFormParams: ReqFileList = reactive({
    page: 1,
    pageSize: 999,
    type: 'post_temporary',
  });

  async function getCategory() {
    categoryList.value = await getCategoryList();
    publishCategory.value.show();
  }

  async function getImgList() {
    formParams.postImg = (await fileList(fileFormParams)).list;
  }

  onMounted(() => {
    getImgList();
  });

  function chooseCategory(item: { slug: any; name: any }) {
    category.value.slug = item.slug;
    category.value.name = item.name;
    formParams.category = item.slug;
    publishCategory.value.hide();
  }

  async function uploadPostImg(event: any) {
    const fd = new FormData();
    fd.append('file', event.target.files[0]);
    fd.append('uuid', auth.userInfo?.uuid as string);
    fd.append('type', 'post_temporary');
    const { code, message, result } = await upload(fd);
    if (code === 200) {
      proxy.$message({
        type: 'success',
        msg: '上传成功',
      });
      formParams.postImg?.push(result.file!);
      if (formParams.cover?.length === 0) {
        formParams.cover = formParams.postImg?.[0].url;
      }
    } else {
      proxy.$message({
        type: 'warning',
        msg: message || '上传失败',
      });
    }
    event.target.value = '';
  }

  function insertPostImg(img: any) {
    wangEditor.value.insertCover('/' + img.url);
  }

  async function deletePostImg(index: any) {
    const { code, message } = await deleteFile({
      url: formParams.postImg?.[index].url,
    });
    if (code === 200) {
      if (formParams.postImg?.[index].url === formParams.cover) {
        formParams.cover = formParams.postImg?.[0].url;
      }
      formParams.postImg?.splice(index, 1);
      proxy.$message({
        type: 'success',
        msg: '删除成功',
      });
    } else {
      proxy.$message({
        type: 'warning',
        msg: message || '删除失败',
      });
    }
  }

  async function publishPost() {
    if (
      checkForm(
        formParams.title,
        '[\u4e00-\u9fa5]{1,50}$|^[\\dA-Za-z_]{1,50}$',
        '请输入正确的文章标题格式',
      ) &&
      checkForm(
        formParams.exerpt,
        '[\u4e00-\u9fa5]{1,120}$|^[\\dA-Za-z_]{1,120}$',
        '请输入正确的文章摘要格式',
      ) &&
      checkForm(
        formParams.content,
        '[\u4e00-\u9fa5]{1,120}$|^[\\dA-Za-z_]{1,120000}$',
        '请输入正确的文章内容格式',
      )
    ) {
      const { code, message } = await createPost(formParams);
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '发布成功',
        });
        setTimeout(() => {
          router.push('/user-center/posts');
        }, 1000);
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '发布失败',
        });
      }
    }
  }

  return {
    category,
    categoryList,
    publishCategory,
    wangEditor,
    message,
    formParams,
    getCategory,
    publishPost,
    chooseCategory,
    uploadPostImg,
    insertPostImg,
    deletePostImg,
    formPublish,
  };
};
