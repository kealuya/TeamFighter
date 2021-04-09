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
              <img style=" width:100px;height: 100px" :src="avatars[profile.avatar]">
            </div>
          </template>
        </van-popover>
      </van-col>
    </van-row>
    <div style="height: 10px"></div>
    <van-row>
      <van-field v-model="profile.id" label="文本" placeholder="请输入用户名"/>
    </van-row>

    <van-field v-model="profile.id" label="文本" placeholder="请输入用户名"/>

    <van-row type="flex" justify="end">
      <van-col span="6">
        <van-button @click="logout" round type="primary">退出登录</van-button>
      </van-col>
    </van-row>

  </div>

</template>

<script>
import {getCurrentInstance} from "vue";
import {Cell, Col, Grid, GridItem, Popover, Row, Field, Tag, Button} from "vant";
import {ref} from 'vue';
import utils from "@/utils/common";

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

      profile: {
        id: "234523",
        avatar: "4",
      },
      avatars: utils.avatars
    }
  },
  methods: {
    changeAvatar: function (i) {
      this.showPopover = false;
      this.profile.avatar = i
    },
    logout: function () {
      utils.ipcAccess("store", {
        method: [utils.storeMethod.del],
        payload: [utils.storeKey.userInfo]
      }).then(() => {
        localStorage.removeItem("userInfo")
        utils.ipcAccess("operate", {operate: "logout"})
      })
    }
  },

  computed: {},
  setup() {
    const showPopover = ref(false);
    return {showPopover};
  },
}
</script>

<style scoped>

</style>