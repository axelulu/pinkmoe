<!--
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-06-04 12:07:03
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 09:23:05
 * @FilePath: /pinkmoe_admin/src/views/content/post/wangEditor.vue
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
-->
<template>
  <div style="border: 1px solid #ccc; margin-top: 10px;">
    <Toolbar :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode"
             style="border-bottom: 1px solid #ccc" />
    <Editor :defaultConfig="editorConfig" :mode="mode" v-model="valueHtml"
            style="height: 400px; overflow-y: hidden;" @onCreated="handleCreated"
            @onChange="handleChange"
            @onDestroyed="handleDestroyed" @onFocus="handleFocus" @onBlur="handleBlur"
            @customAlert="customAlert"
            @customPaste="customPaste" />
  </div>
</template>

<script lang="ts">
import { onBeforeUnmount, ref, shallowRef, onMounted, defineComponent } from "vue";
import "@wangeditor/editor/dist/css/style.css";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";

export default defineComponent({
  components: { Editor, Toolbar },
  props: {
    value: {
      type: String,
      default: "<p></p>"
    }
  },
  setup(props, context) {
    // 编辑器实例，必须用 shallowRef，重要！
    const editorRef = shallowRef();

    const toolbarConfig = {};
    const editorConfig = { placeholder: "请输入内容..." };
    // 内容 HTML
    const valueHtml = ref("");
    // 模拟 ajax 异步获取内容
    onMounted(() => {
      setTimeout(() => {
        valueHtml.value = props.value;
      }, 1000);
    });

    // 组件销毁时，也及时销毁编辑器，重要！
    onBeforeUnmount(() => {
      const editor = editorRef.value;
      if (editor == null) return;

      editor.destroy();
    });

    // 编辑器回调函数
    const handleCreated = (editor) => {
      // console.log('created', editor);
      editorRef.value = editor; // 记录 editor 实例，重要！
    };
    const handleChange = (editor) => {
      context.emit("update:value", editor.getHtml());
    };
    const handleDestroyed = (editor) => {
      // console.log('destroyed', editor)
    };
    const handleFocus = (editor) => {
      // console.log('focus', editor)
    };
    const handleBlur = (editor) => {
      // console.log('blur', editor)
    };
    const customAlert = (info, type) => {
      // alert(`【自定义提示】${type} - ${info}`)
    };
    const customPaste = (editor, event, callback) => {
      // const html = event.clipboardData.getData("text/html"); // 获取粘贴的 html
      const text = event.clipboardData.getData("text/plain"); // 获取粘贴的纯文本
      // const rtf = event.clipboardData.getData('text/rtf') // 获取 rtf 数据（如从 word wsp 复制粘贴）
      // 自定义插入内容
      editor.insertText(text);

      // 返回值（注意，vue 事件的返回值，不能用 return）
      callback(false); // 返回 false ，阻止默认粘贴行为
      // callback(true) // 返回 true ，继续默认的粘贴行为
    };

    function insertCover(cover) {
      const editor = editorRef.value;
      if (editor == null) return;
      editor.dangerouslyInsertHtml("<img src=\"" + cover + "\">");
    }

    const printHtml = () => {
      const editor = editorRef.value;
      if (editor == null) return;
      console.log(editor.getHtml());
    };

    const disable = () => {
      const editor = editorRef.value;
      if (editor == null) return;
      editor.disable();
    };

    return {
      editorRef,
      mode: "default",
      valueHtml,
      toolbarConfig,
      editorConfig,
      handleCreated,
      handleChange,
      handleDestroyed,
      handleFocus,
      handleBlur,
      customAlert,
      customPaste,
      insertCover,
      printHtml,
      disable
    };
  }
});
</script>
