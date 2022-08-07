/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:07:55
 * @FilePath: /pinkmoe_index/src/hooks/search.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { ResCategory } from '/@/api/category/types';
import { getCategoryList } from '/@/api/category';
import { getSearchPostList } from '/@/api/search';
import { ReqSearchPost, SearchSlug } from '/@/api/search/types';
import { ResPage } from '/@/api/common/types';
import { ResPost } from '/@/api/home/types';
import { useAppStore } from '../store/modules/app';

export const useSearch = () => {
  const loading = ref<boolean>(false);
  const hasMore = ref<boolean>(true);
  const isShow = ref<boolean>(false);
  const { siteBasic } = useAppStore();
  const router = useRouter();
  const currentCategory = ref<SearchSlug>({
    slug: '0',
    type: 'all',
  });
  const children = ref<any>(null);
  const searchReCategory = ref();
  const searchPostList = ref<ResPage<Array<ResPost>>>();
  const sort = ref<any>([
    {
      title: '按最新',
      type: 'updated_at',
    },
    {
      title: '按标题',
      type: 'title',
    },
    {
      title: '按作者',
      type: 'author',
    },
    {
      title: '按查看',
      type: 'view',
    },
  ]);
  const route = useRoute();

  const formParams: ReqSearchPost = reactive({
    category: route.params.slug,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  });

  const categoryList = ref<Array<ResCategory>>();

  // 获取分类
  const getCategory = async () => {
    const categories = await getCategoryList();
    categoryList.value = categories.list;
  };

  // 获取搜索文章
  const getSearchPost = async () => {
    loading.value = true;
    searchPostList.value = await getSearchPostList(formParams);
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  formParams.title = route.query.keyword as string;

  const jumpTo = (key: string) => {
    formParams.title = key;
    router.push({
      path: '/search',
      query: {
        keyword: key,
      },
    });
  };

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++;
    const res = await getSearchPostList(formParams);
    if (!res.list || res.list.length <= 0) {
      hasMore.value = false;
    } else {
      searchPostList.value?.list?.push(...res.list!);
    }
    setTimeout(() => {
      loading.value = false;
    }, 300);
  };

  const sortPost = async (type: string, descs: boolean) => {
    formParams.orderKey = type;
    formParams.desc = descs;
    formParams.page = 1;
    hasMore.value = true;
    getSearchPost();
  };

  const getChildCategory = (item) => {
    isShow.value = true;
    children.value = item;
    if (searchReCategory.value) {
      searchReCategory.value.hide();
    }
    getCategoryPost(item);
  };

  const getCategoryPost = (item) => {
    hasMore.value = true;
    formParams.page = 1;
    currentCategory.value = item;
    formParams.category = item.slug;
    getSearchPost();
  };

  onMounted(async () => {
    getCategory();
    getSearchPost();
  });

  return {
    categoryList,
    jumpTo,
    formParams,
    sort,
    route,
    siteBasic,
    sortPost,
    searchPostList,
    loading,
    hasMore,
    isShow,
    children,
    currentCategory,
    searchReCategory,
    getChildCategory,
    getCategoryPost,
    nextPage,
  };
};
