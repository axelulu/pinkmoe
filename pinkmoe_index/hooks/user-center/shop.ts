/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-04 07:27:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:59:56
 * @FilePath: /pinkmoe_index/hooks/user-center/shop.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

import { useAppStore } from "/@/store/modules/app"
import { useUserStore } from "/@/store/modules/user"

export const useUserCenterShop = () => {
  const loading = ref(true)
  const currentIndex = ref(0)
  const auth = useUserStore()
  const { proxy } = getCurrentInstance()
  const { userShop } = useAppStore()

  function buyShop() {
    proxy.$shopConfirm({
      authorShop: userShop[currentIndex.value],
    })
  }

  onMounted(() => {
    loading.value = true
    setTimeout(() => {
      loading.value = false
    }, 300)
  })

  return {
    loading,
    currentIndex,
    auth,
    userShop,
    buyShop,
  }
}
