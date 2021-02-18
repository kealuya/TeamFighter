<template>
  <div style="height:700px;overflowY:scroll">
    2
  </div>

</template>

<script>
import {getCurrentInstance} from "vue";

export default {
  name: "Setting",

  computed: {
    screenHeight: () => {
      const {ctx} = getCurrentInstance();
      return ctx.$screenHeightOther() + 'px'
    }
  },
}
</script>

<style scoped>

</style>