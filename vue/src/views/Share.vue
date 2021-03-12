<template>
  <div class="about">
    <h1>This is an about page</h1>

    <van-button square type="primary">方形按钮</van-button>
    <van-button round type="primary">圆形按钮</van-button>

    <van-tag type="success">标签</van-tag>
  </div>
</template>

<script>

import {Button} from 'vant';
import { Tag } from 'vant';
export default {
  name: 'Share',
  components: {
    [Button.name]: Button,
    [Tag.name]: Tag,
  }
}


</script>
