/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 15:34:15
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 15:35:22
 * @FilePath: /pinkmoe_vitesse/config/vite/proxy.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import type { ProxyOptions } from 'vite'
import {
  API_BASE_URL,
  API_SOCKET_TARGET_URL,
  API_SOCKET_URL,
  API_TARGET_URL,
  API_UPLOAD_TARGET_URL,
  API_UPLOAD_URL,
  MOCK_API_BASE_URL,
  MOCK_API_TARGET_URL,
} from '../../config/constant'
type ProxyTargetList = Record<string, ProxyOptions>

const init: ProxyTargetList = {
  // test
  [API_BASE_URL]: {
    target: API_TARGET_URL,
    changeOrigin: true,
    rewrite: path => path.replace(new RegExp(`^${API_BASE_URL}`), ''),
  },
  [API_SOCKET_URL]: {
    target: API_SOCKET_TARGET_URL,
    ws: true,
    changeOrigin: true,
    rewrite: path => path.replace(new RegExp(`^${API_SOCKET_URL}`), ''),
  },
  [API_UPLOAD_URL]: {
    target: API_UPLOAD_TARGET_URL,
    changeOrigin: true,
    rewrite: path => path.replace(new RegExp(`^${API_UPLOAD_URL}`), ''),
  },
  // mock
  [MOCK_API_BASE_URL]: {
    target: MOCK_API_TARGET_URL,
    changeOrigin: true,
    rewrite: path => path.replace(new RegExp(`^${MOCK_API_BASE_URL}`), '/api'),
  },
}

export default init
