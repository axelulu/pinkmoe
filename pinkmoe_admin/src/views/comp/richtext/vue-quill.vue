<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-04-28 19:46:27
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:22:33
 * @FilePath: /pinkmoe_admin/src/views/comp/richtext/vue-quill.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe 
 * 问题反馈qq群:749150798
 * xanaduCms程序上所有内容(包括但不限于 文字，图片，代码等)均为指针科技原创所有，采用请注意许可
 * 请遵循 “非商业用途” 协议。商业网站或未授权媒体不得复制内容，如需用于商业用途或者二开，请联系作者捐助任意金额即可，我们将保存所有权利。
 * Copyright (c) 2022 by 指针科技, All Rights Reserved. 
-->
<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="富文本">
        富文本，用于展示图文信息，比如商品详情，文章详情等...
      </n-card>
    </div>
    <n-card :bordered="false" class="mt-4 proCard">
      <QuillEditor
        ref="quillEditor"
        :options="options"
        v-model:content="myContent"
        style="height: 350px"
        @ready="readyQuill"
        class="quillEditor"
      />
      <template #footer>
        <n-space>
          <n-button @click="addText">增加文本</n-button>
          <n-button @click="addImg">增加图片</n-button>
          <n-button @click="getHtml">获取HTML</n-button>
        </n-space>
      </template>
    </n-card>
    <n-card :bordered="false" class="mt-4 proCard" title="HTML 内容">
      <n-input
        v-model:value="myContentHtml"
        type="textarea"
        placeholder="html"
        :autosize="{
          minRows: 3,
          maxRows: 6,
        }"
      />
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, reactive } from 'vue';
  import { QuillEditor } from '@vueup/vue-quill';
  import '@vueup/vue-quill/dist/vue-quill.snow.css';
  const quillEditor = ref();
  const myContent = ref(
    '<h4>Naive Ui Admin 是一个基于 vue3,vite2,TypeScript 的中后台解决方案</h4>'
  );
  const myContentHtml = ref(
    '<h4>Naive Ui Admin 是一个基于 vue3,vite2,TypeScript 的中后台解决方案</h4>'
  );

  const options = reactive({
    modules: {
      toolbar: {
        container: [
          ['bold', 'italic', 'underline', 'strike'], // toggled buttons
          ['blockquote', 'code-block'],

          [{ header: 1 }, { header: 2 }], // custom button values
          [{ list: 'ordered' }, { list: 'bullet' }],
          [{ script: 'sub' }, { script: 'super' }], // superscript/subscript
          [{ indent: '-1' }, { indent: '+1' }], // outdent/indent
          [{ direction: 'rtl' }], // text direction

          [{ size: ['small', false, 'large', 'huge'] }], // custom dropdown
          [{ header: [1, 2, 3, 4, 5, 6, false] }],

          [{ color: [] }, { background: [] }], // dropdown with defaults from theme
          [{ font: [] }],
          [{ align: [] }],
          ['clean'],
          ['image'],
        ],
        handlers: {
          image: function (value) {
            if (value) {
              // 触发input框选择图片文件
              console.log(value);
              // document.querySelector('.avatar-uploader input').click();
            } else {
              // this.quill.format('image', false);
            }
          },
          // link: function(value) {
          //   if (value) {
          //     var href = prompt('请输入url');
          //     this.quill.format("link", href);
          //   } else {
          //     this.quill.format("link", false);
          //   }
          // },
        },
      },
    },
    serverUrl: '/v1/blog/imgUpload', // 这里写你要上传的图片服务器地址
    header: {
      token: sessionStorage.token,
    },
    theme: 'snow',
    placeholder: '输入您喜欢的内容吧！',
  });

  function readyQuill() {
    console.log('Quill准备好了');
  }

  function getHtml() {
    myContentHtml.value = getHtmlVal();
  }

  function addText() {
    const html = getHtmlVal() + '新增加的内容';
    quillEditor.value.setHTML(html);
  }

  function addImg() {
    const html =
      getHtmlVal() +
      '<img style="width:100px" src="https://www.baidu.com/img/flexible/logo/pc/result.png"/>';
    quillEditor.value.setHTML(html);
  }

  function getHtmlVal() {
    return quillEditor.value.getHTML();
  }
</script>

<style lang="less">
  .ql-toolbar.ql-snow {
    border-top: none;
    border-left: none;
    border-right: none;
    border-bottom: 1px solid #eee;
    margin-top: -10px;
  }

  .ql-container.ql-snow {
    border: none;
  }
</style>
