/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:03:33
 * @FilePath: /pinkmoe_index/src/utils/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { resolve } from 'path';
const fs = require('fs');

function pathResolve(dir: string) {
  return resolve(process.cwd(), '.', dir);
}

export const getFolder = (path: any) => {
  const components: Array<string> = [];
  const files = fs.readdirSync(path);
  files.forEach(function (item: string) {
    const stat = fs.lstatSync(path + '/' + item);
    if (stat.isDirectory() === true && item != 'components') {
      components.push(path + '/' + item);
      components.push(pathResolve(path + '/' + item));
    }
  });
  return components;
};

/**
 *
 * @param {*} UA ,就是userAgent
 * @returns  type: 设备类型
 *           env: 访问环境(微信/微博/qq)
 *           masklayer: 就是给外部拿到判断是否显示遮罩层的,一些特殊环境要引导用户到外部去打开访问
 */
function isWechat(UA) {
  return /MicroMessenger/i.test(UA) ? true : false;
}

function isWeibo(UA) {
  return /Weibo/i.test(UA) ? true : false;
}

function isQQ(UA) {
  return /QQ/i.test(UA) ? true : false;
}

function isMoible(UA) {
  return /(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(UA) ? true : false;
}

function isIOS(UA) {
  return /iPhone|iPad|iPod/i.test(UA) ? true : false;
}

function isAndroid(UA) {
  return /Android/i.test(UA) ? true : false;
}

export function deviceType(UA) {
  if (isMoible(UA)) {
    if (isIOS(UA)) {
      if (isWechat(UA)) {
        return {
          type: 'ios',
          env: 'wechat',
          masklayer: true,
        };
      }
      if (isWeibo(UA)) {
        return {
          type: 'ios',
          env: 'weibo',
          masklayer: true,
        };
      }
      if (isQQ(UA)) {
        return {
          type: 'ios',
          env: 'qq',
          masklayer: true,
        };
      }
      return {
        type: 'ios',
      };
    }
    if (isAndroid(UA)) {
      if (isWechat(UA)) {
        return {
          type: 'android',
          env: 'wechat',
          masklayer: true,
        };
      }
      if (isWeibo(UA)) {
        return {
          type: 'android',
          env: 'weibo',
          masklayer: true,
        };
      }
      if (isQQ(UA)) {
        return {
          type: 'android',
          env: 'qq',
          masklayer: true,
        };
      }
      return {
        type: 'android',
      };
    }
    return {
      type: 'mobile',
    };
  } else {
    return {
      type: 'pc',
    };
  }
}
