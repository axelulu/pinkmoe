/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-28 16:34:17
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 15:21:50
 * @FilePath: /pinkmoe_vitesse/src/hooks/util.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useAppStore } from '../store'

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
        msg,
      })
      return false
    }
    else {
      return true
    }
  }

  return { formatDate, getLevel, checkForm }
}
