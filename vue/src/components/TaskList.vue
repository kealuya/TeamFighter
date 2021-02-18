<template>
  <div>
    <div style="height:100%">
      <!--      <van-pull-refresh v-model="state.refreshing" @refresh="onRefresh">-->
      <van-list
          v-model:loading="state.loading"
          :finished="state.finished"
          finished-text="没有更多了"
          @load="onLoad"
      >
        <template v-for="item in state.list" :key="item">
          <div style="width: 100%;height: 90px;">
            <div
                @contextmenu.prevent.native="rightClick(item,$event)"
                style="margin: 15px;height: 100%;border:1px solid whitesmoke;box-shadow: 1px 1px 5px #bbb;
                    display: flex;flex-direction: row;align-items: center;justify-content: center">

              <div style="flex:3;display: flex;justify-content: center">
                <!--                <img src="../../public/目标.png" style="width: 30px;height:30px;">-->
                <div @click="showFromName(item)" class="font" :style="{background:getColorFromId(item.fromId)}">
                  {{ item.fromName }}
                </div>
              </div>
              <div style="flex: 21; height: 90%;width: 80%;font-size: 14px;display: flex;flex-direction: column; justify-content: space-between;
                  padding-left: 10px;padding-right: 10px;padding-top: 5px">
                <div style=" text-align: start;height: 28px;border-bottom: 1px solid red;
                    text-overflow:ellipsis;white-space:nowrap;overflow:hidden;">
                  Todo：{{ item.todo }}
                </div>
                <div style="text-align: end">
                  <van-tag style="margin-left: 10px" type="primary">Bug</van-tag>
                  <van-tag style="margin-left: 10px" type="success">额外</van-tag>
                  <van-tag style="margin-left: 10px" type="warning">待确认</van-tag>
                </div>
                <div style="display: flex;justify-content: space-between;align-items: center">
                  <van-slider :step="25" v-model="item.progress" active-color="#ee0a24">
                    <template #button>
                      <div class="custom-button">{{ item.progress }}</div>
                    </template>
                  </van-slider>
                  <div style="width: 40px"></div>
                  <van-rate v-model="item.stars" :count="3"/>
                </div>
              </div>

            </div>
          </div>

        </template>
      </van-list>
      <!--      </van-pull-refresh>-->

    </div>
    <div style="width: 100%">

    </div>
  </div>
</template>

<script>
import {reactive} from 'vue';
import '@vant/touch-emulator';
import {List, Rate, Cell, Col, Row, Slider, Tag, Popover} from 'vant';
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
  },
  data() {
    return {
      sort: "createTime", // createTime,stars,

    }
  },
  created() {
    remote = window.require("electron").remote
  },
  methods: {
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
      menu.append(new MenuItem({type: 'separator'}))
      menu.append(new MenuItem({
        label: '  额外  ',
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

    }
  },
  computed: {
    screenHeight: () => {
      const {ctx} = getCurrentInstance();
      return ctx.$screenHeight() + 'px'
    }
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
        for (let i = 0; i < 8; i++) {
          state.list.push({
            todo: "修改作业" + i,
            fromName: "自",
            fromId: "112345",
            stars: 2,
            state: "wait",//waiting confirmed done
            info: "",
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
          progress: 50,
        });
        state.list.push({
          todo: "春眠不觉晓，处处闻啼鸟，夜处处闻啼鸟，夜来风雨声，花落知多少",
          fromName: "李四喜",
          fromId: "112349",
          stars: 3,
          state: "wait",//waiting confirmed done
          info: "春晓，王安石",
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
<style scoped>
.font {
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

.custom-button {
  width: 26px;
  color: #fff;
  font-size: 10px;
  line-height: 18px;
  text-align: center;
  background-color: #ee0a24;
  border-radius: 100px;
}
</style>
