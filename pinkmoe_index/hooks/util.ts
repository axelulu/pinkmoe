/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-28 16:34:17
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:18:53
 * @FilePath: /pinkmoe_index/hooks/util.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { BASEURL } from '../constant'
import { useAppStore } from '../store/modules/app'

export const useUtil = () => {
  const { userLevel } = useAppStore()
  const { proxy } = getCurrentInstance()

  const formatDate = (dateString?: string | Date) => {
    const time = Date.now() - new Date(dateString)
    /*
    1秒 1000
    1分钟 1000 * 60 60000
    1小时 1000 * 60 * 60 3600000
    1天 1000 * 60 * 60 * 24 86400000
    1个月 1000 * 60 * 60 * 24 * 30 2592000000
    1年 1000 * 60 * 60 * 24 * 365 31536000000
  */

    const times = [31536000000, 2592000000, 86400000, 3600000, 60000, 1000]
    const units = ['年前', '月前', '天前', '小时前', '分钟前', '秒前']

    for (let i = 0; i < times.length; i++) {
      const res = time / times[i]
      if (res >= 1)
        return Math.round(res) + units[i]
    }
  }

  const getLevel = (exp) => {
    if (userLevel) {
      for (const lev of userLevel) {
        if (exp >= lev.headExp && exp < lev.footExp)
          return lev
      }
    }
    else {
      return null
    }
  }

  const checkForm = (value, regular, msg) => {
    const email = RegExp(regular)
    if (!email.test(<string>value)) {
      proxy.$message({
        type: 'warning',
        successMsg: msg,
        loadFun: async () => {
          const code = 200
          return { code }
        },
      })
      return false
    }
    else {
      return true
    }
  }

  const getThumbnail = (url: string) => {
    return `${BASEURL}/Cms/Images/ImageThumbnail?width=${200}&height=${300}&image=${url}`
  }

  return { formatDate, getThumbnail, getLevel, checkForm }
}
