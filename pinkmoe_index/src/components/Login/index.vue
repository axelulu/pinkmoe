<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-22 08:13:00
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:10:38
 * @FilePath: /pinkmoe_index/src/components/Login/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<script lang="ts" setup name="Login">
  import Dialog from '/@/components/Dialog/index.vue';
  import BasicInput from '/@/components/Basicinput/index.vue';
  import GreenBtn from '/@/components/Greenbtn/index.vue';
  import { useLogin } from '/@/hooks/login';
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
  });

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
    showAnimate,
  } = useLogin(props);
</script>

<template>
  <!-- Login -->
  <Dialog
    ref="dialog"
    @hide="close"
    :hide-fun="() => (showAnimate = false)"
    :show-fun="() => (showAnimate = true)"
  >
    <div
      :class="showAnimate ? 'animate-pinkZoomInUp30' : 'animate-pinkZoomOutDown30'"
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
              @click="loginType = 'login'"
              :class="loginType === 'login' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
            >
              登陆
            </div>
            <div
              @click="loginType = 'reg'"
              :class="loginType === 'reg' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
            >
              注册
            </div>
            <div
              @click="loginType = 'forgetPwd'"
              :class="loginType === 'forgetPwd' ? 'bg-pink-400 text-white' : ''"
              class="text-sm px-4 py-1.5 cursor-pointer hover:text-white hover:bg-pink-400 duration-300"
            >
              丢失密码
            </div>
          </div>
          <div
            class="text-sm dark:bg-gray-700 dark:hover:bg-gray-800 dark:text-gray-200 px-4 py-1.5 bg-gray-200 cursor-pointer hover:text-white hover:bg-pink-400 duration-300 rounded-bl-md"
            @click="dialog.hide()"
          >
            <font-awesome-icon :icon="['fas', 'times']" />
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
              :icon="['fas', 'at']"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              v-model:value="formParams.emailCode"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              :icon="['fas', 'unlock-alt']"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              v-model:value="formParams.password"
              type="password"
              placeholder="请输入密码"
            />
            <BasicInput
              :icon="['fas', 'shield-alt']"
              :required="true"
              pattern="[0-9]{5,6}$"
              v-model:value="formParams.captchaCode"
              type="number"
              placeholder="请输入验证码"
              width="w-4/6"
            >
              <div class="w-2/6" @click="getCaptchas">
                <img
                  class="cursor-pointer px-1 object-cover bg-white"
                  :src="captcha.captchaImg"
                  alt=""
                />
              </div>
            </BasicInput>
          </div>
          <div v-if="loginType === 'reg'" class="animate-fadeIn30">
            <BasicInput
              :icon="['fas', 'user']"
              :required="true"
              pattern="[a-zA-Z0-9_-]{4,16}$"
              v-model:value="regFormParams.username"
              type="text"
              placeholder="请输入用户名"
            />
            <BasicInput
              :icon="['fas', 'at']"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              v-model:value="regFormParams.emailCode"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              :icon="['fas', 'key']"
              :required="true"
              pattern="[0-9]{5,6}$"
              v-model:value="regFormParams.captchaCode"
              type="number"
              width="w-3/6"
              placeholder="请输入验证码"
            >
              <div
                @click="getEmailCaptcha('reg')"
                :class="sendCaptcha ? '' : 'cursor-not-allowed'"
                class="px-4 py-1.5 bg-sky-400 ml-4 w-3/6 text-xs text-white flex justify-center items-center cursor-pointer"
              >
                {{ sendCaptcha ? '获取验证码' : '还有' + captchaTimer + '秒' }}
              </div>
            </BasicInput>
            <BasicInput
              :icon="['fas', 'lock']"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              v-model:value="regFormParams.password"
              type="password"
              placeholder="请输入密码 "
            />
            <BasicInput
              :icon="['fas', 'lock']"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              v-model:value="regFormParams.rePassword"
              type="password"
              placeholder="请确认密码"
            />
          </div>
          <div v-if="loginType === 'forgetPwd'" class="animate-fadeIn30">
            <BasicInput
              :icon="['fas', 'at']"
              :required="true"
              pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
              v-model:value="regFormParams.emailCode"
              type="email"
              placeholder="请输入邮箱"
            />
            <BasicInput
              :icon="['fas', 'key']"
              :required="true"
              pattern="[0-9]{5,6}$"
              v-model:value="regFormParams.captchaCode"
              type="number"
              width="w-3/6"
              placeholder="请输入验证码"
            >
              <div
                :class="sendCaptcha ? '' : 'cursor-not-allowed'"
                @click="getEmailCaptcha('forget')"
                class="px-4 py-1.5 bg-sky-400 ml-4 w-3/6 text-xs text-white flex justify-center items-center cursor-pointer"
              >
                {{ sendCaptcha ? '获取验证码' : '还有' + captchaTimer + '秒' }}
              </div>
            </BasicInput>
            <BasicInput
              :icon="['fas', 'lock']"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              v-model:value="regFormParams.password"
              type="password"
              placeholder="请输入密码 "
            />
            <BasicInput
              :icon="['fas', 'lock']"
              :required="true"
              pattern="[a-zA-Z]\w{5,17}$"
              v-model:value="regFormParams.rePassword"
              type="password"
              placeholder="请确认密码"
            />
          </div>
        </form>
        <GreenBtn
          v-if="loginType === 'login'"
          :classes="'border-0'"
          @click="loginUser"
          value="登陆"
          :icon="['fas', 'arrow-alt-circle-right']"
          class="h-9"
        />
        <GreenBtn
          v-if="loginType === 'reg'"
          :classes="'border-0'"
          @click="reg"
          value="注册"
          :icon="['fas', 'user-plus']"
          class="h-9"
        />
        <GreenBtn
          v-if="loginType === 'forgetPwd'"
          :classes="'border-0'"
          @click="forget"
          value="忘记密码"
          :icon="['fas', 'sync']"
          class="h-9"
        />
      </div>
    </div>
  </Dialog>
</template>

<style lang="less" scoped></style>
