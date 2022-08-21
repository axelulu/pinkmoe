/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 21:11:41
 * @FilePath: /pinkmoe_vitesse/src/utils/auth.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
const TokenKey = 'pink-token'
const TokenPrefix = 'Bearer '
const isLogin = () => {
  if (!import.meta.env.SSR)
    return !!localStorage.getItem(TokenKey)
}
const getToken = () => {
  if (!import.meta.env.SSR)
    return localStorage.getItem(TokenKey)
}
const setToken = (token: string) => {
  if (!import.meta.env.SSR)
    localStorage.setItem(TokenKey, token)
}
const clearToken = () => {
  if (!import.meta.env.SSR)
    localStorage.removeItem(TokenKey)
}
export { TokenPrefix, isLogin, getToken, setToken, clearToken }
