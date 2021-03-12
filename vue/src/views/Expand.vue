<template>
  <div class=" ">
    <van-grid :border="false" :column-num="3">
      <van-grid-item @click="weixinSend">
        <img :style="{width: sw}" src="../../public/日历.png"/>
        <div style="font-size:12px">企业微信推送</div>
      </van-grid-item>
      <van-grid-item>
        <img :style="{width: sw}" src="../../public/flower_icon.png"/>
      </van-grid-item>
      <van-grid-item>
        <img :style="{width: sw}" src="../../public/flower_icon.png"/>
      </van-grid-item>
      <van-grid-item>
        <img :style="{width: sw}" src="../../public/flower_icon.png"/>
      </van-grid-item>
      <van-grid-item>
        <img :style="{width: sw}" src="../../public/flower_icon.png"/>
      </van-grid-item>
      <van-grid-item>
        <img :style="{width: sw}" src="../../public/flower_icon.png"/>
      </van-grid-item>
    </van-grid>
  </div>
</template>

<script>

import {Button} from 'vant';
import {Tag} from 'vant';

import {Image as VanImage} from 'vant';
import {Grid, GridItem} from 'vant';

const electron = window.require('electron')
export default {
  name: 'Expand',
  components: {
    [Grid.name]: Grid,
    [GridItem.name]: GridItem,
    [VanImage.name]: VanImage,
  },
  data() {
    return {
      sw: ""
    }
  },
  created() {
    // console.log(this.taskListWidth())
    this.sw = 40 + 'px'
  },
  methods: {
    weixinSend: function () {
      //主进程交互，传送数据
      let p = {
        target: "weixinSend",
        payload: {
          userid: "123123123",
          content: "吃饭了"
        }
      };
      electron.ipcRenderer.send("expand", JSON.stringify(p))
      electron.ipcRenderer.on('expand_reply', (event, arg) => {
        const message = `异步消息回复: ${arg}`
        console.log(message);
      })
    }
  }
}


</script>
