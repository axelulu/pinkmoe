/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-21 22:34:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:21:21
 * @FilePath: /pinkmoe_admin/src/utils/http/axios/checkStatus.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
export function checkStatus(status: number, msg: string): void {
  const $message = window['$message'];
  switch (status) {
    case 400:
      $message.error(msg);
      break;
    // 401: 未登录
    // 未登录则跳转登录页面，并携带当前页面的路径
    // 在登录成功后返回当前页面，这一步需要在登录页操作。
    case 401:
      $message.error('用户没有权限（令牌、用户名、密码错误）!');
      break;
    case 403:
      $message.error('用户得到授权，但是访问是被禁止的。!');
      break;
    // 404请求不存在
    case 404:
      $message.error('网络请求错误，未找到该资源!');
      break;
    case 405:
      $message.error('网络请求错误，请求方法未允许!');
      break;
    case 408:
      $message.error('网络请求超时');
      break;
    case 500:
      $message.error('服务器错误,请联系管理员!');
      break;
    case 501:
      $message.error('网络未实现');
      break;
    case 502:
      $message.error('网络错误');
      break;
    case 503:
      $message.error('服务不可用，服务器暂时过载或维护!');
      break;
    case 504:
      $message.error('网络超时');
      break;
    case 505:
      $message.error('http版本不支持该请求!');
      break;
    default:
      $message.error(msg);
  }
}
