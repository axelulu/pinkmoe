/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-20 21:16:47
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:08:49
 * @FilePath: /pinkmoe_index/src/api/home/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get } from '/@/utils/http/axios'
import type { ResHomeList } from '/@/api/home/types'
// import axios from 'axios';
enum URL {
  list = '/Cms/Home/HomeList',
}

const getHomeList = async () =>
  get<ResHomeList>({
    url: URL.list,
    params: {
      page: 1,
      pageSize: 999,
    },
  })
export { getHomeList }
