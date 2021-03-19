<template>
  <div class="about">
    <van-button @click="testNotify" square type="primary">测试消息通知</van-button>
    <div style="height: 30px"></div>
    <van-button round type="primary">圆形按钮</van-button>
  </div>
</template>

<script>

import {Button} from 'vant';
import {Tag} from 'vant';

const electron = window.require('electron')
export default {
  name: 'Control',
  components: {
    [Button.name]: Button,
    [Tag.name]: Tag,
  },
  methods: {
    testNotify: function () {

      electron.ipcRenderer.send("notify", {msg: "ddddddd"})
      electron.ipcRenderer.on('notify_reply', (event, arg) => {
        const message = `异步消息回复: ${arg}`
      })
    }
  }

}


</script>
