/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-27 22:05:52
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:58:33
 * @FilePath: /pinkmoe_index/hooks/postVideo.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useUserStore } from '/@/store/modules/user';
import { buyPostVideo, getPostVideo } from '/@/api/post'
import MsgConfirm from '/@/components/Msgconfirm/index.vue'
import { createDialogModal } from '../utils/createModal'
import artplayerPluginDanmuku from 'artplayer-plugin-danmuku'

export const usePostVideo = async (props) => {
  const auth = useUserStore()
  const postVideo = ref()
  const currentVideo = ref()
  const art = ref()
  const route = useRoute()
  const router = useRouter()
  const style = ref({
    width: '100%',
    height: '100%',
    margin: '',
    zIndex: 1,
  })
  const { proxy } = getCurrentInstance()

  const showLogin = () => {
    proxy.$login({
      type: 'login',
      router,
      route,
    })
  }

  const options = reactive({
    url: '',
    autoSize: true,
    fullscreen: true,
    fullscreenWeb: true,
    flip: true,
    playbackRate: true,
    aspectRatio: true,
    setting: true,
    settings: [
      {
        html: '画中画模式',
        tooltip: '关闭',
        icon: '<img width="22" heigth="22" src="/assets/img/state.svg">',
        switch: false,
        onSwitch(item, $dom, event) {
          const nextState = !item.switch
          art.value.pip = nextState

          // 改变提示文本
          item.tooltip = nextState ? '开启' : '关闭'

          // 改变按钮状态
          return nextState
        },
      },
      {
        html: '选择字幕',
        width: 250,
        tooltip: '字幕 01',
        selector: [
          {
            default: true,
            html: '<span style="color:red">字幕 01</span>',
            url: '/assets/sample/subtitle.srt?id=1',
          },
          {
            html: '<span style="color:yellow">字幕 02</span>',
            url: '/assets/sample/subtitle.srt?id=2',
          },
        ],
        onSelect(item, $dom, event) {
          art.subtitle.url = item.url

          // 改变提示文本
          return item.html
        },
      },
      {
        html: '选择画质',
        width: 150,
        tooltip: '1080P',
        selector: [
          {
            default: true,
            html: '1080P',
            url: '/assets/sample/video.mp4?id=1080',
          },
          {
            html: '720P',
            url: '/assets/sample/video.mp4?id=720',
          },
          {
            html: '360P',
            url: '/assets/sample/video.mp4?id=360',
          },
        ],
        onSelect(item, $dom, event) {
          art.switchQuality(item.url, item.html)

          // 改变提示文本
          return item.html
        },
      },
    ],
    plugins: [
      artplayerPluginDanmuku({
        // 弹幕数组
        danmuku: [
          {
            text: '111', // 弹幕文本
            time: 1, // 发送时间，单位秒
            color: '#fff', // 弹幕局部颜色
            border: false, // 是否显示描边
            mode: 0, // 弹幕模式: 0表示滚动、1静止
          },
          {
            text: '222',
            time: 2,
            color: 'red',
            border: true,
            mode: 0,
          },
          {
            text: '333',
            time: 3,
            color: 'green',
            border: false,
            mode: 1,
          },
        ],
        speed: 5, // 弹幕持续时间，单位秒，范围在[1 ~ 10]
        opacity: 1, // 弹幕透明度，范围在[0 ~ 1]
        fontSize: 25, // 字体大小，支持数字和百分比
        color: '#FFFFFF', // 默认字体颜色
        mode: 0, // 默认模式，0-滚动，1-静止
        margin: [10, '25%'], // 弹幕上下边距，支持数字和百分比
        antiOverlap: true, // 是否防重叠
        useWorker: true, // 是否使用 web worker
        synchronousPlayback: false, // 是否同步到播放速度
        filter: danmu => danmu.text.length < 50, // 弹幕过滤函数，返回 true 则可以发送
        lockTime: 5, // 输入框锁定时间，单位秒，范围在[1 ~ 60]
        maxLength: 100, // 输入框最大可输入的字数，范围在[0 ~ 500]
        minWidth: 200, // 输入框最小宽度，范围在[0 ~ 500]，填 0 则为无限制
        maxWidth: 400, // 输入框最大宽度，范围在[0 ~ Infinity]，填 0 则为 100% 宽度
        theme: 'dark', // 输入框自定义挂载时的主题色，默认为 dark，可以选填亮色 light
        beforeEmit: danmu => !!danmu.text.trim(), // 发送弹幕前的自定义校验，返回 true 则可以发送

        // 通过 mount 选项可以自定义输入框挂载的位置，默认挂载于播放器底部，仅在当宽度小于最小值时生效
        // mount: document.querySelector('.artplayer-danmuku'),
      }),
    ],
  })

  const getVideo = async () => {
    postVideo.value = await getPostVideo({ postId: props.postId })
    currentVideo.value = postVideo.value?.[0]
    options.url = postVideo.value?.[0]?.url
  }

  const changeVideo = async (video) => {
    if (video.url === '') {
      proxy.$message({
        type: 'warning',
        msg: '商品未购买！',
      })
      return
    }
    if (!auth.isLogin) {
      showLogin()
      return
    }
    currentVideo.value = video
    options.url = video.url
    art.value.seek = 0
    art.value.toggle()
  }

  const buyVideo = async (item) => {
    const res = await createDialogModal(MsgConfirm, {
      msg: '确认购买此商品？',
    })
    if (res) {
      if (!auth.isLogin) {
        proxy.$login({
          type: 'login',
          router,
          route,
        })
        return
      }
      const { code, message } = await buyPostVideo({
        postId: props.postId,
        videoId: item.ID,
      })
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '购买成功',
        })
        getVideo()
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '购买失败',
        })
      }
    }
  }

  const getInstance = (artp) => {
    art.value = artp
  }

  await getVideo()

  return {
    postVideo,
    buyVideo,
    auth,
    changeVideo,
    currentVideo,
    options,
    showLogin,
    getInstance,
    style,
  }
}
