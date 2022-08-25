/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-25 20:21:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-25 20:24:28
 * @FilePath: /pinkmoe_index/ecosystem.config.js
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
module.exports = {
  apps: [
    {
      name: 'pinkmoe_index',
      exec_mode: 'cluster',
      instances: 'max',
      script: './.output/server/index.mjs',
      watch: true,
      env: {
        PORT: 3000,
        NODE_ENV: 'development',
      },
      env_production: {
        PORT: 3000,
        NODE_ENV: 'production',
      },
    },
  ],
}
