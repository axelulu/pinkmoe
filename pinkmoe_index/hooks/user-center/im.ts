/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 18:21:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:20:00
 * @FilePath: /pinkmoe_index/hooks/user-center/im.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import {
  createChatRelationList,
  deleteChatRelationList,
  getChatList,
  getChatRelationList,
} from '/@/api/chat'
import type { ReqChat, ResChat, ResChatRelation } from '/@/api/chat/types'
import type { ResPage } from '/@/api/common/types'
import { useUserStore } from '/@/store/modules/user'
import { useSocketStore } from '/@/store/modules/socket'

export const useUserCenterIm = () => {
  const { socket } = useSocketStore()
  const socketWs = ref<any>()
  const msgRef = ref()
  const msg = ref('')
  const route = useRoute()
  const currentSendId = ref()
  const relationChat = ref<any>()
  const chatList = ref<ResPage<Array<ResChat>>>()
  const chatRelationList = ref<ResPage<Array<ResChatRelation>>>()
  const { proxy } = getCurrentInstance()
  const auth = useUserStore()
  const formParams: ReqChat = reactive({
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
    sendId: '',
  })

  const relationFormParams: ReqChat = reactive({
    page: 1,
    pageSize: 12,
    orderKey: 'updated_at',
    desc: true,
    sendId: '',
  })

  async function getChatRelation() {
    chatRelationList.value = await getChatRelationList(relationFormParams)
  }

  async function getChat() {
    chatList.value = await getChatList(formParams)
  }

  async function sendMsg() {
    if (socketWs.value) {
      socketWs.value.send(
        JSON.stringify({
          type: 'chat',
          chatMsg: { content: msg.value, sendId: formParams.sendId, userId: auth.userInfo?.uuid },
        }),
      )
      msg.value = ''
    }
    else {
      proxy.$message({
        type: 'warning',
        successMsg: '请选择用户',
        loadFun: async () => {
          const code = 200
          return { code }
        },
      })
    }
  }

  async function addChatRelation() {
    proxy.$message({
      successMsg: '添加成功',
      failedMsg: '添加失败',
      loadFun: async () => {
        const { code, message } = await createChatRelationList({
          sendId: relationChat.value,
        })
        return { code, message }
      },
    }).then(async (res) => {
      if (res === 200) {
        await getChatRelation()
        formParams.sendId = relationChat.value
        getChat()
      }
    })
  }

  async function addChat(sendId) {
    currentSendId.value = sendId
    formParams.sendId = sendId
    formParams.page = 1
    await getChat()
    scrollToBottom()
  }

  function scrollToBottom() {
    proxy.$nextTick(() => {
      msgRef.value.scrollTo({
        top: msgRef.value.scrollHeight,
        // 滚动过渡效果
        behavior: 'smooth',
      })
    })
  }

  const scrollHandler = async () => {
    // 当前滚动高度
    const curScrollTop = msgRef.value.scrollTop
    const curScrollBottom = msgRef.value.scrollHeight - msgRef.value.clientHeight - curScrollTop
    // if (curScrollTop === 0) {
    //   (formParams.page as number)--
    //   chatList.value = await getChatList(formParams);
    // }
    if (curScrollBottom === 0) {
      (formParams.page as number)++
      const res = await getChatList(formParams)
      chatList.value?.list?.push(...res.list)
      scrollToBottom()
    }
  }

  watchEffect(() => {
    nextTick(() => {
      if (typeof msgRef.value !== 'undefined') {
        msgRef.value.removeEventListener('scroll', scrollHandler)
        msgRef.value.addEventListener('scroll', scrollHandler)
      }
    })
  })

  function getSocketData(res) {
    if (socketWs.value) {
      if (res.detail.data.type === 'chat')
        chatList.value?.list?.push(res.detail.data.chatMsg)

      scrollToBottom()
    }
  }

  function getSocketOpenData(res) {
    socketWs.value = res.detail.data
  }

  async function deleteChatRelation(sendId) {
    await deleteChatRelationList({
      sendId,
    })
    getChatRelation()
    chatList.value = undefined
  }

  onMounted(async () => {
    await getChatRelation()
    socketWs.value = socket
    if (route.query?.sendId) {
      await createChatRelationList({
        sendId: route.query.sendId,
      })
      getChatRelation()
      addChat(route.query.sendId)
    }
    else {
      await createChatRelationList({
        sendId: chatRelationList.value?.list?.[0]?.sendIdRelation?.uuid,
      })
      addChat(chatRelationList.value?.list?.[0]?.sendIdRelation?.uuid)
    }
    window.addEventListener('onmessageWS', getSocketData)
    window.addEventListener('onopenWS', getSocketOpenData)
  })

  onUnmounted(() => {
    window.removeEventListener('onmessageWS', getSocketData)
    window.removeEventListener('onopenWS', getSocketOpenData)
  })

  return {
    msgRef,
    sendMsg,
    addChatRelation,
    deleteChatRelation,
    auth,
    addChat,
    relationChat,
    msg,
    chatList,
    chatRelationList,
    currentSendId,
  }
}
