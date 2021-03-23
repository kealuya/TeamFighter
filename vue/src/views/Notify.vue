<template>
  <div style="width: 500px;height: 120px; text-align: center;background-color: #6d5ac3">
    <van-row justify="space-between">
      <van-col span="6">span: 6</van-col>
      <van-col span="6">span: 6</van-col>
      <van-col span="6">
        <van-row justify="end">
          <van-col @click="close" span="6">X</van-col>
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
      content: ""
    }
  },
  mounted() {
    setTimeout(() => {
      this.close()
    }, 4000)
  },
  created() {

    let ipcRenderer = window.require("electron").ipcRenderer
    ipcRenderer.on('param', (event, message) => { // 监听父页面定义的端口
      this.content = message.content
      console.log(message.content)
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
