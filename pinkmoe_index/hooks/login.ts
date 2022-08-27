/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 14:58:23
 * @FilePath: /pinkmoe_index/hooks/login.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
 */
import { useUserStore } from '/@/store/modules/user';
import type { LoginData, RegData } from '/@/api/user'
import { getCaptcha, updateUserEmailCaptcha } from '/@/api/user'
import type { ReqUserEmailCaptcha } from '../api/user/types'
import { useUtil } from './util'

export const useLogin = (props) => {
  const dialog = ref()
  const formLogin = ref()
  const formReg = ref()
  const formForgetPwd = ref()
  const showAnime = ref<string>()
  const loginType = ref<string>()
  const { proxy } = getCurrentInstance()
  const { checkForm } = useUtil()
  const sendCaptcha = ref<boolean>(true)
  const captchaTimer = ref<number>(60)
  const captcha = ref<any>({
    captchaId: '',
    captchaImg: '',
  })
  const formParams: LoginData = reactive({
    emailCode: '',
    password: '',
    captchaCode: '',
    captchaId: '',
  })

  const regFormParams: RegData = reactive({
    username: '',
    password: '',
    rePassword: '',
    emailCode: '',
    captchaCode: '',
    captchaType: 'reg',
  })

  const emailFormParams: ReqUserEmailCaptcha = reactive({
    emailType: 'reg',
    emailCode: '',
  })

  const auth = useUserStore()

  const isShow = ref<boolean>(false)
  const showAnimate = ref<boolean>(false)

  loginType.value = props.type

  const getCaptchas = async () => {
    captcha.value = await getCaptcha()
  }

  const getEmailCaptcha = async (type) => {
    const email = RegExp('[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,3}$')
    if (!email.test(<string>regFormParams.emailCode)) {
      proxy.$message({
        type: 'warning',
        msg: '请输入正确的邮箱格式',
      })
      return
    }
    if (sendCaptcha.value) {
      emailFormParams.emailType = type
      emailFormParams.emailCode = regFormParams.emailCode
      const { code, message } = await updateUserEmailCaptcha(emailFormParams)
      if (code == 200) {
        proxy.$message({
          type: 'success',
          msg: '发送成功',
        })
        sendCaptcha.value = false
        const timer = setInterval(() => {
          captchaTimer.value--
          if (captchaTimer.value < 1) {
            clearInterval(timer)
            sendCaptcha.value = true
          }
        }, 1000)
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '发送失败',
        })
      }
    }
  }

  const loginUser = async () => {
    if (
      checkForm(
        formParams.emailCode,
        '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-z]{2,3}$',
        '请输入正确的邮箱格式',
      )
      && checkForm(formParams.captchaCode.toString(), '[0-9]{5,6}$', '请输入正确的验证码格式')
      && checkForm(
        formParams.password,
        '(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$',
        '请输入正确的密码格式',
      )
    ) {
      formParams.captchaId = captcha.value.captchaId
      formParams.captchaCode = formParams.captchaCode.toString()
      const { result, code, message } = await auth.login(formParams)
      if (code === 200 && result?.token) {
        props.close()
        proxy.$message({
          type: 'success',
          msg: '登陆成功',
        })
        setTimeout(() => {
          props.router.push(`${props.route.path}?t=${Date.parse(new Date().toString())}`)
        }, 1000)
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '登陆失败',
        })
        getCaptchas()
      }
    }
  }

  const reg = async () => {
    if (
      checkForm(regFormParams.username, '[a-zA-Z0-9_-]{4,16}$', '请输入正确的用户名格式')
      && checkForm(
        regFormParams.emailCode,
        '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-z]{2,3}$',
        '请输入正确的邮箱格式',
      )
      && checkForm(regFormParams.captchaCode?.toString(), '[0-9]{5,6}$', '请输入正确的验证码格式')
      && checkForm(
        regFormParams.password,
        '(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$',
        '请输入正确的密码格式',
      )
      && checkForm(
        regFormParams.rePassword,
        '(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$',
        '请输入正确的密码格式',
      )
    ) {
      regFormParams.captchaCode = regFormParams.captchaCode?.toString()
      const { code, message } = await auth.reg(regFormParams)
      if (code == 200) {
        props.close()
        proxy.$message({
          type: 'success',
          msg: '注册成功',
        })
        setTimeout(() => {
          props.router.push(`${props.route.path}?t=${Date.parse(new Date().toString())}`)
        }, 1000)
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '注册失败',
        })
      }
    }
  }

  const forget = async () => {
    if (
      checkForm(
        regFormParams.emailCode,
        '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-z]{2,3}$',
        '请输入正确的邮箱格式',
      )
      && checkForm(regFormParams.captchaCode?.toString(), '[0-9]{5,6}$', '请输入正确的验证码格式')
      && checkForm(
        regFormParams.password,
        '(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$',
        '请输入正确的密码格式',
      )
      && checkForm(
        regFormParams.rePassword,
        '(?=.*([a-zA-Z].*))(?=.*[0-9].*)[a-zA-Z0-9-*/+.~!@#$%^&*()]{6,20}$',
        '请输入正确的密码格式',
      )
    ) {
      regFormParams.captchaCode = regFormParams.captchaCode?.toString()
      regFormParams.captchaType = 'forget'
      const { code, message } = await auth.forget(regFormParams)
      if (code == 200) {
        props.close()
        proxy.$message({
          type: 'success',
          msg: '修改密码成功',
        })
        setTimeout(() => {
          props.router.push(`${props.route.path}?t=${Date.parse(new Date().toString())}`)
        }, 1000)
      }
      else {
        proxy.$message({
          type: 'warning',
          msg: message || '修改密码失败',
        })
      }
    }
  }

  const changeType = (type: string) => {
    loginType.value = type
    if (type === 'login')
      getCaptchas()
  }

  onMounted(() => {
    dialog.value.show()
    getCaptchas()
  })

  return {
    dialog,
    loginType,
    captcha,
    formParams,
    emailFormParams,
    regFormParams,
    captchaTimer,
    sendCaptcha,
    showAnime,
    formLogin,
    formReg,
    formForgetPwd,
    loginUser,
    reg,
    forget,
    getCaptchas,
    getEmailCaptcha,
    changeType,
    showAnimate,
    isShow,
  }
}
