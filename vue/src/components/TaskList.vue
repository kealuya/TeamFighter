<template>
  <div>
    <div class="my_scroll" style="height:546px;overflow-y:scroll">
      <!-- 不需要下拉刷新，使用右键刷新菜单，下拉刷新不应该出现在pc端
           <van-pull-refresh v-model="state.refreshing" @refresh="onRefresh">-->
      <van-list
          v-model:loading="state.loading"
          :finished="state.finished"
          finished-text="干完了(*⊙~⊙)"
          @load="onLoad"
      >
        <template v-for="item in state.list" :key="item">
          <div style="width: 100%;height: 90px;">
            <!--右键菜单div，在范围内可以右键-->
            <div @contextmenu.prevent.native="rightClick(item,$event)"
                 style="margin: 15px;height: 100%;border:1px solid whitesmoke;box-shadow: rgb(222 222 222) 3px 3px 5px  ;
                    display: flex;flex-direction: row;align-items: center;justify-content: center">
              <!--任务单元-->
              <!--左侧头像、人名部分-->
              <div style=" display: flex;flex-direction: column;justify-content: center;align-items: center">
                <!--头像图片-->
                <div style="width: 45px;height:45px;" @click="showFromName(item)">
                  <img :src="getAvatar(item.avatar)" style="width: 45px;height:45px;">
                </div>
                <!--头像人名-->
                <div style="font-size: 10px">六三</div>
              </div>
              <!--右侧代办任务部分-->
              <div style="flex: 21; height: 90%;width: 80%;font-size: 14px;display: flex;flex-direction: column; justify-content: space-between;
                  padding : 5px">
                <!--任务事项-->
                <div style=" text-align: start;height: 28px;border-bottom: 1px solid #ebedf0;
                    text-overflow:ellipsis;white-space:nowrap;overflow:hidden;">
                  {{ item.todo }}
                </div>
                <!--任务属性-->
                <div style="text-align: end">
                  <van-tag style="margin-left: 10px;font-size: 10px" plain type="primary">Bug</van-tag>
                  <van-tag style="margin-left: 10px;font-size: 10px" plain type="primary">2天</van-tag>
                  <van-tag style="margin-left: 10px;font-size: 10px" plain type="primary">详情</van-tag>
                  <van-tag style="margin-left: 10px;font-size: 10px" type="primary">待确认</van-tag>
                </div>
                <!--任务情况-->
                <div style="display: flex;justify-content: space-between;align-items: center">
                  <!--任务进度-->
                  <van-slider :step="25" v-model="item.progress">
                    <template #button>
                      <div class="task-list-custom-button">{{ item.progress }}</div>
                    </template>
                  </van-slider>
                  <div style="width: 40px"></div>
                  <!--任务优先级-->
                  <van-rate v-model="item.stars" :count="3"/>
                </div>
              </div>
            </div>
          </div>
        </template>
      </van-list>
    </div>
    <!--任务录入-->
    <div style="height: 99px;width: 100%;border-top:1px solid whitesmoke">
      <!--任务录入框-->
      <van-field class="my_scroll" style="height: 70px"
                 v-model="todo"
                 rows="1"
                 label="事项"
                 type="textarea"
                 maxlength="30"
                 placeholder="待办事项"
                 show-word-limit
                 colon
                 label-width="40"
      />
      <div style="height: 29px;width: 100%">
        <van-row justify="start">
          <!--任务类型-->
          <van-col span="5">
            <van-popover placement="top" v-model:show="todo_type_popover" :actions="todo_type_actions"
                         @select="todoTypeSelect">
              <template #reference>
                <van-button style="width:60px" plain type="primary" size="mini">{{ todo_type_val }}</van-button>
              </template>
            </van-popover>
          </van-col>
          <!--任务指向人-->
          <van-col span="5">
            <van-popover placement="top" v-model:show="todo_user_show">
              <div class="my_scroll" style="height: 120px;width: 84px;overflow-y: scroll">
                <template v-for="user in todo_user_list">
                  <div style="font-size: 10px;width: 80px;height: 22px;text-align: center;"
                       @click="todoUserSelect(user)">
                    {{ user.userName }}
                  </div>
                </template>
              </div>
              <template #reference>
                <van-button style="width:60px" plain type="primary" size="mini">{{ todo_user_val }}</van-button>
              </template>
            </van-popover>
          </van-col>
          <!--任务提交按钮-->
          <van-col span="12">
            <van-row justify="end">
              <van-col>
                <van-button @click="todoSend" style="width:60px" type="primary" size="mini">提交</van-button>
              </van-col>
            </van-row>
          </van-col>
        </van-row>
      </div>
    </div>
  </div>

</template>

<script>
import {reactive, ref} from 'vue';
import '@vant/touch-emulator';
import {List, Rate, Cell, Col, Row, Slider, Tag, Popover, Field, Button} from 'vant';
//vue3.0 global组件
import {getCurrentInstance} from 'vue';
import utils from "@/utils/common";

let remote;
export default {
  name: 'TaskList',
  components: {
    [Cell.name]: Cell,
    [Row.name]: Row,
    [Col.name]: Col,
    [Rate.name]: Rate,
    [List.name]: List,
    [Slider.name]: Slider,
    [Tag.name]: Tag,
    [Popover.name]: Popover,
    [Field.name]: Field,
    [Button.name]: Button,
  },
  data() {
    return {
      sort: "createTime", // createTime,stars,
      todo: "",
      todo_type_popover: false,
      todo_type_actions: [],
      todo_type_val: "需求",
      todo_user_show: false,
      todo_user_val: "自己",
      todo_user_id_val: "",
      todo_user_list: [
        {userId: "111", userName: "智能柜1"},
        {userId: "131", userName: "智能柜2"},
        {userId: "1151", userName: "智能柜3"},
        {userId: "1161", userName: "智能柜4"},
        {userId: "1131", userName: "智能柜5"},
        {userId: "131", userName: "智能柜2"},
        {userId: "1151", userName: "智能柜3"},
        {userId: "1161", userName: "智能柜4"},
        {userId: "1131", userName: "智能柜5"},
      ]
    }
  },
  beforeUnmount() {
    // 解除快捷键绑定
    window.removeEventListener('keydown', this.shortcut)
  },
  created() {
    remote = window.require("electron").remote
    // 追加快捷键绑定
    window.addEventListener('keydown', this.shortcut)

    // to do type的初始化设定
    this.todo_type_popover = false;
    this.todo_type_actions = [
      {text: '需求', className: 'task-list-msg-type-class'},
      {text: 'Bug', className: 'task-list-msg-type-class'},
      {text: '整理', className: 'task-list-msg-type-class'},
      {text: '颠覆', className: 'task-list-msg-type-class'},
    ];

  },
  methods: {
    // 获取头像图片
    getAvatar: function (a) {
      return utils.avatars[a]
    },
    todoUserSelect: function (user) {
      this.todo_user_val = user.userName
      this.todo_user_id_val = user.userId
      this.todo_user_show = false
    },
    todoSend: function () {

      this.eventBus.emit('title_notify', {msg: '操作成功'})
      //恢复原样
      this.todo_type_val = "需求"
      this.todo_user_val = "自己"
      this.todo_user_id_val = ""

    },
    todoTypeSelect: function (action, index) {
      this.todo_type_val = action.text
    },
    shortcut: function (e) {
      if (e.ctrlKey && e.keyCode === 13) {   //用户点击了ctrl+enter触发
        this.todoSend()
      }
    },
    rightClick: function (item, $event) {
      let thatOnRefresh = this.onRefresh
      const {Menu, MenuItem} = remote
      const menu = new Menu()

      menu.append(new MenuItem({
        label: '  确认  ',
        click() {

        }
      }))
      menu.append(new MenuItem({
        label: '  完成  ',
        click() {

        }
      }))
      menu.append(new MenuItem({
        label: '  废弃  ',
        click() {

        }
      }))
      menu.append(new MenuItem({type: 'separator'}))
      menu.append(new MenuItem({
        label: '  详情  ',
        click() {

        }
      }))
      menu.append(new MenuItem({type: 'separator'}))
      menu.append(new MenuItem({
        label: '  刷新  ',
        click() {
          thatOnRefresh();
        }
      }))

      menu.append(new MenuItem({
        label: '  排序  ',
        submenu: [
          {
            label: '  时间  ',
            type: 'checkbox',
            checked: true,
            click() {
              console.log('时间排序')
            }
          },
          {
            label: '  优先级  ',
            type: 'checkbox',
            checked: false,
            click() {
              console.log('时间排序')
            }
          }
        ]
      }))

      menu.append(new MenuItem({
        label: '  显示  ',
        submenu: [
          {
            label: '  待处理  ',
            type: 'checkbox',
            checked: true,
            click() {
              console.log('时间排序')
            }
          },
          {
            label: '  全部  ',
            type: 'checkbox',
            checked: false,
            click() {
              console.log('时间排序')
            }
          }
        ]
      }))


      menu.popup({
        window: remote.getCurrentWindow(),
      })

    },
    getColorFromId: function (fromId) {
      return utils.getColorFromId(fromId)
    },
    showFromName: function () {

    },
    keyDown: function (e) {
      console.log(e)
      if (e.ctrlKey && e.keyCode === 13) {   //用户点击了ctrl+enter触发
        console.log(this.todo)
      }
    }
  },
  computed: {},
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
        for (let i = 0; i < 8; i++) {
          state.list.push({
            todo: "修改作业" + i,
            fromName: "自",
            fromId: "112345",
            stars: 2,
            state: "wait",//waiting confirmed done
            info: "",
            avatar: 2,
            progress: 30,
          });
        }

        state.list.push({
          todo: "必须今天下午之前把文档打印好2222111114444",
          fromName: "张三个",
          fromId: "112342",
          stars: 3,
          state: "wait",//waiting confirmed done
          info: "",
          avatar: 12,
          progress: 50,
        });
        state.list.push({
          todo: "春眠不觉晓，处处闻啼鸟，夜处处闻啼鸟，夜来风雨声，花落知多少",
          fromName: "李四喜",
          fromId: "112349",
          stars: 3,
          state: "wait",//waiting confirmed done
          info: "春晓，王安石",
          avatar: 6,
          progress: 90,
        });


        state.loading = false;

        if (state.list.length >= 40) {
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.task-list-font {
  width: 30px;
  height: 30px;
  top: 30px;
  line-height: 30px;
  border-radius: 30px;
  font-size: 12px;
  color: white;
  background: #6d5ac3;
  text-align: center;
}

.task-list-custom-button {
  width: 26px;
  color: #fff;
  font-size: 10px;
  line-height: 18px;
  text-align: center;
  /*TMD all_color 颜色 */
  background-color: #29b7cb;
  border-radius: 100px;
}

.task-list-msg-type-class {
  font-size: 10px;
  height: 22px;
}

/*滚动条样式设定*/
.my_scroll::-webkit-scrollbar {
  width: 3px;
  /*height: 100px;*/
}

/*滚动条样式设定*/
.my_scroll::-webkit-scrollbar-thumb {
  border-radius: 10px;
  background: #cbcbcb;
}
</style>
