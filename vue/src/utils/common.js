// 用import获取头像图片，可以作为对象进行处理，很方便
// <img :src="getAvatar(item.avatar)" style="width: 45px;height:45px;"> 这样引入
import headImg0 from '../../public/profile/default0.png'
import headImg1 from '../../public/profile/default1.png'
import headImg2 from '../../public/profile/default2.png'
import headImg3 from '../../public/profile/default3.png'
import headImg4 from '../../public/profile/default4.png'
import headImg5 from '../../public/profile/default5.png'
import headImg6 from '../../public/profile/default6.png'
import headImg7 from '../../public/profile/default7.png'
import headImg8 from '../../public/profile/default8.png'
import headImg9 from '../../public/profile/default9.png'
import headImg10 from '../../public/profile/default10.png'
import headImg11 from '../../public/profile/default11.png'
import headImg12 from '../../public/profile/default12.png'
import headImg13 from '../../public/profile/default13.png'
import headImg14 from '../../public/profile/default14.png'
//================================================================//
import pic_logo from '../../public/logo.png'
import pic_teamFighter from '../../public/内卷系统.png'
import picMsgMainLogo from '../../public/picMsgMainLogo.png'
import picMission from '../../public/铃铛.png'

const electron = window.require('electron')

let utils = {
    // nodejs ipcRenderer 交互，通过promise进行封装
    ipcAccess: function (channel, arg) {
        let d = new Date()
        arg.onceId = d.getTime()
        return new Promise((resolve, reject) => {
            electron.ipcRenderer.send(channel, arg)
            electron.ipcRenderer.once(channel + '_reply_' + arg.onceId, (event, arg) => {
                resolve(arg)
            })
        })
    },
    // 头像图片对象数组
    avatars: [
        headImg0,
        headImg1,
        headImg2,
        headImg3,
        headImg4,
        headImg5,
        headImg6,
        headImg7,
        headImg8,
        headImg9,
        headImg10,
        headImg11,
        headImg12,
        headImg13,
        headImg14,
    ],
    //通过id自动生成对应颜色
    getColorFromId: (fromId) => {
        let color = "#FFF"
        switch (fromId.substr(fromId.length - 1, 1)) {
            case '0':
                color = "#ed5a65"
                break;
            case '1':
                color = "#b55176"
                break;
            case '2':
                color = "#8b2671"
                break;
            case '3':
                color = "#126bae"
                break;
            case '4':
                color = "#44889b"
                break;
            case '5':
                color = "#d6cb28"
                break;
            case '6':
                color = "#20894d"
                break;
            case '7':
                color = "#fcd337"
                break;
            case '8':
                color = "#ff9900"
                break;
            case '9':
                color = "#ac1f18"
                break;
            default:
                color = "#a55353"
                break;
        }
        return color
    },
    //公司logo
    picLogo: pic_logo,
    //系统title的logo图片
    picTeamFighter: pic_teamFighter,
    picMsgMainLogo: picMsgMainLogo,
    picMission: picMission,
    // kv存储对应key值，防止key值写错
    /*🎉*/storeMethod: {
        get: "get",
        put: "put",
        del: "del",
    },
    /*🎉*/storeKey: {
        userInfo: "userInfo",
    },
    //基础请求路径，后续通过配置进行选择
    /*🎉*/httpBaseUrl: "http://127.0.0.1:8000/v1/",

    //日期格式化
    dateFtt: function (fmt, date) {
        //author: meizz
        let o = {
            'M+': date.getMonth() + 1,                 //月份
            'd+': date.getDate(),                    //日
            'h+': date.getHours(),                   //小时
            'm+': date.getMinutes(),                 //分
            's+': date.getSeconds(),                 //秒
            'q+': Math.floor((date.getMonth() + 3) / 3), //季度
            'S': date.getMilliseconds()             //毫秒
        };
        if (/(y+)/.test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
        }
        for (let k in o) {
            if (new RegExp('(' + k + ')').test(fmt)) {
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? (o[k]) : (('00' + o[k]).substr(('' + o[k]).length)))
            }
        }
        return fmt
    },

}

export default utils;