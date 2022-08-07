/*
 * @Author: coderzhaolu && izhaicy@163.com
 * @Date: 2022-07-29 19:46:40
 * @LastEditors: coderzhaolu && izhaicy@163.com
 * @LastEditTime: 2022-08-07 15:28:40
 * @FilePath: /xanaduCms/pinkmoe_index/src/store/modules/socket/index.ts
 * @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 * QQ:2419857357;支付宝:13135986153
 * Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
 */
import { defineStore } from 'pinia';
import { useUserStore } from '../user';
import { getToken, isLogin } from '/@/utils/auth';

export interface SocketState {
  socket?: any;
  lockReconnect: boolean;
  tt: any;
  heartCheck: any;
}

export const useSocketStore = defineStore({
  id: 'socket',
  state: (): SocketState => ({
    // 连接状态
    socket: WebSocket,
    lockReconnect: false,
    tt: null,
    heartCheck: {
      timeout: 3000,
      timeoutObj: null,
      serverTimeoutObj: null,
    },
  }),
  actions: {
    start() {
      const auth = useUserStore();
      this.heartCheck.timeoutObj && clearTimeout(this.heartCheck.timeoutObj);
      this.heartCheck.serverTimeoutObj && clearTimeout(this.heartCheck.serverTimeoutObj);
      this.heartCheck.timeoutObj = setTimeout(() => {
        //这里发送一个心跳，后端收到后，返回一个心跳消息，
        this.socket.send(
          JSON.stringify({
            type: 'heartbeat',
            chatMsg: {
              userId: auth.userInfo?.uuid,
            },
          }),
        );
        // this.heartCheck.serverTimeoutObj = setTimeout(() => {
        //   console.log(111);
        //   console.log(this.socket);
        //   this.socket.close();
        //   // createWebSocket();
        // }, this.heartCheck.timeout);
      }, this.heartCheck.timeout);
    },
    closeSocket() {
      this.socket.close();
    },
    reconnect() {
      if (this.lockReconnect) {
        return;
      }
      this.lockReconnect = true;
      //没连接上会一直重连，设置延迟避免请求过多
      this.tt && clearTimeout(this.tt);
      this.tt = setTimeout(() => {
        this.createWebSocket();
        this.lockReconnect = false;
      }, 4000);
    },
    createWebSocket() {
      if (isLogin()) {
        try {
          this.socket = new WebSocket('/ws?token=' + getToken());
          this.init();
        } catch (e) {
          this.reconnect();
        }
      }
    },
    init() {
      this.socket.onclose = () => {
        this.reconnect();
      };
      this.socket.onerror = (e) => {
        this.reconnect();
      };
      this.socket.onopen = () => {
        //心跳检测重置
        this.start();
        // 自定义全局监听事件
        window.dispatchEvent(
          new CustomEvent('onopenWS', {
            detail: {
              data: this.socket,
            },
          }),
        );
      };
      this.socket.onmessage = (event) => {
        this.start();
        window.dispatchEvent(
          new CustomEvent('onmessageWS', {
            detail: {
              data: JSON.parse(event.data),
            },
          }),
        );
      };
    },
  },
});
