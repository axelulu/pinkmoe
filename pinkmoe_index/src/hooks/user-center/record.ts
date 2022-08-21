/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 22:32:30
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:17
 * @FilePath: /pinkmoe_vitesse/src/hooks/user-center/record.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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

    console.log(authorRecordList.value)
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
