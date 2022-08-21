/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-04 07:41:38
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 17:47:23
 * @FilePath: /pinkmoe_vitesse/src/hooks/shopConfirm.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { buyShop } from '../api/user'
import type { ReqUserBuyShop } from '../api/user/types'
import { createDialogModal } from '../utils/createModal'
import MsgConfirm from '/@/components/Msgconfirm/index.vue'

export const useShopConfirm = (props) => {
  const dialog = ref()
  const num = ref(1)
  const currentVip = ref()
  const key = ref()
  const showAnimate = ref<boolean>(false)
  const { proxy } = getCurrentInstance()
  const formParams: ReqUserBuyShop = reactive({
    shopType: '',
    priceType: '',
    shopCredit: 0,
  })

  async function submitShop() {
    const res = await createDialogModal(MsgConfirm, {
      msg: '确认购买此商品？',
    })
    if (res) {
      formParams.shopType = currentVip.value.shopType
      formParams.priceType = currentVip.value.priceType
      formParams.shopCredit = parseInt(currentVip.value.shopCredit)
      formParams.shopValue = currentVip.value.shopValue
      formParams.shopNum = num.value
      formParams.shopKey = key.value
      const { code, message } = await buyShop(formParams)
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
    currentVip.value = props.authorShop
  })

  return {
    showAnimate,
    dialog,
    num,
    submitShop,
    currentVip,
    key,
  }
}
