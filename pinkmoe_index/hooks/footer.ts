/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 13:20:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-30 11:06:17
 * @FilePath: /pinkmoe_index/hooks/footer.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */

import { useAppStore } from '../store/modules/app'

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
