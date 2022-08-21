/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 10:47:46
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 15:43:16
 * @FilePath: /pinkmoe_vitesse/src/hooks/topic.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
