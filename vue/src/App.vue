<template>
  <div  >
    <router-view></router-view>
  </div>
</template>
<script>


export default {
  name: 'App',

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


</style>
