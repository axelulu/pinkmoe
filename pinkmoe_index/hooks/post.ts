/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:59:35
 * @FilePath: /pinkmoe_index/hooks/post.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { getPostItem, postView } from '/@/api/post'
import type { ReqPostItem, ResComment, ResPostItem } from '/@/api/post/types'
import { getAuthorPostList } from '/@/api/author'
import type { ResPost } from '/@/api/home/types'
import type { ReqPage, ResPage } from '/@/api/common/types'
import { getCommentList } from '/@/api/comment'
import type { ReqComment } from '/@/api/comment/types'
import { useUserStore } from '../store/modules/user'
import { useAppStore } from '../store/modules/app'

export const usePostItem = async () => {
  const route = useRoute()
  const router = useRouter()
  const comment = ref()
  const postItem = ref<ResPostItem>()
  const recommendPost = ref<ResPage<Array<ResPost>>>()
  const commentList = ref<ResPage<Array<ResPost>>>()
  const user = useUserStore()
  const loading = ref<boolean>(false)
  const loadingPost = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const auth = useUserStore()
  const { siteBasic } = useAppStore()
  const { proxy } = getCurrentInstance()
  const formParams: ReqPostItem = reactive({
    postId: route.params.id,
  })

  const recoFormParams: ReqPage = reactive({
    page: 1,
    pageSize: 6,
    orderKey: 'updated_at',
  })

  const commentFormParams: ReqComment = reactive({
    postId: route.params.id,
    page: 1,
    pageSize: 12,
  })

  // 获取文章信息
  const getPost = async () => {
    loadingPost.value = true
    postItem.value = await getPostItem(formParams)
    setTimeout(() => {
      loadingPost.value = false
    }, 300)
  }

  // 获取推荐文章列表
  const getRecommendPost = async () => {
    recommendPost.value = await getAuthorPostList(recoFormParams)
  }

  const showComment = (commentMeta) => {
    commentFormParams.page = 1
    if (auth.isLogin) {
      comment.value.show(commentMeta)
    }
    else {
      proxy.$login({
        type: 'login',
        router,
        route,
      })
    }
  }

  const getComment = async () => {
    loading.value = true
    commentList.value = await getCommentList(commentFormParams)
    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const refreshComment = async () => {
    hasMore.value = true
    getComment()
  }

  const nextPage = async () => {
    loading.value = true
    commentFormParams.page! += 1
    const res = await getCommentList(commentFormParams)
    if (!res.list || res.list?.length <= 0)
      hasMore.value = false

    else
      commentList.value?.list?.push(...(res?.list as Array<ResComment>))

    setTimeout(() => {
      loading.value = false
    }, 300)
  }

  const share = (type) => {
    switch (type) {
      case 'qq':
        window.open(
          `https://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=/${postItem.value?.post?.postId}&title=${postItem.value?.post?.title}&pics=${postItem.value?.post?.cover}&desc=&summary=&site=`,
          '_blank',
        )
        break
      case 'weibo':
        window.open(
          `https://service.weibo.com/share/share.php?url=/${postItem.value?.post?.postId}&title=${postItem.value?.post?.title}&pic=${postItem.value?.post?.cover}&searchPic=true`,
          '_blank',
        )
        break
      case 'weixin':
        window.open(
          `https://service.weibo.com/share/share.php?url=/${postItem.value?.post?.postId}&title=${postItem.value?.post?.title}&pic=${postItem.value?.post?.cover}&searchPic=true`,
          '_blank',
        )
        break
      case 'bold':
        window.open(
          `http://tieba.baidu.com/f/commit/share/openShareApi?url=/${postItem.value?.post?.postId}&title=${postItem.value?.post?.title}&desc=&comment=&pic=${postItem.value?.post?.cover}&red_tag=j2697365826`,
          '_blank',
        )
        break
      default:
        window.open(
          `https://service.weibo.com/share/share.php?url=/${postItem.value?.post?.postId}&title=${postItem.value?.post?.title}&pic=${postItem.value?.post?.cover}&searchPic=true`,
          '_blank',
        )
        break
    }
  }

  onMounted(() => {
    postView({ postId: route.params.id })
    getComment()
    getRecommendPost()
  })
  await getPost()

  return {
    postItem,
    recommendPost,
    user,
    route,
    siteBasic,
    comment,
    commentList,
    hasMore,
    loading,
    loadingPost,
    formParams,
    getPost,
    share,
    nextPage,
    showComment,
    refreshComment,
  }
}
