<template>
  <div class="home">
    <van-tabs v-model:active="tab_active">
      <van-tab>
        <template #title>
          <van-icon name="todo-list-o"/>
          一览
        </template>
        <TaskList/>
      </van-tab>

      <van-tab>
        <template #title>
          <van-icon name="todo-list-o"/>
          履历
        </template>
        <History/>
      </van-tab>

      <van-tab>
        <template #title>
          <van-icon name="todo-list-o"/>
          统计
        </template>
        在做了。。。😂
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>

import {Tab, Tabs, Icon} from 'vant';
import TaskList from "@/components/TaskList";
import History from "@/components/History";

export default {
  name: 'Task',
  components: {
    History,
    TaskList,
    [Tab.name]: Tab,
    [Tabs.name]: Tabs,
    [Icon.name]: Icon,
  },
  data() {
    return {
      tab_active: 0
    }
  }
}
</script>
