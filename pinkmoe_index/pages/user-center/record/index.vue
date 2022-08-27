<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-23 09:11:52
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-27 11:26:18
 * @FilePath: /pinkmoe_index/pages/user-center/record/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts" setup name="UserCenterRecordIndex">
import UserCenterLayout from '/@/components/Usercenterlayout/index.vue'
import MoreBtn from '/@/components/Morebtn/index.vue'
import Spin from '/@/components/Spin/index.vue'
import NotFound from '/@/components/NotFound/index.vue'
import { useUserCenterRecord } from '/@/hooks/user-center/record'
import { useAppStore } from '/@/store/modules/app'
import { useHead } from '@vueuse/head'
import { useUtil } from '/@/hooks/util'
const { formatDate } = useUtil()
const { authorRecordList, hasMore, loading, nextPage } = useUserCenterRecord()

definePageMeta({
  middleware: ['user-auth'],
})

const { siteBasic } = useAppStore()
useHead({
  title: '消费记录 - 用户中心',
  meta: [
    { name: 'og:type', content: 'record' },
    {
      name: 'og:title',
      content: '消费记录 - 用户中心',
    },
    { name: 'og:url', content: siteBasic?.url },
  ],
})
</script>

<template>
  <UserCenterLayout>
    <div class="ml-6">
      <div
        class="w-full mt-3 mb-2 inline-block bg-white dark:bg-gray-700 rounded-md shadow-sm relative"
      >
        <div class="absolute flex justify-center items-center -top-3 text-xs bg-sky-600 px-1.5 py-1 text-white ml-4 cursor-pointer">
          <i class="inline-block i-ri:chat-history-fill" />
          <span class="ml-1 select-none">消费记录</span>
        </div>
        <div class="p-4">
          <div
            class="px-4 py-3 bg-gray-100 mt-2 text-xs dark:bg-gray-800 dark:text-gray-200 text-gray-500"
          >
            <p>这里会展示您在本站进行的消费记录！</p>
          </div>
        </div>
        <div class="flex flex-wrap items-center px-4 pb-4">
          <Spin :show="loading">
            <div v-if="authorRecordList?.list && authorRecordList?.list?.length" class="w-full">
              <div
                v-for="(item, index) in authorRecordList?.list"
                :key="index"
                class="w-full text-xs text-gray-500 dark:hover:bg-gray-800 dark:text-gray-200 py-2 px-2 duration-300 cursor-pointer hover:bg-pink-100"
              >
                <div v-if="item.type === 'publishPost'" class="flex justify-between items-center">
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-majesticons:newspaper" />
                    <div>
                      您的文章<NuxtLink
                        class="text-pink-400"
                        :to="`/post/${item?.postRelation?.postId}`"
                      >
                        《{{ item?.postRelation?.title }}》
                      </NuxtLink>审核成功！
                    </div>
                  </div>
                  <div class="min-w-16 text-right">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'checkIn_credit'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-mdi:map-marker" />
                    <div>您获取了{{ item.credit }}积分签到奖励。</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'checkIn_cash'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-lucide:gem" />
                    <div>您获取了{{ item.cash }}现金签到奖励。</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div v-else-if="item.type === 'reg'" class="flex justify-between items-center">
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-ph:user-fill" />
                    <div>恭喜您注册成功！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'updateAvatar'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-ion:md-images" />
                    <div>恭喜您更新头像成功！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'updateEmail'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-jam:envelope-f" />
                    <div>恭喜您更新邮箱成功！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'updatePwd'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-material-symbols:key-rounded" />
                    <div>恭喜您更新密码成功！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'buyDownload'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-material-symbols:cloud-download" />
                    <div>
                      恭喜您成功购买了文章<NuxtLink
                        class="text-pink-400"
                        :to="`/post/${item?.postRelation?.postId}`"
                      >
                        《{{ item?.postRelation?.title }}》
                      </NuxtLink>的下载链接！
                    </div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div v-else-if="item.type === 'buyVideo'" class="flex justify-between items-center">
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-ic:round-videocam" />
                    <div>
                      恭喜您成功购买了文章<NuxtLink
                        class="text-pink-400"
                        :to="`/post/${item?.postRelation?.postId}`"
                      >
                        《{{ item?.postRelation?.title }}》
                      </NuxtLink>的下载链接！
                    </div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div v-else-if="item.type === 'buyMusic'" class="flex justify-between items-center">
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-ic:round-music-note" />
                    <div>
                      恭喜您成功购买了文章<NuxtLink
                        class="text-pink-400"
                        :to="`/post/${item?.postRelation?.postId}`"
                      >
                        《{{ item?.postRelation?.title }}》
                      </NuxtLink>的下载链接！
                    </div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'shopCash_key'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.cash > 0 ? `+${item.cash}` : item.cash
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-lucide:gem" />
                    <div>恭喜您成功使用卡密充值了{{ item?.cash }}现金！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'shopCredit_key'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.cash > 0 ? `+${item.cash}` : item.cash
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-lucide:gem" />
                    <div>恭喜您成功使用卡密充值了{{ item?.cash }}现金！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'shopCredit_cash'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-lucide:gem" />
                    <div>恭喜您使用{{ item?.cash }}现金成功购买了{{ item?.credit }}积分！</div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
                <div
                  v-else-if="item.type === 'comment' || item.type === 'commentPost'"
                  class="flex justify-between items-center"
                >
                  <div class="flex flex-row justify-center items-center">
                    <div class="text-xl text-green-600">
                      {{
                        item.credit > 0 ? `+${item.credit}` : item.credit
                      }}
                    </div>
                    <i class="inline-block text-xl px-4 i-uil:comments" />
                    <div>
                      您对用户<NuxtLink
                        class="text-pink-400"
                        :to="`/author/${item?.userRelation?.uuid}/userInfo`"
                      >
                        {{ item?.userRelation?.nickName }}
                      </NuxtLink>在文章<NuxtLink
                        class="text-pink-400"
                        :to="`/post/${item?.postRelation?.postId}`"
                      >
                        《{{ item?.postRelation?.title }}》
                      </NuxtLink>发表的评论进行了回复。
                    </div>
                  </div>
                  <div class="ml-4 min-w-16">
                    {{ formatDate(item?.UpdatedAt) }}
                  </div>
                </div>
              </div>
            </div>
            <NotFound v-else />
          </Spin>
        </div>
      </div>
      <div class="w-full text-gray-500">
        <MoreBtn v-if="hasMore" @click="nextPage" />
      </div>
    </div>
  </UserCenterLayout>
</template>
