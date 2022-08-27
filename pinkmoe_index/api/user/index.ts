/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:09:30
 * @FilePath: /pinkmoe_index/src/api/user/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { ResAuthor } from '/@/api/home/types'
import type {
  ReqUserAvatar,
  ReqUserBuyShop,
  ReqUserBuyVip,
  ReqUserDetail,
  ReqUserEmail,
  ReqUserEmailCaptcha,
  ReqUserPwd,
  ResCaptcha,
  ResCheckIn,
  ResLogin,
  ResUserProfile,
} from '/@/api/user/types'
import type { IResponse } from '/@/utils/http/axios/type'
import type { ReqAuthorPost } from '../author/types'
// import axios from 'axios';
enum URL {
  userInfo = '/Api/User/UserInfo',
  captcha = '/Api/User/Captcha',
  login = '/Api/User/Login',
  reg = '/Api/User/Reg',
  forget = '/Api/User/PwdForget',
  checkIn = '/Api/User/UserCheckIn',
  checkInStatus = '/Api/User/UserCheckInStatus',
  logout = '/Api/User/Logout',
  profile = '/Api/User/User',
  detail = '/Api/User/UserDetailUpdate',
  pwd = '/Api/User/UserPwdUpdate',
  email = '/Api/User/UserEmailUpdate',
  emailCaptcha = '/Api/User/EmailCaptcha',
  avatar = '/Api/User/UserAvatarUpdate',
  authorityList = '/Api/Authority/AuthorityList',
  userBuyVip = '/Api/User/UserBuyVip',
  userBuyShop = '/Api/User/UserBuyShop',
}
interface LoginRes {
  token: string
}

export interface LoginData {
  emailCode: string
  password: string
  captchaCode: string
  captchaId: string
}

export interface RegData {
  username?: string
  password?: string
  rePassword?: string
  emailCode?: string
  captchaCode?: string
  captchaType?: string
}

const getUserInfo = async params => get<ResAuthor>({ url: URL.userInfo, params })
const getCaptcha = async () => get<ResCaptcha>({ url: URL.captcha })
const getUserProfile = async () => get<ResUserProfile>({ url: URL.profile })
const login = async (data: LoginData) =>
  post<IResponse<ResLogin<ResAuthor>>>({ url: URL.login, data }, true)
const reg = async (data: RegData) => post<IResponse>({ url: URL.reg, data }, true)
const forget = async (data: RegData) => post<IResponse>({ url: URL.forget, data }, true)
const checkIn = async () => post<IResponse<ResCheckIn>>({ url: URL.checkIn }, true)
const checkInStatus = async () => post<IResponse>({ url: URL.checkInStatus }, true)
const logout = async () => post<LoginRes>({ url: URL.logout })

const updateUserDetail = async (data: ReqUserDetail) => post<any>({ url: URL.detail, data }, true)
const updateUserPwd = async (data: ReqUserPwd) => post<IResponse>({ url: URL.pwd, data }, true)
const updateUserEmail = async (data: ReqUserEmail) =>
  post<IResponse>({ url: URL.email, data }, true)
const updateUserAvatar = async (data: ReqUserAvatar) =>
  post<IResponse>({ url: URL.avatar, data }, true)
const updateUserEmailCaptcha = async (data: ReqUserEmailCaptcha) =>
  post<IResponse>({ url: URL.emailCaptcha, data }, true)
const getAuthorityList = async (params: ReqAuthorPost) =>
  get<any>({ url: URL.authorityList, params })
const buyVip = async (data: ReqUserBuyVip) => post<any>({ url: URL.userBuyVip, data }, true)
const buyShop = async (data: ReqUserBuyShop) => post<any>({ url: URL.userBuyShop, data }, true)

export {
  buyVip,
  buyShop,
  getAuthorityList,
  getUserInfo,
  getCaptcha,
  getUserProfile,
  logout,
  login,
  reg,
  forget,
  checkIn,
  checkInStatus,
  updateUserDetail,
  updateUserPwd,
  updateUserEmail,
  updateUserEmailCaptcha,
  updateUserAvatar,
}
