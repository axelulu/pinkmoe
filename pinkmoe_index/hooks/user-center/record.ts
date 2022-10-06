/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 22:32:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 19:02:04
 * @FilePath: /pinkmoe_index/hooks/user-center/record.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { ReqAuthorPost } from '/@/api/author/types'
import { getNotificationList } from '/@/api/notification'

export const useUserCenterRecord = () => {
  const loading = ref(true)
  const hasMore = ref(true)
  const authorRecordList = ref<any>([])
  const formParams: ReqAuthorPost = reactive({
    page: 1,
    pageSize: 12,
  })

  async function getUserRecordList() {
    loading.value = true
    authorRecordList.value = await getNotificationList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  async function nextPage() {
    loading.value = true;
    (formParams.page as number)++
    const res = await getNotificationList(formParams)
    if (!res.list || res.list?.length <= 0)
      hasMore.value = false

    else
      authorRecordList.value?.list?.push(...res?.list)

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  onMounted(() => {
    getUserRecordList()
  })

  return {
    authorRecordList,
    hasMore,
    loading,
    nextPage,
  }
}
