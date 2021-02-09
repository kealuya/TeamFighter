<template>
  <div class="home">

    <img alt="Vue logo" src="../../public/flower_icon.png">
    <HelloWorld msg="2222 to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'

const electron = window.require('electron')
export default {
  name: 'Test',
  components: {
    HelloWorld,
  },
  mounted() {

    electron.ipcRenderer.send("get_version")
    electron.ipcRenderer.on('get_version_reply', (event, arg) => {
      const message = `异步消息回复: ${arg}`
      console.log(message);
    })
  }
}
</script>
