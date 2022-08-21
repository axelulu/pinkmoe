/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 11:00:23
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:47:29
 * @FilePath: /pinkmoe_vitesse/src/hooks/vipConfirm.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { buyVip } from '../api/user'
import type { ReqUserBuyVip } from '../api/user/types'
import { createDialogModal } from '../utils/createModal'
import MsgConfirm from '/@/components/Msgconfirm/index.vue'

export const useVipConfirm = (props) => {
  const dialog = ref()
  const num = ref(1)
  const currentVip = ref()
  const showAnimate = ref<boolean>(false)
  const { proxy } = getCurrentInstance()
  const formParams: ReqUserBuyVip = reactive({
    authorityId: '',
    authorityParamId: '',
    authorityParamKey: '',
    authorityParamValue: 0,
    authorityParamDay: 0,
    authorityNum: 0,
  })

  function changeVip(vip) {
    currentVip.value = vip
  }

  async function submitVip() {
    const res = await createDialogModal(MsgConfirm, {
      msg: '确认购买此商品？',
    })
    if (res) {
      formParams.authorityId = currentVip.value.authorityId
      formParams.authorityParamId = currentVip.value.authorityParamId
      formParams.authorityParamKey = currentVip.value.authorityParamKey
      formParams.authorityParamValue = currentVip.value.authorityParamValue
      formParams.authorityParamDay = currentVip.value.authorityParamDay
      formParams.authorityNum = num.value
      const { code, message } = await buyVip(formParams)
      if (code === 200) {
        proxy.$message({
          type: 'success',
          msg: '购买成功',
        })
        dialog.value.hide()
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '购买失败',
        })
      }
    }
  }

  onMounted(() => {
    dialog.value.show()
  })

  return {
    showAnimate,
    dialog,
    num,
    changeVip,
    submitVip,
    currentVip,
  }
}
