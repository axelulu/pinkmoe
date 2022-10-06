/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:47:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 16:24:35
 * @FilePath: /pinkmoe_index/hooks/topic.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { useAppStore } from '../store/modules/app'
import type { ResPage } from '/@/api/common/types'
import { getTopicPostList } from '/@/api/topic'
import type { ReqTopicPost, ResTopicPost } from '/@/api/topic/types'

export const useTopic = () => {
  const topicPostList = ref<ResPage<ResTopicPost>>()
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
  const formParams: ReqTopicPost = reactive({
    value: route.params.slug,
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
  })

  const nextPage = async () => {
    loading.value = true;
    (formParams.page as number)++
    const res = await getTopicPostList(formParams)
    if (!res?.list?.post || res?.list?.post?.length <= 0)
      hasMore.value = false

    else
      topicPostList.value?.list?.post?.push(...res.list.post!)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const sortPost = async (type: string, descs: boolean) => {
    formParams.orderKey = type
    formParams.desc = descs
    formParams.page = 1
    getTopicPost()
    hasMore.value = true
  }

  // 获取话题文章列表
  const getTopicPost = async () => {
    loading.value = true
    topicPostList.value = await getTopicPostList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getTopicPost()
  })

  return {
    topicPostList,
    sort,
    loading,
    hasMore,
    formParams,
    siteBasic,
    route,
    nextPage,
    sortPost,
  }
}
