import Vuex from 'vuex'
import createPersistedState from "vuex-persistedstate"
import application from './module/application'
import system from './module/system'


export default new Vuex.Store({
    // plugins: [createPersistedState({key: "szht_ct", storage: window.sessionStorage})],
    plugins: [createPersistedState({
        key: "szht_store",
        storage: {
            getItem: (key) =>  window.localStorage.getItem(key),
            setItem: (key, value) => {
                value = value + "";//此处如果是数字，就会发生错误
                window.localStorage.setItem(key, value)
            },
            removeItem: key => window.localStorage.removeItem(key)
        }
    })],
    modules: {
        // app: application,//应用管理
        sys: system,//系统配置
    }
})
