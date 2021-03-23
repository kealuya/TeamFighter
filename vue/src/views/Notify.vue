<template>
  <div style="width: 450px;height: 120px; text-align: center;background-color: #ffffff">
    <van-row style="height: 15px" justify="space-between">
      <van-col span="12"></van-col>
      <van-col span="6">
        <van-row justify="end">
          <van-col @click="close" span="6">X</van-col>
        </van-row>
      </van-col>
    </van-row>
    <van-row style="height: 90px" justify="star">
      <van-col span="6">
        <img style="height: 80px;width: 80px" :src="picMsgMainLogo">
      </van-col>
      <van-col span="18">
        <van-row style=" " justify="center">
          <div style="font-size: 20px;font-weight: bold"> {{ title }}</div>
        </van-row>
        <van-row style=" " justify="center">
          <div style="font-size: 18px"> {{ content }}</div>
        </van-row>
      </van-col>
    </van-row>
  </div>
</template>

<script>

import {Button} from 'vant';
import {Tag} from 'vant';
import {Field} from 'vant';
import {Col, Row} from 'vant';

import utils from "@/utils/common";

export default {
  name: 'Notify',
  components: {
    [Button.name]: Button,
    [Tag.name]: Tag,
    [Field.name]: Field,
    [Col.name]: Col,
    [Row.name]: Row,
  },
  data() {
    return {
      title: "",
      content: "",
      picMsgMainLogo: utils.picMsgMainLogo
    }
  },
  mounted() {
    setTimeout(() => {
      this.close()
    }, 4000)
  },
  beforeCreate() {
    let ipcRenderer = window.require("electron").ipcRenderer
    ipcRenderer.on('param', (event, message) => { // 监听父页面定义的端口
      this.title = message.title
      this.content = message.content
    });
  },
  methods: {
    close: function () {
      let remote = window.require("electron").remote
      remote.getCurrentWindow().close()
    }
  }
}


</script>
