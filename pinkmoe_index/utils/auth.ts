/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 22:09:40
 * @FilePath: /pinkmoe_index/utils/auth.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
const TokenKey = 'pink-token'
const TokenPrefix = 'Bearer '
const isLogin = () => {
  const { token } = useAuthStorage()
  return !!token.value
}
const getToken = () => {
  const { token } = useAuthStorage()
  return token.value
}
const setToken = (token: string) => {
  const { store } = useAuthStorage()
  store(token)
}
const clearToken = () => {
  const { clear } = useAuthStorage()
  clear()
}
export { TokenPrefix, TokenKey, isLogin, getToken, setToken, clearToken }
