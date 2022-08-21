/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:47:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:55
 * @FilePath: /pinkmoe_vitesse/src/hooks/author/post.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ReqCategoryPost } from '/@/api/category/types'
import type { ResPage } from '/@/api/common/types'
import type { ResPost } from '/@/api/home/types'
import { getAuthorPostList } from '/@/api/author'

export const useAuthorPost = () => {
  const authorPostList = ref<ResPage<Array<ResPost>>>()
  const route = useRoute()
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const formParams: ReqCategoryPost = reactive({
    author: route.params.author,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
  })

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getAuthorPostList(formParams)
    if (!res.list || res.list.length <= 0)
      hasMore.value = false

    else
      authorPostList.value?.list?.push(...res.list!)

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
    authorPostList.value = await getAuthorPostList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getCategoryPost()
  })

  return {
    authorPostList,
    loading,
    hasMore,
    formParams,
    nextPage,
    sortPost,
  }
}
