/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 14:35:57
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:58:49
 * @FilePath: /pinkmoe_index/hooks/user-center/comments.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { getUserComment } from '/@/api/comment'
import type { ReqAuthorPost } from '/@/api/author/types'
import { useUserStore } from '/@/store/modules/user';
import type { ResPost } from '/@/api/home/types'

export const useUserCenterComments = () => {
  const loading = ref(true)
  const hasMore = ref(true)
  const authorCommentList = ref<any>([])
  const auth = useUserStore()
  const formParams: ReqAuthorPost = reactive({
    page: 1,
    pageSize: 12,
    formUid: auth.userInfo?.uuid,
  })

  async function getUserCommentList() {
    loading.value = true
    authorCommentList.value = await getUserComment(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  async function nextPage() {
    loading.value = true;
    (formParams.page as number)++
    const res = await getUserComment(formParams)
    if (!res.list || res.list?.length <= 0)
      hasMore.value = false

    else
      authorCommentList.value?.list?.push(...(res?.list as Array<ResPost>))

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getUserCommentList()
  })

  return {
    authorCommentList,
    hasMore,
    loading,
    nextPage,
  }
}
