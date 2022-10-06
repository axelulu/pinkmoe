/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 18:13:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:52
 * @FilePath: /pinkmoe_vitesse/src/hooks/author/follow.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { ResPage } from '/@/api/common/types'
import { getAuthorFollowList } from '/@/api/author'
import type { ReqAuthorFollow, ResFollow } from '/@/api/author/types'

export const useAuthorFollow = () => {
  const authorFollowList = ref<ResPage<Array<ResFollow>>>()
  const route = useRoute()
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const formParams: ReqAuthorFollow = reactive({
    author: '',
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
  })

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getAuthorFollowList(formParams)
    if (!res.list || res.list.length <= 0)
      hasMore.value = false

    else
      authorFollowList.value?.list?.push(...res.list!)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  // 获取关注列表
  const getAuthorFollow = async () => {
    loading.value = true
    formParams.formUid = route.params.author
    authorFollowList.value = await getAuthorFollowList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getAuthorFollow()
  })

  return {
    authorFollowList,
    loading,
    hasMore,
    formParams,
    nextPage,
  }
}
