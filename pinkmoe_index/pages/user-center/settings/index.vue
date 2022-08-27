<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:11:59
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:26:20
 * @FilePath: /pinkmoe_index/pages/user-center/settings/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterSettingsIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import BasicInput from '/@/components/Basicinput/index.vue'
import GreenBtn from '/@/components/Greenbtn/index.vue'
import TextareaInput from '/@/components/TextareaInput/index.vue'
import { useUserCenterSettings } from '/@/hooks/user-center/settings'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
const {
  formUserDetail,
  formUserPwd,
  formUserEmail,
  detailFormParams,
  emailFormParams,
  avatarFormParams,
  pwdFormParams,
  sendCaptcha,
  captchaTimer,
  updateDetail,
  uploadPostImg,
  updatePwd,
  getEmailCaptcha,
  updateEmail,
  loading,
} = useUserCenterSettings()

definePageMeta({
  middleware: ['user-auth'],
})

const { siteBasic } = useAppStore()
useHead({
  title: '设置 - 用户中心',
  meta: [
    { name: 'og:type', content: 'settings' },
    {
      name: 'og:title',
      content: '设置 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <UserCenterLayout>
    <Spin :show="loading" class="flex flex-wrap">
      <div class="ml-6 columns-2 gap-4">
        <div
          class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
        >
          <div
            class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer"
          >
            <i class="inline-block i-ph:user-fill" />
            <span class="ml-1 select-none">我的头像</span>
          </div>
          <div class="p-4">
            <div
              class="px-4 py-3 bg-gray-100 mt-2 text-xs text-gray-500 dark:bg-gray-800 dark:text-gray-200"
            >
              修改头像需要 10 个积分哦！
            </div>
          </div>
          <div class="flex flex-row items-center px-4 pb-4">
            <img
              v-lazy="avatarFormParams.avatar"
              class="w-28 h-28 object-cover rounded-full animate-lazyloaded cursor-pointer object-cover"
              alt=""
            >
            <label
              for="avatar"
              class="border-2 border-dashed hover:bg-gray-100 dark:border-gray-200 dark:hover:bg-gray-800 duration-300 border-gray-500 ml-4 cursor-pointer h-28 w-full flex justify-center items-center"
            >
              <i class="text-2xl inline-block text-gray-500 dark:text-gray-200 i-material-symbols:photo-camera-rounded" />
              <span class="ml-2 text-lg text-gray-500 dark:text-gray-200 select-none">更改我的头像</span>
              <input
                id="avatar"
                accept="image/gif, image/jpg, image/jpeg, image/png, image/webp"
                type="file"
                hidden
                value=""
                @change="uploadPostImg"
              >
            </label>
          </div>
        </div>
        <div
          class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
        >
          <div
            class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer"
          >
            <i class="inline-block i-ph:user-fill" />
            <span class="ml-1 select-none">个人资料</span>
          </div>
          <div class="p-4">
            <div
              class="px-4 py-3 bg-gray-100 mt-2 text-xs text-gray-500 dark:bg-gray-800 dark:text-gray-200"
            >
              修改一个你个性化的昵称和描述吧！
            </div>
            <form ref="formUserDetail" onsubmit="return false" class="mt-4">
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  昵称
                </div>
                <BasicInput
                  v-model:value="detailFormParams.nickName"
                  :required="true"
                  pattern="[\u4e00-\u9fa5]{1,7}$|^[\dA-Za-z_]{1,14}$"
                  class="w-10/12 text-xs"
                  icon="i-bx:bxs-user-circle"
                  type="text"
                  placeholder="请输入昵称"
                />
              </div>
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  描述
                </div>
                <TextareaInput
                  v-model:value="detailFormParams.desc"
                  :required="true"
                  pattern="[\u4e00-\u9fa5]{3,100}$|^[\dA-Za-z_]{3,100}$"
                  class="w-10/12 text-xs"
                  rows="5"
                  maxlen="120"
                  placeholder="请输入描述"
                />
              </div>
              <div class="flex justify-end mt-4">
                <GreenBtn
                  classes="w-10/12"
                  value="更新个人资料"
                  icon="i-material-symbols:check-circle"
                  @click="updateDetail"
                />
              </div>
            </form>
          </div>
        </div>
        <div
          class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
        >
          <div
            class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer"
          >
            <i class="inline-block i-ic:baseline-email" />
            <span class="ml-1 select-none">账户邮箱</span>
          </div>
          <div class="p-4">
            <div
              class="px-4 py-3 bg-gray-100 mt-2 text-xs text-gray-500 dark:bg-gray-800 dark:text-gray-200"
            >
              修改账户邮箱，需要 10 个积分哦！
            </div>
            <form ref="formUserEmail" onsubmit="return false" class="mt-4">
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  邮箱
                </div>
                <BasicInput
                  v-model:value="emailFormParams.emailCode"
                  :required="true"
                  pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,3}$"
                  class="w-10/12 text-xs"
                  icon="i-ic:baseline-email"
                  type="email"
                  placeholder="请输入邮箱"
                />
              </div>
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  验证码
                </div>
                <BasicInput
                  v-model:value="emailFormParams.emailCaptcha"
                  :required="true"
                  pattern="[0-9]{5,6}$"
                  class="w-6/12 text-xs"
                  icon="i-material-symbols:lock"
                  type="number"
                  placeholder="请输入验证码"
                />
                <div class="w-4/12 pl-4 flex justify-center items-center">
                  <div
                    :class="sendCaptcha ? '' : 'opacity-50 cursor-not-allowed'"
                    class="w-full select-none flex justify-center items-center text-xs text-center cursor-pointer text-gray-500 dark:border-gray-800 dark:text-gray-200 border-2 border-gray-100 px-2 py-1.5"
                    @click="getEmailCaptcha"
                  >
                    <i class="inline-block i-jam:envelope-f" />
                    <span> {{ sendCaptcha ? '发送验证码' : captchaTimer }} </span>
                  </div>
                </div>
              </div>
              <div class="flex justify-end mt-4">
                <GreenBtn
                  classes="w-10/12"
                  value="更新邮箱"
                  icon="i-material-symbols:check-circle"
                  @click="updateEmail"
                />
              </div>
            </form>
          </div>
        </div>
        <div
          class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
        >
          <div
            class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer"
          >
            <i class="inline-block i-material-symbols:lock" />
            <span class="ml-1 select-none">我的密码</span>
          </div>
          <div class="p-4">
            <div
              class="px-4 py-3 bg-gray-100 mt-2 text-xs text-gray-500 dark:bg-gray-800 dark:text-gray-200"
            >
              您可以修改您的密码，建议复杂一些，需要 10 个积分哦！。
            </div>
            <form ref="formUserPwd" onsubmit="return false" class="mt-4">
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  旧密码
                </div>
                <BasicInput
                  v-model:value="pwdFormParams.oldPassword"
                  :required="true"
                  pattern="[a-zA-Z]\w{5,17}$"
                  class="w-10/12 text-xs"
                  icon="i-material-symbols:lock"
                  type="password"
                  placeholder="请输入旧密码"
                />
              </div>
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  新密码
                </div>
                <BasicInput
                  v-model:value="pwdFormParams.password"
                  :required="true"
                  pattern="[a-zA-Z]\w{5,17}$"
                  class="w-10/12 text-xs"
                  icon="i-material-symbols:lock"
                  type="password"
                  placeholder="请输入新密码"
                />
              </div>
              <div class="flex flex-row items-center">
                <div class="text-xs w-2/12 text-gray-500 dark:text-gray-200">
                  确认新密码
                </div>
                <BasicInput
                  v-model:value="pwdFormParams.reNewPwd"
                  :required="true"
                  pattern="[a-zA-Z]\w{5,17}$"
                  class="w-10/12 text-xs"
                  icon="i-material-symbols:lock"
                  type="password"
                  placeholder="请确认新密码"
                />
              </div>
              <div class="flex justify-end mt-4">
                <GreenBtn
                  classes="w-10/12"
                  value="更新密码"
                  icon="i-material-symbols:check-circle"
                  @click="updatePwd"
                />
              </div>
            </form>
          </div>
        </div>
      </div>
    </Spin>
  </UserCenterLayout>
</template>
