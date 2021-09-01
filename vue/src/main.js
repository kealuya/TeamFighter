import {createApp} from 'vue'
import {createStore} from 'vuex'
import mitt from 'mitt'
import App from './App.vue'
import router from './router'
import store from './store'

let app = createApp(App);
app.use(store)

//vue3.0 global组件 // app.config.globalProperties 必须写到app.use(router).mount('#app'); 这句话的上面才起作用，TMD
/*
      const {ctx} = getCurrentInstance();
      return ctx.$screenHeightOther() + 'px'
      上面的获取global变量的方式是错的，网上写这个的人都是sb，还是要看官网啊
        app.config.globalProperties.foo = 'bar'
        app.component('child-component', {
          mounted() {
            console.log(this.foo) // 'bar'
          }
        })
*/
app.config.globalProperties.taskListHeight = () => {
    // 30: 标题栏高度， 44：tabs的高度，20：状态栏，剩下的是主工作区的高度 fixme
    return 750 - 40 - 20; // width:380
}

app.config.globalProperties.taskListWidth = () => {
    // 30: 标题栏高度， 44：tabs的高度，20：状态栏，剩下的是主工作区的高度 fixme
    return 400 - 65; // width:380
}

app.config.globalProperties.eventBus = mitt();


app.use(router).mount('#app');