<template>
  <div style="height:100%;">
    <van-row justify="center">
      <van-col>
        <van-popover v-model:show="showPopover">
          <van-grid
              square
              clickable
              :border="false"
              column-num="3"
              style="width: 240px;"
          >
            <van-grid-item
                v-for="(item,i) in avatars"
                :key="i"
                text="选项"
                icon="photo-o"
                @click="changeAvatar(i)"
            >
              <img style=" width:80px;height: 80px" :src="item">
            </van-grid-item>
          </van-grid>
          <template #reference>
            <div>
              <img style=" width:100px;height: 100px" :src="avatars[userInfo.avatar]">
            </div>
          </template>
        </van-popover>
      </van-col>
    </van-row>
    <div style="height: 10px"></div>
    <van-field readonly v-model="level" label="职务"/>
    <van-field readonly v-model="userInfo.userid" label="用户id"/>
    <van-field readonly v-model="userInfo.username" label="用户名称"/>

    <van-row style="margin-top: 40px" type="flex" justify="end">
      <van-col span="10">
        <van-button @click="logout" round type="primary">退出登录</van-button>
      </van-col>
    </van-row>

  </div>

</template>

<script>
import {getCurrentInstance, reactive} from "vue";
import {Cell, Col, Grid, GridItem, Popover, Row, Field, Tag, Button} from "vant";
import {ref} from 'vue';
import utils from "@/utils/common";
import {useStore} from 'vuex';

export default {
  name: "Setting",
  components: {
    [Cell.name]: Cell,
    [Row.name]: Row,
    [Button.name]: Button,
    [Col.name]: Col,
    [Grid.name]: Grid,
    [GridItem.name]: GridItem,
    [Popover.name]: Popover,
    [Field.name]: Field,
  },
  data() {
    return {
      userInfo: this.getUserInfo(),
      avatars: utils.avatars,
      level: ""
    }
  },
  methods: {
    changeAvatar: function (i) {
      this.showPopover = false;
      this.userInfo.avatar = i
      this.setUserInfo(this.userInfo)
      let param = {
        "avatar": this.userInfo.avatar,
        "userid": this.userInfo.userid
      }
      utils.ipcAccess("http", {
        url: utils.httpBaseUrl + "u/set_user_info",
        method: "post",
        parameter: param
      }).then(ro => {
        //logging
        utils.ipcAccess("logging", {
          logType: "info",
          logContent: "u/set_user_info" + "::" + JSON.stringify(param) + "::" + JSON.stringify(ro)
        })
      })
    },
    logout: function () {
      this.setUserInfo({})
      utils.ipcAccess("operate", {operate: "logout"})
      // utils.ipcAccess("store", {
      //   method: [utils.storeMethod.del],
      //   payload: [utils.storeKey.userInfo]
      // }).then(() => {
      //   this.setUserInfo({})
      //   utils.ipcAccess("operate", {operate: "logout"})
      // })
    }
  },
  mounted() {

    let lv = "";
    switch (this.userInfo.level) {
      case "top":
        lv = "服务员"
        break
      case "normal":
        lv = "打工人"
        break
      default:
        lv = "打工人的人"
        break
    }
    this.level = lv
  },

  computed: {},
  setup() {
    const store = useStore()
    const getUserInfo = () => store.state.sys.userInfo
    const setUserInfo = (userInfo) => {
      store.commit("setUserInfo", userInfo)
    }
    const showPopover = ref(false);
    return {showPopover, getUserInfo, setUserInfo};
  },
}
</script>

<style scoped>

</style>