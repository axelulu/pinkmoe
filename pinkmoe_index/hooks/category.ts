/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 09:02:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-30 11:24:34
 * @FilePath: /pinkmoe_index/hooks/category.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useAppStore } from '../store/modules/app'
import { getCategoryPostList } from '/@/api/category'
import type { ReqCategoryPost, ResCategoryPost } from '/@/api/category/types'
import type { ResPage } from '/@/api/common/types'

export const useCategory = async () => {
  const categoryPostList = ref<ResPage<ResCategoryPost>>()
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
    const res = await getCategoryPostList(formParams)
    if (!res.list?.post || res.list?.post.length <= 0)
      hasMore.value = false

    else
      categoryPostList.value?.list?.post?.push(...res?.list?.post)

    setTimeout(() => {
      loading.value = false
      scollBottom(648 * (formParams.page - 1))
    }, 300)
  }

  // 获取分类文章列表
  const getCategoryPost = async () => {
    loading.value = true
    categoryPostList.value = await getCategoryPostList(formParams)
    title.value = categoryPostList.value.list?.category?.[0].name as string
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

  function scollBottom(height) {
    //  设置一个定时器
    const upRoll = setInterval(() => {
      if (process.client) {
        const top = height // 每次获取页面被卷去的部分
        const speed = top / 10 // 每次滚动多少 （步长值）
        if (document.documentElement.scrollTop < height)
          document.documentElement.scrollTop += speed // 不在顶部 每次滚动到的位置
        else
          clearInterval(upRoll) // 回到顶部清除定时器
      }
    }, 20)
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
