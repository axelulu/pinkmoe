/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:03:01
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 16:31:24
 * @FilePath: /pinkmoe_index/hooks/postDownload.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useUserStore } from '/@/store/modules/user'
import { buyPostDownload, getPostDownload } from '/@/api/post'
import { useClipboard } from '@vueuse/core'
import { createDialogModal } from '../utils/createModal'
import MsgConfirm from '/@/components/Msgconfirm/index.vue'

export const usePostDownload = (props) => {
  const auth = useUserStore()
  const postDownload = ref()
  const route = useRoute()
  const router = useRouter()
  const { proxy } = getCurrentInstance()
  const getDownload = async () => {
    postDownload.value = await getPostDownload({ postId: props.postId })
  }

  const showLogin = () => {
    proxy.$login({
      type: 'login',
      router,
      route,
    })
  }

  const copyText = async (text: string) => {
    const { copy } = useClipboard({ source: text })
    await copy()
    proxy.$message({
      type: 'success',
      successMsg: '复制成功',
      loadFun: async () => {
        const code = 200
        return { code }
      },
    })
  }

  const buyDownload = async (item) => {
    const res = await createDialogModal(MsgConfirm, {
      msg: '确认购买此商品？',
    })
    if (res) {
      proxy.$message({
        successMsg: '购买成功',
        failedMsg: '购买失败',
        loadFun: async () => {
          const { code, message } = await buyPostDownload({
            postId: props.postId,
            downloadId: item.ID,
          })
          return { code, message }
        },
      }).then((res) => {
        if (res.code === 200)
          getDownload()
      })
    }
  }

  onMounted(() => {
    getDownload()
  })

  return {
    auth,
    postDownload,
    showLogin,
    copyText,
    buyDownload,
  }
}
