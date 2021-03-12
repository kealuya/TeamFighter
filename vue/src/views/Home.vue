<template>
  <div style="height: 100vh; ">
    <!--    标题栏-->
    <van-row style="height: 39px;border-bottom: 1px solid whitesmoke ;-webkit-app-region: drag;-webkit-user-select:none"
             align="center">
      <van-col span="8"></van-col>
      <van-col span="8">TeamFighter</van-col>
      <van-col span="8">
        <van-row justify="end" align="center">
          <van-col span="3">X</van-col>
          <van-col span="3">X</van-col>
          <van-col span="3">X</van-col>
        </van-row>
      </van-col>
    </van-row>

    <van-row style="">
      <!--       个人头像-->
      <van-col span="4" style="border-right: 1px solid whitesmoke ;">
        <!--      空白-->
        <van-row style="height: 10px;"></van-row>
        <van-row justify="center" align="center">
          <van-col>
            <img style="width: 60px;height: 60px;" src="../../public/profile/default5.png"/>
          </van-col>
        </van-row>
        <!--      空白-->
        <van-row style="height: 10px;"></van-row>
        <van-sidebar v-model="active" @change="onChange">
          <van-sidebar-item title="任务" badge="5"/>
          <van-sidebar-item title="扩展"/>
          <van-sidebar-item title="分享"/>
          <van-sidebar-item title="设置"/>
          <van-sidebar-item title="管理"/>
        </van-sidebar>
      </van-col>
      <!--      主工作区-->
      <van-col span="20">
        <router-view :style="{height:'690px',width: '100%',overflowY:'scroll'}"/>
      </van-col>
    </van-row>
    <van-row align="center"
             style="background-color: #cbcaca;height:19px;font-size: 12px;border-top: 1px solid whitesmoke ;">
      <van-col span="20">状态栏</van-col>
    </van-row>
  </div>
</template>
<script>

import {Col, Row, Sidebar, SidebarItem, Notify} from 'vant';
import {getCurrentInstance, ref} from 'vue';

export default {
  name: 'Home',
  components: {
    [Col.name]: Col,
    [Row.name]: Row,
    [Sidebar.name]: Sidebar,
    [SidebarItem.name]: SidebarItem,
    [Notify.name]: Notify,
  },
  data() {
    return {
      active: 0
    }
  },
  created() {
    this.$router.push({path: "task"})
    this.eventBus.on('title_notify', this.onNotifyMsg)
  },
  beforeUnmount() {
    this.eventBus.off('title_notify', this.onNotifyMsg)
  },
  methods: {
    onChange: function (index) {
      switch (index) {
        case 0:
          this.$router.push({path: "task"})
          break;
        case 1:
          this.$router.push({path: "expand"})
          break;
        case 2:
          this.$router.push({path: "share"})
          break
        case 3:
          this.$router.push({path: "setting"})
          break
        case 4:
          this.$router.push({path: "control"})
          break
      }
    },
    onNotifyMsg: function (e) {
      Notify({type: 'primary', message: e.msg});
    }
  },

}
</script>
<style>
/* TMD electron 禁止选择页面文字   */
body {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}


#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}


</style>
