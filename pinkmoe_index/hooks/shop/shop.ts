/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 08:40:11
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:48:42
 * @FilePath: /pinkmoe_vitesse/src/hooks/shop/shop.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import type { ResCategory } from '/@/api/category/types'
import type { ResPost } from '/@/api/home/types'
import { getShopList } from '/@/api/shop'
import type { ResContent } from '/@/api/shop/types'

export const useShop = () => {
  const content = ref<Array<ResContent>>()
  const popular = ref<Array<ResPost>>()
  const shopCategory = ref<Array<ResCategory>>()
  const loading = ref(false)

  // 获取首页
  const getHome = async () => {
    loading.value = true
    const homeList = await getShopList()
    content.value = homeList.content
    popular.value = homeList.popular
    shopCategory.value = homeList.shopCategory
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const shopMenu = ref()
  const scrollMenu = (isRight) => {
    shopMenu.value.scrollTo({
      // y方向坐标800px（代码中不需要写明单位）
      left: isRight ? shopMenu.value.scrollLeft + 140 : shopMenu.value.scrollLeft - 140,
      // 滚动方式是平滑滚动 默认是auto 即instant 直接跳到目标位置  不常用
      behavior: 'smooth',
    })
  }

  onMounted(() => {
    getHome()
  })

  return {
    scrollMenu,
    shopMenu,
    content,
    popular,
    shopCategory,
    loading,
  }
}
