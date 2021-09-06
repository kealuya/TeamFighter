<template>
  <div style="padding: 30px;text-align: center">
    <div style="text-align: left">
      <img :src="logo" style="width: 70px;height:25px;"/>
    </div>
    <div>
      <img :src="teamFighter" style="width: 180px;height:85px;"/>
    </div>
    <div style="width: 200px">

    </div>

    <div style="display: flex;width: 100%;justify-content: center">
      <van-field v-model="userid" type="number" label="工号" placeholder="请输入工号"/>
    </div>
    <div style="display: flex;width: 100%;justify-content: center">
      <van-field v-model="password" type="password" label="密码" placeholder="请输入密码"/>
    </div>
    <div style="display: flex;width: 100%;justify-content: center;margin-top: 20px">
      <van-radio-group v-model="channel" direction="horizontal">
        <van-radio name="htjy">浩天教育</van-radio>
        <van-radio name="test">测试组</van-radio>
      </van-radio-group>
    </div>


    <div style="display: flex;width: 100%;justify-content: center;
    height: 40px;
    font-size: 14px;
    color: red;
     white-space:normal;
     word-break:break-all;
     word-wrap:break-word; ">
      {{ msg }}
    </div>

    <van-button @click="login" style="width: 200px" square type="primary">登录</van-button>

    <div style="width: 200px;margin: 10px;"/>
    <van-button @click="logout" style="width: 200px" square type="primary">退出</van-button>
    <div style="width: 200px;margin: 45px;"/>
    <div style="text-align: right">
      version 0.2.11
    </div>
  </div>
</template>

<script>

import {Button} from 'vant';
import {Tag} from 'vant';
import {Field} from 'vant';
import {Col, Row, Picker, RadioGroup, Radio} from 'vant';
import {useStore} from 'vuex';
import utils from "@/utils/common";

export default {
  name: 'Login',
  components: {
    [Button.name]: Button,
    [Tag.name]: Tag,
    [Field.name]: Field,
    [Col.name]: Col,
    [Row.name]: Row,
    [RadioGroup.name]: RadioGroup,
    [Radio.name]: Radio,

  },
  data() {
    return {
      teamFighter: utils.picTeamFighter,
      logo: utils.picLogo,
      userid: "",
      password: "",
      msg: "",
      channel: "htjy",
    }
  },

  async beforeCreate() {

    let ui = this.getUserInfo()
    if (JSON.stringify(ui) !== "{}") {
      utils.ipcAccess("operate", {operate: "login"})
    }

    // await utils.ipcAccess("store", {
    //   method: [utils.storeMethod.get],
    //   payload: [utils.storeKey.userInfo]
    // }).then(result => {
    //   if (result != null) {
    //     utils.ipcAccess("operate", {operate: "login"})
    //   }
    // })
  },
  methods: {
    login: function () {
      this.msg = ""
      if (this.userid.trim() === "" || this.password.trim() === "") {
        this.msg = "用户名或密码不正确"
        return
      }

      utils.ipcAccess("http", {
        url: utils.httpBaseUrl + "u/login",
        method: "post",
        parameter: {userid: this.userid, password: this.password,}
      }).then(ro => {
        if (!ro.success) {
          this.msg = ro.msg.replace("\r").replace("\n")
        } else {
          if (JSON.stringify(ro.data) === "{}") {
            this.msg = "工号或密码输入不正确"
          } else {
            // ⛹🏻‍暂时采用vuex处理登录问题
            this.setUserInfo(ro.data)
            utils.ipcAccess("operate", {operate: "login"})
            // utils.ipcAccess("store", {
            //   method: [utils.storeMethod.put],
            //   payload: [utils.storeKey.userInfo, ro.data]
            // }).then(() => {
            //   // localStorage.setItem("userInfo", JSON.stringify(ro.data))//localstorage里只能存string
            //
            //   utils.ipcAccess("operate", {operate: "login"})
            // })
          }
        }
      })
    },
    logout: function () {
      console.log("quit")
      utils.ipcAccess("operate", {
        operate: "quit",
        parameter: ""
      })
    }
  },
  setup() {
    const store = useStore()
    const getUserInfo = () => store.state.sys.userInfo
    const setUserInfo = (userInfo) => {
      store.commit("setUserInfo", userInfo)
    }
    return {
      setUserInfo, getUserInfo
    }
  }
}


</script>
