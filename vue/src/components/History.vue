<template>

  <van-list
      v-model:loading="state.loading"
      :finished="state.finished"
      finished-text="没有更多了"
      @load="onLoad"
  >


    <van-steps direction="vertical" :active="-1">


      <template v-for="item in state.list" :key="item">
        <van-step style=" font-size: 14px;text-align: start">
          <div>【{{ item.fromName }}】需要【{{item.toName}}】交作业</div>
          <div>{{ item.todoType }}</div>
        </van-step>
      </template>


    </van-steps>


  </van-list>
</template>

<script>
import {Icon, Step, Steps, List, Tabs} from 'vant';
import TaskList from "@/components/TaskList";
import {reactive} from "vue";
import utils from "@/utils/common";

export default {
  name: "History",
  components: {
    TaskList,
    [Step.name]: Step,
    [Steps.name]: Steps,
    [List.name]: List,
  },
  data() {
    return {}
  },
  methods: {},
  created() {

  },
  setup() {

    const state = reactive({
      list: [],
      loading: false,
      finished: false,
      refreshing: false,
    });

    //当前页初始化
    let currentPage = 0;
    const onLoad = () => {
      /*  setTimeout(() => {

        }, 1000);*/
      if (state.refreshing) {
        state.list = [];
        state.refreshing = false;
        currentPage = 1;
      } else {
        currentPage += 1;
      }
      // 数据部分
      /*for (let i = 0; i < 20; i++) {
        state.list.push({
          thing: "修改作业" + i,
          date: "2020-01-12 12:33:43",
        });
      }*/
      let ui = JSON.parse(localStorage.getItem("userInfo"))//localStorage取得对象需要转换成json使用
      utils.ipcAccess("http", {
        url: utils.httpBaseUrl + "t/get_record_list",
        method: "post",
        parameter: {userid: ui.userid, page: currentPage, query: ""}
      }).then(response => {
        if (response.success){
          let count = response.data.count;
          let tasks = response.data.tasks;
          // 数据部分
          state.list = state.list.concat(tasks)
          state.loading = false;
          if (state.list.length === count) {
            state.finished = true;
          }
        }else {
          //logging
          utils.ipcAccess("logging", {logType: "error", logContent: response})
          return
        }
      })
    };

    const onRefresh = () => {
      // 清空列表数据
      state.list = [];
      state.finished = false;
      state.refreshing = true;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      state.loading = true;
      onLoad();
    };

    return {
      state,
      onLoad,
      onRefresh,
      currentPage
    };
  },

}
</script>

<style scoped>

</style>
