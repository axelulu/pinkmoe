/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-21 14:16:37
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 19:52:37
 * @FilePath: /pinkmoe_index/hooks/login.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
 */
import { useUserStore } from '/@/store/modules/user'
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
        successMsg: '请输入正确的邮箱格式',
        loadFun: async () => {
          const code = 200
          return { code }
        },
      })
      return
    }
    if (sendCaptcha.value) {
      emailFormParams.emailType = type
      emailFormParams.emailCode = regFormParams.emailCode
      proxy.$message({
        successMsg: '发送成功',
        failedMsg: '发送失败',
        loadFun: async () => {
          const { code, message } = await updateUserEmailCaptcha(emailFormParams)
          return { code, message }
        },
      }).then((res) => {
        if (res.status === 200) {
          sendCaptcha.value = false
          const timer = setInterval(() => {
            captchaTimer.value--
            if (captchaTimer.value < 1) {
              clearInterval(timer)
              sendCaptcha.value = true
            }
          }, 1000)
        }
      })
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
      proxy.$message({
        successMsg: '登陆成功',
        failedMsg: '登陆失败',
        loadFun: async () => {
          const { code, message } = await auth.login(formParams)
          const result = props
          return { code, message, result }
        },
      }).then((res: any) => {
        if (res.status === 200) {
          setTimeout(() => {
            dialog.value.hide()
          }, 300)
          setTimeout(() => {
            res.data.router.push(`${res.data.route.path}?t=${Date.parse(new Date().toString())}`)
          }, 300)
        }
        else {
          getCaptchas()
        }
      })
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
      proxy.$message({
        successMsg: '注册成功',
        failedMsg: '注册失败',
        loadFun: async () => {
          const { code, message } = await auth.reg(regFormParams)
          return { code, message }
        },
      }).then((res) => {
        if (res.status === 200) {
          setTimeout(() => {
            dialog.value.hide()
          }, 300)
          setTimeout(() => {
            res.data.router.push(`${res.data.route.path}?t=${Date.parse(new Date().toString())}`)
          }, 300)
        }
      })
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
      proxy.$message({
        successMsg: '修改成功',
        failedMsg: '修改失败',
        loadFun: async () => {
          const { code, message } = await auth.forget(regFormParams)
          return { code, message }
        },
      }).then((res) => {
        if (res.status === 200) {
          setTimeout(() => {
            dialog.value.hide()
          }, 300)
          setTimeout(() => {
            res.data.router.push(`${res.data.route.path}?t=${Date.parse(new Date().toString())}`)
          }, 300)
        }
      })
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
