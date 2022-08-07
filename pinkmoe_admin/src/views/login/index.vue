<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-10 19:31:31
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:24:30
 * @FilePath: /pinkmoe_admin/src/views/login/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <div class="view-account">
    <div class="view-account-header"></div>
    <div class="view-account-container">
      <div class="view-account-top">
        <div class="view-account-top-logo">
          <img src="~@/assets/images/account-logo.png" alt="" />
        </div>
        <div class="view-account-top-desc">XanaDuCms 程序后台</div>
      </div>
      <div class="view-account-form">
        <n-form
          ref="formRef"
          label-placement="left"
          size="large"
          :model="formInline"
          :rules="rules"
        >
          <n-form-item v-if="loginType.type === 'login'" path="emailCode">
            <n-input v-model:value="formInline.emailCode" placeholder="请输入邮箱">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <PersonOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item v-if="loginType.type === 'reg'" path="username">
            <n-input v-model:value="formInline.username" placeholder="请输入用户名">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <PersonOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item v-if="loginType.type === 'reg'" :show-feedback="false">
            <n-grid :cols="6" :x-gap="24">
              <n-form-item-gi :span="4" path="emailCode">
                <n-input v-model:value="formInline.emailCode" placeholder="请填写邮箱">
                  <template #prefix>
                    <n-icon size="18" color="#808695">
                      <MailOutline />
                    </n-icon>
                  </template>
                </n-input>
              </n-form-item-gi>
              <n-form-item-gi :span="2">
                <n-button :disabled="!formInline.sendCaptcha" @click="getEmailCode">{{
                    formInline.sendCaptcha ? "获取验证码" : "还有" + formInline.captchaTimer + "秒"
                  }}
                </n-button>
              </n-form-item-gi>
            </n-grid>
          </n-form-item>
          <n-form-item v-if="loginType.type === 'reg'" path="captchaCode">
            <n-input class="w-full" v-model:value="formInline.captchaCode" placeholder="请输入验证码">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <ShieldCheckmarkOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="password">
            <n-input
              v-model:value="formInline.password"
              type="password"
              showPasswordOn="click"
              placeholder="请输入密码"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item v-if="loginType.type === 'reg'" path="rePassword">
            <n-input
              v-model:value="formInline.rePassword"
              type="password"
              showPasswordOn="click"
              placeholder="请再次输入密码"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item v-if="loginType.type === 'login'" :show-feedback="false">
            <n-grid :cols="5" :x-gap="24">
              <n-form-item-gi :span="3" path="captchaCode">
                <n-input v-model:value="formInline.captchaCode" placeholder="验证码" />
              </n-form-item-gi>
              <n-form-item-gi :span="2">
                <n-image @click="getCapt" width="200" preview-disabled :src="captcha.captchaImg" />
              </n-form-item-gi>
            </n-grid>
          </n-form-item>
          <n-form-item class="default-color">
            <div class="flex justify-between">
              <div class="flex-initial">
                <n-checkbox v-model:checked="autoLogin">自动登录</n-checkbox>
              </div>
              <div class="flex-initial order-last">
                <a href="javascript:">忘记密码</a>
              </div>
            </div>
          </n-form-item>
          <n-form-item>
            <n-button type="primary" @click="handleSubmit" size="large" :loading="loading" block>
              {{
                loginType.type === "login" ? "登录" : loginType.type === "reg" ? "注册" : "忘记密码"
              }}
            </n-button>
          </n-form-item>
          <n-form-item class="default-color">
            <div class="flex view-account-other">
              <div class="flex-initial" style="margin-left: auto">
                <a v-if="loginType.type === 'login'" @click="loginType.type = 'reg'">注册账号</a>
                <a v-if="loginType.type === 'reg'" @click="loginType.type = 'login'">登陆账号</a>
              </div>
            </div>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useUserStore } from "@/store/modules/user";
import { useMessage } from "naive-ui";
import {
  PersonOutline,
  LockClosedOutline,
  LogoGithub,
  LogoFacebook,
  MailOutline,
  ShieldCheckmarkOutline
} from "@vicons/ionicons5";
import { ResultEnum } from "@/enums/httpEnum";
import { PageEnum } from "@/enums/pageEnum";
import { getCaptcha, getEmailCaptcha } from "@/api/system/user";

interface LoginFormState {
  emailCode: string;
  password: string;
  captchaId: string;
  captchaCode: string;
}

interface RegFormState {
  username: string;
  password: string;
  rePassword: string;
  emailCode: string;
  captchaCode: string;
  captchaType: string;
}

interface EmailFormState {
  emailCode: string;
  emailType: string;
}

const formRef = ref();
const message = useMessage();
const loading = ref(false);
const autoLogin = ref(true);
const LOGIN_NAME = PageEnum.BASE_LOGIN_NAME;

const formInline = reactive({
  username: "admin",
  password: "123456",
  rePassword: "123456",
  emailCode: "",
  isCaptcha: true,
  captchaId: "",
  captchaCode: "",
  sendCaptcha: true,
  captchaTimer: 60
});

const captcha = reactive({
  captchaImg: ""
});

const loginType = reactive({
  type: "login"
});

const rules = {
  username: {
    required: true,
    message: "4到16位（字母，数字，下划线，减号）",
    trigger: "blur",
    pattern: /^[a-zA-Z0-9_-]{4,16}$/ //验证用户名 4到16位（字母，数字，下划线，减号）
  },
  password: {
    required: true,
    message: "6-20位，包含大小写字母和数字",
    trigger: "blur",
    pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[\s\S]{6,20}$/ //验证密码 6-20位，包含大小写字母和数字
  },
  rePassword: {
    required: true,
    message: "6-20位，包含大小写字母和数字",
    trigger: "blur",
    pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[\s\S]{6,20}$/ //验证密码 6-20位，包含大小写字母和数字
  },
  emailCode: {
    required: true,
    type: "email",
    message: "请检查邮箱格式",
    trigger: "blur"
    // pattern: /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/, //验证邮箱
  },
  captchaCode: {
    required: true,
    message: "6位验证码，包含数字",
    trigger: "blur",
    pattern: /^[0-9]{6}$/ //验证验证码 6位验证码，包含数字
  }
};

const userStore = useUserStore();

const router = useRouter();
const route = useRoute();

const getCapt = async () => {
  const { captchaId, captchaImg } = await getCaptcha();
  formInline.captchaId = captchaId;
  captcha.captchaImg = captchaImg;
};

getCapt();

const getEmailCode = async () => {
  const { emailCode } = formInline;
  if (emailCode.length < 1) {
    message.info("请填写邮箱");
    return;
  }
  const params: EmailFormState = {
    emailCode,
    emailType: loginType.type
  };
  const { code, message: msg } = await getEmailCaptcha(params);
  if (code == ResultEnum.SUCCESS) {
    message.success("发送验证码成功");
    formInline.sendCaptcha = false;
    const timer = setInterval(() => {
      formInline.captchaTimer--;
      if (formInline.captchaTimer < 1) {
        clearInterval(timer);
        formInline.sendCaptcha = true;
      }
    }, 1000);
  } else {
    message.error(msg || "发送验证码失败");
  }
};

const submitReg = async () => {
  const { username, password, rePassword, emailCode, captchaCode } = formInline;
  message.loading("注册中...");
  loading.value = true;
  const params: RegFormState = {
    username,
    password,
    rePassword,
    emailCode,
    captchaCode,
    captchaType: loginType.type
  };
  try {
    const { code, message: msg } = await userStore.reg(params);
    message.destroyAll();
    if (code == ResultEnum.SUCCESS) {
      message.success("注册成功，即将返回登陆");
      loginType.type = "login";
    } else {
      message.error(msg || "注册失败");
    }
  } finally {
    loading.value = false;
  }
};

const submitLogin = async () => {
  const { emailCode, password, captchaId, captchaCode } = formInline;
  message.loading("登录中...");
  loading.value = true;

  const params: LoginFormState = {
    emailCode,
    password,
    captchaId,
    captchaCode
  };

  try {
    const { code, message: msg } = await userStore.login(params);
    message.destroyAll();
    if (code == ResultEnum.SUCCESS) {
      const toPath = decodeURIComponent((route.query?.redirect || "/") as string);
      message.success("登录成功，即将进入系统");
      if (route.name === LOGIN_NAME) {
        await router.replace("/");
      } else await router.replace(toPath);
    } else {
      await getCapt();
      message.error(msg || "登录失败");
    }
  } finally {
    loading.value = false;
  }
};

const submitForget = async () => {
};

const handleSubmit = (e) => {
  e.preventDefault();
  formRef.value.validate(async (errors) => {
    if (!errors) {
      switch (loginType.type) {
        case "login":
          await submitLogin();
          break;
        case "reg":
          await submitReg();
          break;
        case "forget":
          await submitForget();
          break;
      }
    } else {
      message.error("请填写完整信息，并进行格式验证");
    }
  });
};
</script>

<style lang="less" scoped>
.view-account {
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: auto;

  &-container {
    flex: 1;
    padding: 32px 0;
    width: 384px;
    margin: 0 auto;
  }

  &-top {
    padding: 32px 0;
    text-align: center;

    &-desc {
      font-size: 14px;
      color: #808695;
    }
  }

  &-other {
    width: 100%;
  }

  .default-color {
    color: #515a6e;

    .ant-checkbox-wrapper {
      color: #515a6e;
    }
  }
}

@media (min-width: 768px) {
  .view-account {
    background-image: url('../../assets/images/login.svg');
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100%;
  }

  .page-account-container {
    padding: 32px 0 24px 0;
  }
}
</style>
