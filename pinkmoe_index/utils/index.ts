/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-18 21:44:07
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 15:46:07
 * @FilePath: /pinkmoe_vitesse/src/utils/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { resolve } from 'path'
import fs from 'fs'

function pathResolve(dir: string) {
  return resolve(process.cwd(), '.', dir)
}

export const getFolder = (path: any) => {
  const components: Array<string> = []
  const files = fs.readdirSync(path)
  files.forEach((item: string) => {
    const stat = fs.lstatSync(`${path}/${item}`)
    if (stat.isDirectory() === true && item != 'components') {
      components.push(`${path}/${item}`)
      components.push(pathResolve(`${path}/${item}`))
    }
  })
  return components
}

function isWechat(UA) {
  return !!/MicroMessenger/i.test(UA)
}

function isWeibo(UA) {
  return !!/Weibo/i.test(UA)
}

function isQQ(UA) {
  return !!/QQ/i.test(UA)
}

function isMoible(UA) {
  return !!/(Android|webOS|iPhone|iPod|tablet|BlackBerry|Mobile)/i.test(UA)
}

function isIOS(UA) {
  return !!/iPhone|iPad|iPod/i.test(UA)
}

function isAndroid(UA) {
  return !!/Android/i.test(UA)
}

export function deviceType(UA) {
  if (isMoible(UA)) {
    if (isIOS(UA)) {
      if (isWechat(UA)) {
        return {
          type: 'ios',
          env: 'wechat',
          masklayer: true,
        }
      }
      if (isWeibo(UA)) {
        return {
          type: 'ios',
          env: 'weibo',
          masklayer: true,
        }
      }
      if (isQQ(UA)) {
        return {
          type: 'ios',
          env: 'qq',
          masklayer: true,
        }
      }
      return {
        type: 'ios',
      }
    }
    if (isAndroid(UA)) {
      if (isWechat(UA)) {
        return {
          type: 'android',
          env: 'wechat',
          masklayer: true,
        }
      }
      if (isWeibo(UA)) {
        return {
          type: 'android',
          env: 'weibo',
          masklayer: true,
        }
      }
      if (isQQ(UA)) {
        return {
          type: 'android',
          env: 'qq',
          masklayer: true,
        }
      }
      return {
        type: 'android',
      }
    }
    return {
      type: 'mobile',
    }
  }
  else {
    return {
      type: 'pc',
    }
  }
}
