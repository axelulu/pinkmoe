/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 13:22:29
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-10 19:37:07
 * @FilePath: /pinkmoe_index/src/api/post/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
// 权限问题后期增加
import { get, post } from '/@/utils/http/axios'
import type { ReqPostItem, ResBbsSilder, ResPostItem } from '/@/api/post/types'
import type { ReqPage, ResPage } from '/@/api/common/types'
import type { IResponse } from '/@/utils/http/axios/type'
import type { ResPost } from '../home/types'
// import axios from 'axios';
enum URL {
  item = '/Cms/Post/PostItem',
  list = '/Cms/Post/PostList',
  bbsSilder = '/Cms/Bbs/BbsSilder',
  userList = '/Cms/Post/UserPostList',
  create = '/Cms/Post/PostCreate',
  view = '/Cms/Post/PostViewUpdate',
  collect = '/Cms/Post/PostCollect',
  collectList = '/Cms/Post/CollectPost',
  unCollect = '/Cms/Post/PostUnCollect',
  download = '/Cms/Post/PostItemDownload',
  downloadBuy = '/Cms/Post/PostItemDownloadBuy',
  music = '/Cms/Post/PostItemMusic',
  musicBuy = '/Cms/Post/PostItemMusicBuy',
  video = '/Cms/Post/PostItemVideo',
  videoBuy = '/Cms/Post/PostItemVideoBuy',
}

const getPostItem = async (params: ReqPostItem) =>
  get<ResPostItem>({
    url: URL.item,
    params,
  })

const getPostList = async (params: ReqPage) =>
  get<ResPage<Array<ResPost>>>({
    url: URL.list,
    params,
  })

const getBbsSilder = async () =>
  get<ResBbsSilder>({
    url: URL.bbsSilder,
  })

const getCollectPostList = async (params: ReqPage) =>
  get<ResPage<Array<ResPost>>>({
    url: URL.collectList,
    params,
  })

const getUserPost = async (params: ReqPage) =>
  get<ResPage<Array<ResPost>>>({
    url: URL.userList,
    params,
  })

const getPostDownload = async params =>
  get<any>({
    url: URL.download,
    params,
  })

const getPostMusic = async params =>
  get<any>({
    url: URL.music,
    params,
  })

const getPostVideo = async params =>
  get<any>({
    url: URL.video,
    params,
  })

const postView = async data =>
  post<any>({
    url: URL.view,
    data,
  })

const collectPost = async data =>
  post<any>(
    {
      url: URL.collect,
      data,
    },
    true,
  )

const unCollectPost = async data =>
  post<any>(
    {
      url: URL.unCollect,
      data,
    },
    true,
  )

const buyPostDownload = async data =>
  post<any>(
    {
      url: URL.downloadBuy,
      data,
    },
    true,
  )

const buyPostMusic = async data =>
  post<any>(
    {
      url: URL.musicBuy,
      data,
    },
    true,
  )

const buyPostVideo = async data =>
  post<any>(
    {
      url: URL.videoBuy,
      data,
    },
    true,
  )

const createPost = async data =>
  post<IResponse>(
    {
      url: URL.create,
      data,
    },
    true,
  )

export {
  getPostItem,
  postView,
  getCollectPostList,
  collectPost,
  unCollectPost,
  getPostList,
  getUserPost,
  createPost,
  getPostDownload,
  buyPostDownload,
  getPostMusic,
  buyPostMusic,
  getPostVideo,
  buyPostVideo,
  getBbsSilder,
}
