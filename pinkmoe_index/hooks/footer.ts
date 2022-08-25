/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 13:20:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 19:07:59
 * @FilePath: /pinkmoe_index/hooks/footer.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useAppStore } from '/@/store'

export const useFooter = () => {
  const appStore = useAppStore()
  function scollTop() {
    //  设置一个定时器
    const upRoll = setInterval(() => {
      if (process.client) {
        const top = document.documentElement.scrollTop // 每次获取页面被卷去的部分
        const speed = top / 10 // 每次滚动多少 （步长值）
        if (document.documentElement.scrollTop !== 0)
          document.documentElement.scrollTop -= speed // 不在顶部 每次滚动到的位置
        else
          clearInterval(upRoll) // 回到顶部清除定时器
      }
    }, 20)
  }
  return { scollTop, appStore }
}
