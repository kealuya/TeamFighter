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
      <van-field v-model="username" type="number" label="工号" placeholder="请输入工号"/>
    </div>
    <div style="display: flex;width: 100%;justify-content: center">
      <van-field v-model="password" type="password" label="密码" placeholder="请输入密码"/>
    </div>

    <div style="width: 200px;margin: 30px;"/>

    <van-button style="width: 200px" square type="primary">登录</van-button>

    <div style="width: 200px;margin: 10px;"/>
    <van-button style="width: 200px" square type="primary">退出</van-button>
    <div style="width: 200px;margin: 45px;"/>
    <div style="text-align: right">
      version 0.2
    </div>
  </div>
</template>

<script>

import {Button} from 'vant';
import {Tag} from 'vant';
import {Field} from 'vant';
import {Col, Row} from 'vant';

import utils from "@/utils/common";

export default {
  name: 'Login',
  components: {
    [Button.name]: Button,
    [Tag.name]: Tag,
    [Field.name]: Field,
    [Col.name]: Col,
    [Row.name]: Row,
  },
  data() {
    return {
      teamFighter: utils.picTeamFighter,
      logo: utils.picLogo,
      username: "",
      password: "",
    }
  }
}


</script>
