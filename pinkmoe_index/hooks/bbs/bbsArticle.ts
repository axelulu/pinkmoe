/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-24 09:00:58
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:10:05
 * @FilePath: /pinkmoe_index/hooks/bbs/bbsArticle.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { useUtil } from '../util'
import { createFollow, deleteFollow, followStatus } from '/@/api/follow'
import type { ReqFollowStatus } from '/@/api/follow/types'
import { collectPost, unCollectPost } from '/@/api/post'
import { useUserStore } from '/@/store/modules/user'

export const useBbsArticle = (post) => {
  const status = ref<boolean>(false)
  const auth = useUserStore()
  const router = useRouter()
  const route = useRoute()
  const isCollect = ref(false)
  const formParams: ReqFollowStatus = reactive({
    toUid: post?.AuthorRelation?.uuid,
  })
  const { proxy } = getCurrentInstance()

  const deletePostImg = (content) => {
    const imgReg = /<img.*?(?:>|\/>)/gi // 定义正则表达式（拿到img标签所有值）
    const deArray = content.match(imgReg) // 使用正则表达式，拿到的是数组
    // ["<img src="images/01.gif">", "<img src="images/02.gif">"]
    if (deArray != null && deArray != undefined) {
      for (const vm of deArray)
        content = content.replace(vm, '*') // 放在循环中剔除img标签
    }
  }
  const { getLevel } = useUtil()
  const lev = getLevel(post?.AuthorRelation?.exp)

  const getFollowStatus = async () => {
    if (auth.isLogin)
      status.value = await followStatus(formParams)
  }

  const follow = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
    else {
      proxy.$message({
        successMsg: '取消关注成功',
        failedMsg: '取消关注失败',
        loadFun: async () => {
          const { code, message } = await deleteFollow(formParams)
          return { code, message }
        },
      }).then(async (res) => {
        if (res === 200)
          await getFollowStatus()
      })
    }
  }

  const unFollow = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
    else {
      proxy.$message({
        successMsg: '关注成功',
        failedMsg: '关注失败',
        loadFun: async () => {
          const { code, message } = await createFollow(formParams)
          return { code, message }
        },
      }).then(async (res) => {
        if (res === 200)
          await getFollowStatus()
      })
    }
  }

  const collect = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
    else {
      proxy.$message({
        successMsg: '收藏成功',
        failedMsg: '收藏失败',
        loadFun: async () => {
          const { code, message } = await collectPost({ postId: post.postId })
          return { code, message }
        },
      }).then(async (res) => {
        if (res === 200) {
          isCollect.value = true
          await getFollowStatus()
        }
      })
    }
  }

  const unCollect = async () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
    else {
      proxy.$message({
        successMsg: '取消收藏成功',
        failedMsg: '取消收藏失败',
        loadFun: async () => {
          const { code, message } = await unCollectPost({ postId: post.postId })
          return { code, message }
        },
      }).then(async (res) => {
        if (res === 200) {
          isCollect.value = false
          await getFollowStatus()
        }
      })
    }
  }

  const jump = () => {
    if (!auth.isLogin) {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
    else {
      router.push({
        path: '/user-center/im',
        query: { sendId: post?.AuthorRelation?.uuid },
      })
    }
  }

  onMounted(() => {
    isCollect.value = post?.collectRelation.length
    getFollowStatus()
  })

  return {
    status,
    auth,
    isCollect,
    jump,
    unFollow,
    follow,
    collect,
    unCollect,
    deletePostImg,
    lev,
  }
}
