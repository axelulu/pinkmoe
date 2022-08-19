/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-05-06 19:19:28
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-18 21:11:43
 * @FilePath: /pinkmoe_admin/src/api/system/user.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { http } from '@/utils/http/axios';
import { BasicResponseModel } from '@/api/common';

/**
 * @description: 获取用户信息
 */
export function getUserInfo() {
  return http.request({
    url: '/Admin/User/User',
    method: 'get',
  },
  {
    isTransformResponse: false,
  });
}

/**
 * @description: 获取用户验证码
 */
export function getCaptcha() {
  return http.request({
    url: '/Admin/User/Captcha',
    method: 'get',
  });
}

/**
 * @description: 获取用户邮箱验证码
 * @param params
 */
export function getEmailCaptcha(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/User/EmailCaptcha',
      method: 'post',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户登录
 * @param params
 */
export function login(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/User/Login',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户注册
 * @param params
 */
export function reg(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/User/Reg',
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户修改密码
 * @param params
 */
export function changePassword(params, uid) {
  return http.request(
    {
      url: `/Admin/user/u${uid}/changepw`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 用户登出
 * @param params
 */
export function logout() {
  return http.request({
    url: '/Admin/User/Logout',
    method: 'POST',
  });
}

/**
 * @description: 获取用户列表
 * @param params
 */
export function getUserList(params) {
  return http.request<BasicResponseModel>(
    {
      url: '/Admin/User/UserList',
      method: 'GET',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 创建用户
 * @param params
 */
export function createUser(params) {
  return http.request(
    {
      url: `/Admin/User/UserCreate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 更新用户
 * @param params
 */
export function updateUser(params) {
  return http.request(
    {
      url: `/Admin/User/UserUpdate`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}

/**
 * @description: 删除用户
 * @param params
 */
export function deleteUser(params) {
  return http.request(
    {
      url: `/Admin/User/UserDelete`,
      method: 'POST',
      params,
    },
    {
      isTransformResponse: false,
    }
  );
}
