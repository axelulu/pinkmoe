/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 09:57:35
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:22
 * @FilePath: /pinkmoe_vitesse/src/hooks/user-center/vip.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ReqAuthorPost } from '/@/api/author/types'
import { useUserStore } from '/@/store'
import { getAuthorityList } from '/@/api/user'

export const useUserCenterVip = () => {
  const loading = ref(true)
  const currentIndex = ref(0)
  const authorAuthorityList = ref<any>([])
  const auth = useUserStore()
  const { proxy } = getCurrentInstance()
  const formParams: ReqAuthorPost = reactive({
    page: 1,
    pageSize: 12,
  })

  async function getUserAuthorityList() {
    loading.value = true
    authorAuthorityList.value = await getAuthorityList(formParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  async function buyVip() {
    proxy.$vipConfirm({
      authorAuthority: authorAuthorityList.value?.list[currentIndex.value],
    })
  }

  onMounted(() => {
    getUserAuthorityList()
  })

  return {
    authorAuthorityList,
    loading,
    currentIndex,
    auth,
    buyVip,
  }
}
