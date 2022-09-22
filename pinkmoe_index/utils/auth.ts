/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-09-10 22:09:40
 * @FilePath: /pinkmoe_index/utils/auth.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
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
