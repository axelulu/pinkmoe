/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 08:40:11
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 11:09:04
 * @FilePath: /pinkmoe_index/src/hooks/bbs/bbs.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useI18n } from 'vue-i18n';
import { useUtil } from '../util';
import { ResPage } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';
import { createPost, getPostList } from '/@/api/post';
import { ReqBbsActive, ReqBbsPost } from '/@/api/post/types';
import { useUserStore } from '/@/store';

export const useBbs = () => {
  const postList = ref<ResPage<Array<ResPost>>>();
  const { proxy } = getCurrentInstance();
  const { t } = useI18n();
  const sort = ref<any>([
    {
      title: t('accordingToLatest'),
      type: 'updated_at',
    },
    {
      title: t('accordingToTitle'),
      type: 'title',
    },
    {
      title: t('accordingToAuthor'),
      type: 'author',
    },
    {
      title: t('accordingToView'),
      type: 'view',
    },
  ]);
  const leftBar = ref([
    {
      name: t('bbs.focus'),
      slug: 'focus',
    },
    {
      name: t('bbs.all'),
      slug: 'all',
    },
    {
      name: t('bbs.hot'),
      slug: 'hot',
    },
    {
      name: t('bbs.active'),
      slug: 'active',
    },
    {
      name: t('bbs.post'),
      slug: 'post',
    },
    {
      name: t('bbs.video'),
      slug: 'video',
    },
    {
      name: t('bbs.music'),
      slug: 'music',
    },
    {
      name: t('bbs.download'),
      slug: 'download',
    },
  ]);
  const auth = useUserStore();
  const route = useRoute();
  const router = useRouter();
  const formParams: ReqBbsPost = reactive({
    userId: auth?.userInfo?.uuid,
    category: '',
    type: '',
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  });

  const postFormParams: ReqBbsActive = reactive({
    title: '',
    content: '',
    author: auth.userInfo?.uuid,
    type: 'active',
  });
  const loading = ref<boolean>(false);
  const isTitle = ref<boolean>(false);
  const formPublishRef = ref();
  const currentSlug = ref('all');
  const { checkForm } = useUtil();
  const currentSort = ref('updated_at');
  const hasMore = ref<boolean>(true);

  // 获取分类文章列表
  const getPost = async () => {
    loading.value = true;
    postList.value = await getPostList(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  const sortPost = (type) => {
    formParams.page = 1;
    hasMore.value = true;
    if (type === 'video' || type === 'post' || type === 'music' || type === 'active') {
      formParams.author = '';
      currentSlug.value = type;
      currentSort.value = 'updated_at';
      formParams.orderKey = 'updated_at';
      formParams.type = type;
    } else if (type === 'focus') {
      formParams.author = auth?.userInfo?.uuid;
      currentSlug.value = type;
      currentSort.value = 'updated_at';
      formParams.orderKey = 'updated_at';
      formParams.type = type;
      formParams.orderKey = 'view';
    } else if (type === 'hot') {
      formParams.author = '';
      currentSort.value = 'updated_at';
      currentSlug.value = type;
      formParams.type = '';
      formParams.orderKey = 'view';
    } else if (type === 'view' || type === 'author' || type === 'title' || type === 'updated_at') {
      formParams.author = '';
      currentSort.value = type;
      formParams.orderKey = type;
    } else if (type === 'desc') {
      formParams.author = '';
      formParams.desc = !formParams.desc;
    } else if (type === 'all') {
      formParams.author = '';
      currentSort.value = 'updated_at';
      currentSlug.value = type;
      formParams.type = '';
      formParams.orderKey = 'updated_at';
    }
    getPost();
  };

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++;
    formParams.author = auth?.userInfo?.uuid;
    const res = await getPostList(formParams);
    if (!res.list || res.list?.length <= 0) {
      hasMore.value = false;
    } else {
      postList.value?.list?.push(...(res?.list as Array<ResPost>));
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  function showLogin() {
    proxy.$login({
      type: 'login',
      router: router,
      route: route,
    });
  }

  async function publishPost() {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router: router,
        route: route,
      });
      return;
    }
    if (
      checkForm(
        postFormParams.content,
        '[\u4e00-\u9fa5]{1,50}$|^[\\dA-Za-z_]{1,50}$',
        '请输入正确的内容格式',
      )
    ) {
      const { code, message } = await createPost(postFormParams);
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '发布成功',
        });
        getPost();
      } else {
        proxy.$message({
          type: 'warning',
          msg: message || '发布失败',
        });
      }
    }
  }

  onMounted(() => {
    getPost();
  });

  return {
    postList,
    sort,
    formParams,
    currentSlug,
    currentSort,
    leftBar,
    hasMore,
    loading,
    isTitle,
    postFormParams,
    publishPost,
    sortPost,
    formPublishRef,
    nextPage,
    auth,
    showLogin,
    t,
  };
};
