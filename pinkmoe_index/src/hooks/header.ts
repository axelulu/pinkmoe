/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-22 12:19:32
 * @FilePath: /pinkmoe_index/src/hooks/header.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
// @unocss-include
import { getCategoryList } from '/@/api/category'
import type { ResCategory } from '/@/api/category/types'
import { useAppStore, useUserStore } from '/@/store'
import { useDark, useToggle } from '@vueuse/core'
import { checkIn } from '/@/api/user'

export const useHeader = () => {
  const categoryList = ref<Array<ResCategory>>()
  const { proxy } = getCurrentInstance()
  const router = useRouter()
  const route = useRoute()
  const appStore = useAppStore()
  const auth = useUserStore()
  const children = ref()
  const i18n = useI18n()
  const categoryIndex = ref()
  const isShow = ref(false)
  const headerReCategory = ref()
  const theme = computed(() => {
    return appStore.theme
  })
  const setting = ref<any>([
    {
      icon: 'i-ph:paint-brush-broad-fill',
      name: '文章管理',
      children: [
        {
          icon: 'i-ph:paint-brush-broad-fill',
          name: '资源投稿',
          url: '/user-center/publish',
          func: null,
        },
        {
          icon: 'i-fluent:copy-24-filled',
          name: '我的帖子',
          url: '/user-center/posts',
          func: null,
        },
        {
          icon: 'i-mdi:cards-heart',
          name: '收藏夹',
          url: '/user-center/stars',
          func: null,
        },
        {
          icon: 'i-fa-solid:comments',
          name: '我的吐槽',
          url: '/user-center/comments',
          func: null,
        },
      ],
    },
    {
      icon: 'i-ph:bell-fill',
      name: '消息管理',
      children: [
        {
          icon: 'i-ri:chat-history-fill',
          name: '消费记录',
          url: '/user-center/record',
          func: null,
        },
        {
          icon: 'i-ph:bell-fill',
          name: '我的提醒',
          url: '/user-center/notice',
          func: null,
        },
        {
          icon: 'i-jam:envelope-f',
          name: '即时通讯',
          url: '/user-center/im',
          func: null,
        },
      ],
    },
    {
      icon: 'i-fa6-solid:bag-shopping',
      name: '商城论坛',
      children: [
        {
          icon: 'i-fa-solid:address-card',
          name: '收货地址',
          url: '/user-center/home',
          func: null,
        },
        {
          icon: 'i-fa6-solid:basket-shopping',
          name: '订单记录',
          url: '/user-center/follow',
          func: null,
        },
        {
          icon: 'i-fa6-solid:people-group',
          name: '关注圈子',
          url: '/user-center/fans',
          func: null,
        },
      ],
    },
    {
      icon: 'i-mdi:gamepad-square',
      name: '游戏中心',
      children: [
        {
          icon: 'i-material-symbols:bookmark',
          name: '我的荣誉',
          url: '/user-center/medal',
          func: null,
        },
        {
          icon: 'i-mdi:gift',
          name: '抽奖',
          url: '/user-center/lottery',
          func: null,
        },
        {
          icon: 'i-material-symbols:shopping-bag',
          name: '商城',
          url: '/user-center/shop',
          func: null,
        },
      ],
    },
    {
      icon: 'i-material-symbols:settings-applications-rounded',
      name: '个人设置',
      children: [
        {
          icon: 'i-icon-park-solid:vip-one',
          name: '我的会员',
          url: '/user-center/vip',
          func: null,
        },
        {
          icon: 'i-mdi:cog',
          name: '我的设置',
          url: '/user-center/settings',
          func: null,
        },
        {
          icon: 'i-fa6-solid:power-off',
          name: '退出登陆',
          url: '/user-center/publish',
          func: () => {
            auth.logout()
            proxy.$message({
              type: 'success',
              msg: '退出登陆成功',
            })
            setTimeout(() => {
              router.push(route.path)
            }, 1000)
          },
        },
      ],
    },
  ])
  const headerBarLeft = ref([
    {
      name: '社区',
      icon: 'i-fa6-solid:people-group',
      url: 'bbs',
    },
    {
      name: '商城',
      icon: 'i-material-symbols:shopping-bag',
      url: 'shop',
    },
  ])
  const headerBarRight = ref([
    {
      name: '夜间模式',
      icon: 'i-icon-park-solid:theme',
      type: 'theme',
    },
    {
      name: '语言',
      icon: 'i-fa-solid:language',
      type: 'lang',
    },
  ])
  const isDark = useDark({
    selector: 'body',
    attribute: 'arco-theme',
    valueDark: 'dark',
    valueLight: 'light',
    storageKey: 'arco-theme',
    onChanged(dark: boolean) {
      appStore.toggleTheme(dark)
    },
  })
  const toggleTheme = useToggle(isDark)
  // 列表HTMLElementDom
  const data = reactive<any>({
    // 列表第一项的高度（起始高度）
    initH: 192 - 56,
    // 是否固定
    fixed: false,
  })

  // 获取分类
  const getCategory = async () => {
    const categories = await getCategoryList()
    categoryList.value = categories.list
  }

  const showLogin = () => {
    proxy.$login({
      type: 'login',
      router,
      route,
    })
  }

  const showSearchDialog = () => {
    proxy.$search({
      router,
    })
  }

  const scrollHandler = () => {
    // 当前滚动高度
    if (!import.meta.env.SSR) {
      const curScrollTop = document.documentElement.scrollTop
      data.fixed = curScrollTop > data.initH
    }
  }

  const checkInUser = async () => {
    const { code, message } = await checkIn()
    if (code === 200) {
      proxy.$message({
        type: 'success',
        msg: '签到成功',
      })
      await auth.checkIn()
    }
    else {
      proxy.$message({
        type: 'warning',
        msg: message || '签到失败',
      })
    }
  }

  const mouseenter = (item, index) => {
    isShow.value = true
    children.value = item
    categoryIndex.value = index
    if (headerReCategory.value && headerReCategory.value.length > 0)
      headerReCategory.value[index].hide()
  }

  function seleLanguage() {
    const locale = localStorage.getItem('locale')
    if (locale === 'zh-CN') {
      const idx = ['zh-CN', 'ja'][1] || navigator.language.slice(0, 2)
      localStorage.setItem('locale', idx)
      i18n.locale.value = idx
    }
    else {
      const idx = ['zh-CN', 'ja'][0] || navigator.language.slice(0, 2)
      localStorage.setItem('locale', idx)
      i18n.locale.value = idx
    }
  }

  watchEffect(() => {
    nextTick(() => {
      if (typeof window !== 'undefined') {
        window.removeEventListener('scroll', scrollHandler)
        window.addEventListener('scroll', scrollHandler)
      }
    })
  })

  function getSocketData(res) {
    if (
      res.detail.data.type === 'chat'
      && res.detail.data.chatMsg?.userIdRelation?.uuid !== auth.userInfo?.uuid
    ) {
      proxy.$notify({
        content: res.detail.data.chatMsg,
        type: 'chat',
        router,
      })
    }
    // ...业务处理
  }

  onUnmounted(() => {
    if (typeof window !== 'undefined')
      window.removeEventListener('scroll', scrollHandler)
  })

  onMounted(async () => {
    scrollHandler()
    await getCategory()
    await auth.checkIn()
    window.addEventListener('onmessageWS', getSocketData)
  })

  return {
    data,
    headerBarLeft,
    headerBarRight,
    headerReCategory,
    isShow,
    children,
    auth,
    appStore,
    categoryList,
    toggleTheme,
    categoryIndex,
    setting,
    theme,
    showSearchDialog,
    showLogin,
    seleLanguage,
    mouseenter,
    checkInUser,
    i18n,
  }
}
