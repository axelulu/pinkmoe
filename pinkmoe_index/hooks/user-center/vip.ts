/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 09:57:35
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:59:05
 * @FilePath: /pinkmoe_index/hooks/user-center/vip.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { ReqAuthorPost } from '/@/api/author/types'
import { useUserStore } from '/@/store/modules/user';
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
