<template>
  <div class="about">
    <van-button @click="testNotify" square type="primary">测试消息通知</van-button>
    <div style="height: 30px"></div>
    <van-button round type="primary">圆形按钮</van-button>
    <van-cell is-link @click="popupTaskDetail">展示弹出层</van-cell>
    <van-popup v-model:show="show_task_detail">内容</van-popup>

    <van-button @click="testSetStore" round type="primary">setStore</van-button>
    <van-button @click="testGetStore" round type="primary">getStore</van-button>
  </div>
</template>

<script>

import {Button, Cell} from 'vant';
import {Tag, Popup} from 'vant';
import {useStore} from 'vuex';


const electron = window.require('electron')
export default {
  name: 'Control',
  components: {
    [Button.name]: Button,
    [Cell.name]: Cell,
    [Tag.name]: Tag,
    [Popup.name]: Popup,
  },
  data() {
    return {
      show_task_detail: false
    }
  },
  methods: {
    testNotify: function () {

      electron.ipcRenderer.send("notify", {content: "三级分类经二路", title: "2222323"})
      electron.ipcRenderer.once('notify_reply', (event, arg) => {
        const message = `异步消息回复: ${arg}`
      })
    },
    popupTaskDetail: function () {
      console.log(111)
      this.show_task_detail = true
    },


  },
  setup() {
    const store = useStore()
    const testGetStore = () => {
      console.log("testGetStore")
      console.log(store.state.sys.userInfo)
    }

    const testSetStore = () => {
      console.log("testSetStore")
      store.commit("setUserInfo", {"ttt": 2222})
    }

    return {
      testGetStore,
      testSetStore
    }
  }

}


</script>
