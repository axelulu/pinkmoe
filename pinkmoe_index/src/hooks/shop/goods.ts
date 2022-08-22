/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-22 22:50:29
 * @FilePath: /pinkmoe_index/src/hooks/shop/goods.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { postView } from '/@/api/post'
import type { ResComment } from '/@/api/post/types'
import { useAppStore, useUserStore } from '/@/store'
import type { ResPost } from '/@/api/home/types'
import type { ResPage } from '/@/api/common/types'
import { getCommentList } from '/@/api/comment'
import type { ReqComment } from '/@/api/comment/types'
import { getGoodsItem } from '/@/api/goods'
import type { ReqGoodsItem, ResGoodsItem } from '/@/api/goods/types'

export const useGoodsItem = () => {
  const route = useRoute()
  const router = useRouter()
  const comment = ref()
  const cover = ref()
  const num = ref(0)
  const saleAmount = ref(0)
  const stock = ref(0)
  const goodsItem = ref<ResGoodsItem>()
  const commentList = ref<ResPage<Array<ResPost>>>()
  const user = useUserStore()
  const loading = ref<boolean>(false)
  const hasMore = ref<boolean>(true)
  const auth = useUserStore()
  const { siteBasic } = useAppStore()
  const { proxy } = getCurrentInstance()
  const formParams: ReqGoodsItem = reactive({
    goodsId: route.params.id,
  })

  const commentFormParams: ReqComment = reactive({
    postId: route.params.id,
    page: 1,
    pageSize: 12,
  })

  // 获取文章信息
  const getPost = async () => {
    goodsItem.value = await getGoodsItem(formParams)
    goodsItem.value.goods?.sizeRelation?.forEach((element) => {
      const v = JSON.parse(element.value)
      const vv = []
      v.forEach((ele, index) => {
        if (index === 0) {
          vv.push({
            value: ele,
            status: true,
          })
        }
        else {
          vv.push({
            value: ele,
            status: false,
          })
        }
      })
      element.value = vv
    })
    saleAmount.value = goodsItem.value.goods?.sizeValRelation?.[0].saleAmount
    stock.value = goodsItem.value.goods?.sizeValRelation?.[0].stock
    cover.value = goodsItem.value?.goods?.cover
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
    console.log(route)
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

  const changeSizeVal = (index, size) => {
    size.forEach((element) => {
      element.status = false
    })
    size[index].status = true

    // 处理商品规格变化
    const sizeValPool = []
    goodsItem.value?.goods?.sizeRelation?.forEach((e) => {
      e.value.forEach((ee) => {
        if (ee.status === true)
          sizeValPool.push(ee.value)
      })
    })

    goodsItem.value?.goods?.sizeValRelation?.forEach((e) => {
      if (e.value == JSON.stringify(sizeValPool)) {
        stock.value = e.stock
        saleAmount.value = e.saleAmount
      }
    })
  }

  onMounted(() => {
    postView({ postId: route.params.id })
    getPost()
    getComment()
  })

  return {
    goodsItem,
    user,
    route,
    siteBasic,
    comment,
    commentList,
    hasMore,
    loading,
    stock,
    saleAmount,
    cover,
    num,
    share,
    nextPage,
    changeSizeVal,
    getComment,
    showComment,
    refreshComment,
  }
}
