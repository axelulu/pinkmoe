<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 08:13:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-28 17:13:17
 * @FilePath: /pinkmoe_index/components/Login/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * PinkMoe主题上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved.
-->
<script lang="ts" setup name="Login">
import Dialog from '/@/components/Dialog/index.vue'
import BasicInput from '/@/components/Basicinput/index.vue'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import { useLogin } from '/@/hooks/login'
const props = defineProps({
  type: {
    type: String,
    default: 'login',
  },
  router: {
    type: Object,
    default: null,
  },
  route: {
    type: Object,
    default: null,
  },
  close: {
    type: Function,
    default: null,
  },
})

const {
  formLogin,
  formForgetPwd,
  formReg,
  formParams,
  emailFormParams,
  regFormParams,
  dialog,
  loginType,
  captcha,
  captchaTimer,
  sendCaptcha,
  loginUser,
  reg,
  forget,
  getCaptchas,
  getEmailCaptcha,
  changeType,
  showAnimate,
} = useLogin(props)
</script>

<template>
  <Dialog
    ref="dialog"
    :hide-fun="() => (showAnimate = false)"
    :show-fun="() => (showAnimate = true)"
    @hide="close"
  >
    <div
      :class="showAnimate ? 'animate-pinkTipZoomInUp30' : 'animate-pinkTipZoomOutDown30'"
      class="fixed top-0 bottom-0 left-0 ring-0 flex justify-center items-center h-full w-full pointer-events-none"
    >
      <div
        class="lg:w-2/6 xl:w-2/12 bg-gray-50 dark:bg-gray-700 dark:text-gray-200 pointer-events-auto flex flex-col rounded-md overflow-hidden"
      >
        <div class="flex flex-row justify-between items-center">
          <div
            class="flex flex-row ml-4 bg-gray-200 dark:bg-gray-700 dark:text-gray-200 p-0 rounded-bl-md rounded-br-md overflow-hidden"
          >
            <div
              :class="loginType === 'login' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 select-none cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
              @click="changeType('login')"
            >
              登陆
            </div>
            <div
              :class="loginType === 'reg' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 select-none cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
              @click="changeType('reg')"
            >
              注册
            </div>
            <div
              :class="loginType === 'forgetPwd' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 select-none cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
              @click="changeType('forgetPwd')"
            >
              丢失密码
            </div>
          </div>
          <div
            class="text-sm dark:bg-gray-700 flex justify-center items-center py-2 dark:hover:bg-gray-800 dark:text-gray-200 px-4 py-1.5 bg-gray-200 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <i class="inline-block i-carbon:close-filled" />
          </div>
        </div>
        <form
          :ref="
            loginType === 'login' ? 'formLogin' : loginType === 'reg' ? 'formReg' : 'formForgetPwd'
          "
          class="flex-1 p-4 text-xs"
        >
          <div>欢迎登录 粉萌次元，登录后签个到看看运气如何吧！</div>
          <div v-if="loginType === 'login'" class="animate-fadeIn30">
            <BasicInput
              v-model:value="formParams.emailCode"
              icon="i-ic:baseline-email"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              v-model:value="formParams.password"
              icon="i-uis:unlock-alt"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              type="password"
              placeholder="请输入密码"
            />
            <BasicInput
              v-model:value="formParams.captchaCode"
              icon="i-material-symbols:key"
              :required="true"
              pattern="[0-9]{5,6}$"
              type="number"
              placeholder="请输入验证码"
              width="w-4/6"
            >
              <div class="w-2/6" @click="getCaptchas">
                <img
                  class="cursor-pointer px-1 object-cover bg-white"
                  :src="captcha.captchaImg"
                  alt=""
                >
              </div>
            </BasicInput>
          </div>
          <div v-if="loginType === 'reg'" class="animate-fadeIn30">
            <BasicInput
              v-model:value="regFormParams.username"
              icon="i-ph:user-plus-fill"
              :required="true"
              pattern="[a-zA-Z0-9_-]{4,16}$"
              type="text"
              placeholder="请输入用户名"
            />
            <BasicInput
              v-model:value="regFormParams.emailCode"
              icon="i-ic:baseline-email"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              v-model:value="regFormParams.captchaCode"
              icon="i-material-symbols:key"
              :required="true"
              pattern="[0-9]{5,6}$"
              type="number"
              width="w-3/6"
              placeholder="请输入验证码"
            >
              <div
                :class="sendCaptcha ? '' : 'cursor-not-allowed'"
                class="px-4 py-1.5 bg-sky-400 ml-4 w-3/6 text-xs text-white flex justify-center items-center cursor-pointer"
                @click="getEmailCaptcha('reg')"
              >
                {{ sendCaptcha ? '获取验证码' : `还有${captchaTimer}秒` }}
              </div>
            </BasicInput>
            <BasicInput
              v-model:value="regFormParams.password"
              icon="i-uis:unlock-alt"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              type="password"
              placeholder="请输入密码 "
            />
            <BasicInput
              v-model:value="regFormParams.rePassword"
              icon="i-uis:unlock-alt"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              type="password"
              placeholder="请确认密码"
            />
          </div>
          <div v-if="loginType === 'forgetPwd'" class="animate-fadeIn30">
            <BasicInput
              v-model:value="regFormParams.emailCode"
              icon="i-ic:baseline-email"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              v-model:value="regFormParams.captchaCode"
              icon="i-material-symbols:key"
              :required="true"
              pattern="[0-9]{5,6}$"
              type="number"
              width="w-3/6"
              placeholder="请输入验证码"
            >
              <div
                :class="sendCaptcha ? '' : 'cursor-not-allowed'"
                class="px-4 py-1.5 bg-sky-400 ml-4 w-3/6 text-xs text-white flex justify-center items-center cursor-pointer"
                @click="getEmailCaptcha('forget')"
              >
                {{ sendCaptcha ? '获取验证码' : `还有${captchaTimer}秒` }}
              </div>
            </BasicInput>
            <BasicInput
              v-model:value="regFormParams.password"
              icon="i-uis:unlock-alt"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              type="password"
              placeholder="请输入密码 "
            />
            <BasicInput
              v-model:value="regFormParams.rePassword"
              icon="i-uis:unlock-alt"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              type="password"
              placeholder="请确认密码"
            />
          </div>
        </form>
        <GreenBtn
          v-if="loginType === 'login'"
          classes="border-0"
          value="登陆"
          icon="i-ic:sharp-arrow-circle-right"
          class="h-9"
          @click="loginUser"
        />
        <GreenBtn
          v-if="loginType === 'reg'"
          classes="border-0"
          value="注册"
          icon="i-ph:user-plus-fill"
          class="h-9"
          @click="reg"
        />
        <GreenBtn
          v-if="loginType === 'forgetPwd'"
          classes="border-0"
          value="忘记密码"
          icon="i-mdi:sync-circle"
          class="h-9"
          @click="forget"
        />
      </div>
    </div>
  </Dialog>
</template>
