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
          <div>【张三】需要【你】交作业</div>
          <div>2021-02-14 12:31:12</div>
        </van-step>
      </template>


    </van-steps>


  </van-list>
</template>

<script>
import {Icon, Step, Steps, List, Tabs} from 'vant';
import TaskList from "@/components/TaskList";
import {reactive} from "vue";

export default {
  name: "History",
  components: {
    TaskList,
    [Step.name]: Step,
    [Steps.name]: Steps,
    [List.name]: List,
  },

  setup() {

    const state = reactive({
      list: [],
      loading: false,
      finished: false,
      refreshing: false,
    });

    const onLoad = () => {
      setTimeout(() => {
        if (state.refreshing) {
          state.list = [];
          state.refreshing = false;
        }
        // 数据部分
        for (let i = 0; i < 20; i++) {
          state.list.push({
            thing: "修改作业" + i,
            date: "2020-01-12 12:33:43",
          });
        }


        state.loading = false;

        if (state.list.length >= 140) {
          state.finished = true;
        }
      }, 1000);
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
    };
  },

}
</script>

<style scoped>

</style>