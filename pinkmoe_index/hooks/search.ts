/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 16:24:31
 * @FilePath: /pinkmoe_index/hooks/search.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { ResCategory } from '/@/api/category/types'
import { getCategoryList } from '/@/api/category'
import { getSearchPostList } from '/@/api/search'
import type { ReqSearchPost, SearchSlug } from '/@/api/search/types'
import type { ResPage } from '/@/api/common/types'
import type { ResPost } from '/@/api/home/types'
import { useAppStore } from '../store/modules/app'

export const useSearch = () => {
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const isShow = ref<boolean>(false)
  const { siteBasic } = useAppStore()
  const router = useRouter()
  const currentCategory = ref<SearchSlug>({
    slug: '0',
    type: 'all',
  })
  const children = ref<any>(null)
  const searchReCategory = ref()
  const searchPostList = ref<ResPage<Array<ResPost>>>()
  const sort = ref<any>([
    {
      title: 'accordingToLatest',
      type: 'updated_at',
    },
    {
      title: 'accordingToTitle',
      type: 'title',
    },
    {
      title: 'accordingToAuthor',
      type: 'author',
    },
    {
      title: 'accordingToView',
      type: 'view',
    },
  ])
  const route = useRoute()

  const formParams: ReqSearchPost = reactive({
    category: route.params.slug,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  })

  const categoryList = ref<Array<ResCategory>>()

  // 获取分类
  const getCategory = async () => {
    const categories = await getCategoryList()
    categoryList.value = categories.list
  }

  // 获取搜索文章
  const getSearchPost = async () => {
    loading.value = true
    searchPostList.value = await getSearchPostList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  formParams.title = route.query.keyword as string

  const jumpTo = (key: string) => {
    formParams.title = key
    router.push({
      path: '/search',
      query: {
        keyword: key,
        t: Date.parse(new Date().toString()),
      },
    })
  }

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getSearchPostList(formParams)
    if (!res.list || res.list.length <= 0)
      hasMore.value = false

    else
      searchPostList.value?.list?.push(...res.list!)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const sortPost = async (type: string, descs: boolean) => {
    formParams.orderKey = type
    formParams.desc = descs
    formParams.page = 1
    hasMore.value = true
    getSearchPost()
  }

  const getChildCategory = (item) => {
    isShow.value = true
    children.value = item
    if (searchReCategory.value)
      searchReCategory.value.hide()

    getCategoryPost(item)
  }

  const getCategoryPost = (item) => {
    hasMore.value = true
    formParams.page = 1
    currentCategory.value = item
    formParams.category = item.slug
    getSearchPost()
  }

  onMounted(async () => {
    getCategory()
    getSearchPost()
  })

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
  }
}
