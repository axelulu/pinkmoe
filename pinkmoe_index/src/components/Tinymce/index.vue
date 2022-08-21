<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-08-19 19:23:35
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-19 19:28:32
 * @FilePath: /pinkmoe_vitesse/src/components/Tinymce/index.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved.
-->
<script lang="ts">
import type { PropType } from 'vue'
import {
  computed, defineComponent, onBeforeUnmount, onMounted, ref, unref, watch,
} from 'vue'
import type { Editor, RawEditorSettings } from 'tinymce'
import tinymce from 'tinymce/tinymce'

import { fontFormats, plugins as initialPlugins, toolbar as initialToolbar } from './tinymce'

  type Recordable<T = any> = Record<string, T>

export default defineComponent({
  props: {
    value: {
      type: String,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    options: {
      type: Object as PropType<Partial<RawEditorSettings>>,
      default: () => ({}),
    },
    toolbar: {
      type: String,
      default: initialToolbar,
    },
    plugins: {
      type: Array as PropType<string[]>,
      default: initialPlugins,
    },
    height: {
      type: [Number, String] as PropType<string | number>,
      required: false,
      default: 400,
    },
    width: {
      type: [Number, String] as PropType<string | number>,
      required: false,
      default: 'auto',
    },
  },
  emits: ['change', 'update:value'],
  setup(props, { emit }) {
    const tinymceId = ref<string>()
    const editorRef = ref<Editor>()

    const initOptions = computed((): RawEditorSettings => {
      const publicPath = '/@/'

      const {
        height, options, toolbar, plugins,
      } = props
      return {
        selector: `#${tinymceId.value}`,
        language_url: `${publicPath}tinymce/langs/zh_CN.js`,
        language: 'zh_CN',
        skin_url: `${publicPath}tinymce/skins/ui/oxide`,
        content_css: `${publicPath}tinymce/skins/ui/oxide/content.min.css`,
        images_upload_handler: handleImgUpload,
        images_file_types: 'jpeg,jpg,png,gif,bmp,webp', // 准许的图片格式
        convert_urls: false,
        branding: false, // 隐藏品牌，隐藏状态栏中显示的“ Powered by Tiny ”链接
        placeholder: '请输入内容', // 占位符
        toolbar,
        plugins,
        height,
        toolbar_mode: 'sliding',
        toolbar_sticky: true,
        paste_block_drop: true, // 禁用将内容拖放到编辑器中
        paste_data_images: false, // 粘贴data格式的图像 谷歌浏览器无法粘贴
        font_formats: fontFormats,
        paste_retain_style_properties: 'color border border-left border-right border-bottom border-top', // MS Word 和类似 Office 套件产品保留样式
        paste_webkit_styles: 'none', // 允许在 WebKit 中粘贴时要保留的样式
        paste_tab_spaces: 2, // 将制表符转换成空格的个数
        content_style: `
          html, body                { height:100%; }
          img                       { max-width:100%; display:block;height:auto; }
          a                         { text-decoration: none; }
          p                         { line-height:1.6; margin: 0px; }
          table                     { word-wrap:break-word; word-break:break-all;max-width:100%; border:none; border-color:#999; }
          .mce-object-iframe        { width:100%; box-sizing:border-box; margin:0; padding:0; }
          ul,ol                     { list-style-position:inside; }
          `,
        ...options,
        setup: (editor: Editor) => {
          editorRef.value = editor
          editor.on('init', initSetup)
        },
      }
    })

    onMounted(() => {
      tinymce.init(initOptions.value)
    })

    onBeforeUnmount(() => {
      destory()
    })

    function destory() {
      if (tinymce !== null)
        tinymce?.remove?.(unref(initOptions).selector!)
    }

    // 图片上传自定义逻辑
    function handleImgUpload(blobInfo, success, failure, progress) {
      console.log('blobInfo', blobInfo.blob(), blobInfo.filename())
      const { type: fileType, name: fileName } = blobInfo.blob()
      // xxxx 自定义上传逻辑
    }

    // 编辑器初始化
    function initSetup() {
      const editor = unref(editorRef)
      if (!editor)
        return

      const value = props.value || ''

      editor.setContent(value)
      bindModelHandlers(editor)
    }

    function setValue(editor: Recordable, val: string, prevVal?: string) {
      if (
        editor
          && typeof val === 'string'
          && val !== prevVal
          && val !== editor.getContent()
      )
        editor.setContent(val)
    }

    function bindModelHandlers(editor: any) {
      watch(
        () => props.value,
        (val: string, prevVal) => setValue(editor, val, prevVal),
        { immediate: true },
      )
      watch(
        () => props.disabled,
        (val) => {
          editor.setMode(val ? 'readonly' : 'design')
        },
        { immediate: true },
      )
      editor.on('change keyup undo redo', () => {
        const content = editor.getContent()
        emit('update:value', content)
        emit('change', content)
      })
    }

    return {
      tinymceId,
    }
  },
})
</script>

<template>
  <textarea :id="tinymceId" />
</template>
