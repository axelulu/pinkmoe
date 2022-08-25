/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:02:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 18:59:54
 * @FilePath: /pinkmoe_index/hooks/shop/category.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useAppStore } from '../../store/modules/app'
import { getCategoryGoodsList } from '/@/api/category'
import type { ReqCategoryPost, ResCategoryGoods } from '/@/api/category/types'
import type { ResPage } from '/@/api/common/types'

export const useShopCategory = async () => {
  const categoryPostList = ref<ResPage<ResCategoryGoods>>()
  const title = ref<string>('')
  const route = useRoute()
  const { siteBasic } = useAppStore()
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
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const formParams: ReqCategoryPost = reactive({
    category: route.params.slug,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  })

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getCategoryGoodsList(formParams)
    if (!res.list?.post || res.list?.post.length <= 0)
      hasMore.value = false

    else
      categoryPostList.value?.list?.post?.push(...res?.list?.post)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const sortPost = async (type: string, descs: boolean) => {
    formParams.orderKey = type
    formParams.desc = descs
    formParams.page = 1
    getCategoryPost()
    hasMore.value = true
  }

  // 获取分类文章列表
  const getCategoryPost = async () => {
    loading.value = true
    categoryPostList.value = await getCategoryGoodsList(formParams)
    title.value = categoryPostList.value.list?.category?.[0].name as string
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  await getCategoryPost()

  return {
    categoryPostList,
    sort,
    loading,
    hasMore,
    formParams,
    route,
    title,
    siteBasic,
    nextPage,
    sortPost,
  }
}
