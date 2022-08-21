/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 18:13:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:49
 * @FilePath: /pinkmoe_vitesse/src/hooks/author/comment.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ResPage } from '/@/api/common/types'
import { getAuthorCommentList } from '/@/api/author'
import type { ResComment } from '/@/api/post/types'
import type { ReqCommentPost } from '/@/api/comment/types'

export const useAuthorComment = () => {
  const authorCommentList = ref<ResPage<Array<ResComment>>>()
  const route = useRoute()
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const formParams: ReqCommentPost = reactive({
    formUid: route.params.author,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
  })

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getAuthorCommentList(formParams)
    if (!res.list || res.list.length <= 0)
      hasMore.value = false

    else
      authorCommentList.value?.list?.push(...res.list!)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  // 获取评论列表
  const getAuthorComment = async () => {
    loading.value = true
    authorCommentList.value = await getAuthorCommentList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getAuthorComment()
  })

  return {
    authorCommentList,
    loading,
    hasMore,
    formParams,
    nextPage,
  }
}
