const {app, ipcMain, BrowserView, BrowserWindow, screen} = require('electron');
const axios = require('axios');
const log = require('./common/log_config') // 引入主进程逻辑
const utils = require('./common/tf_utils')
const path = require('path')
const level = require('level')
/*
type httpResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
TMD 统一前后台数据结构
 */
let hook = function () {
    // 所有方法必须异步返回，不能同步，不然可能发生主线程阻塞
    console.log('hook_handler init')

    /*
     getVersion 获取程序版本 来自package.json
     */
    ipcMain.on("get_version", ipc_get_version)

    function ipc_get_version(event, arg) {
        event.sender.send('get_version_reply', app.getVersion());
    }

    /*
     log记录
     */
    ipcMain.on("logging", ipc_logging)

    function ipc_logging(event, arg) {
        let logType = arg.logType
        let logContent = JSON.stringify(arg.logContent)

        log[logType](logContent);
    }

    /*
     扩展功能
     */
    ipcMain.on("expand", ipc_expand)

    function ipc_expand(event, arg) {
        //todo 追加log
        const win = new BrowserWindow({
            width: 400, height: 750,
            title: "浩天差旅统计分析",
            resizable: false,
            // frame:false,
        })

        const view = new BrowserView()
        win.setBrowserView(view)
        view.setBounds({x: 0, y: 0, width: 400, height: 750})
        view.webContents.loadURL('http://pioneer.szhtkj.com.cn:8084/analyse')


        event.sender.send('expand_reply', app.getVersion());
    }


    /*
     消息通知
     */
    ipcMain.on("notify", ipc_notify)

    function ipc_notify(event, arg) {
        // 解决electron console里提示安全warning的问题
        process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';
        const {width, height} = screen.getPrimaryDisplay().workAreaSize

        const win = new BrowserWindow({
            width: 450, height: 120,
            title: "TeamFighter -- 通知",
            resizable: false,
            x: width - 450,
            y: height - 120,
            alwaysOnTop: true,
            frame: false,
            webPreferences: {
                contextIsolation: false,
                enableRemoteModule: true,//来打开remote模块，使得渲染进程中可以调用主进程的方法
                nodeIntegration: true,//渲染进程也可以使用node模块，electron引入vue，vue所在窗口与主线程交互必须要true
            },
        })


        const notifyURL = process.env.NODE_ENV === 'development'
            ? `http://localhost:8080/#/notify`
            : `${__dirname}/vue/dist/index.html/#/notify`
        // AND LOAD THE INDEX.HTML OF THE APP.
        if (process.env.NODE_ENV === 'development') {
            win.loadURL(notifyURL)
        } else {
            win.loadFile(notifyURL)
        }
        win.webContents.on('did-finish-load', function () {
            win.webContents.send('param', arg);
            event.sender.send('notify_reply', {"msg": "通知发送成功", "success": true});
        });
        // win.webContents.openDevTools();

    }

    /*
     * 存储操作
     */
    ipcMain.on("store", ipc_store)

    function ipc_store(event, arg) {
        let method = arg.method;
        let payload = arg.payload;
        store[method](...payload).then(val => {
            event.sender.send('store_reply', val);
        }).catch(err => {
            event.sender.send('store_reply', null);
            log.error("store err::", err, "store_method::", method, "store_payload::", JSON.stringify(payload))
        })
    }

    /*
     * 操作处理
     */
    ipcMain.on("operate", ipc_operate)

    function ipc_operate(event, arg) {
        let operate = arg.operate;
        let parameter = arg.parameter;

        //程序退出
        let func_quit = function () {
            app.quit()
        }

        //登录成功
        let func_login = function () {
            loginWindow.close()//关闭登录页面
            createMainWindow()//打开业务主页面
        }
        //退出登录
        let func_logout = function () {
            mainWindow.close()//关闭业务页面
            createLoginWindow()//打开登录主页面
        }

        switch (operate) {
            case "quit"://程序退出
                func_quit()
                break;
            case "login"://程序退出
                func_login()
                break;
            case "logout":
                func_logout()
                break;
        }
        // event.sender.send('notify_reply', {"msg": "通知发送成功", "success": true});
    }


    // 网络请求
    ipcMain.on("http", ipc_http)

    function ipc_http(event, arg) {
        let timeStamp = Date.now() // 用来记录对应各请求的request和response
        log.info("stamp::", timeStamp, "http request::", JSON.stringify(arg),)
        let url = arg.url;
        let method = arg.method;
        let parameter = arg.parameter;

        let config = {
            method: method,
            url: url,
            data: parameter,
            headers: {
                'Cookie': ''
            }
        };
        axios(config)
            .then(function (response) {
                log.info("stamp::", timeStamp, "http response::", JSON.stringify(response.data))
                event.sender.send('http_reply', response.data);
            })
            .catch(function (error) {
                log.error("stamp::", timeStamp, "http response::", error)
                let ro = utils.getResponseObject()
                ro.msg = error.toString()
                event.sender.send('http_reply', ro);
            });
    }


    // 业务主页面
    let mainWindow

    function createMainWindow() {
        // 解决electron console里提示安全warning的问题
        process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';

        // 如果开发环境直接获取8080，如果生产环境直接读取index.html
        const winURL = process.env.NODE_ENV === 'development'
            ? `http://localhost:8080`
            : `${__dirname}/vue/dist/index.html`

        // Create the browser window.
        mainWindow = new BrowserWindow({
            width: 1200,//400
            height: 750,//750
            frame: false,
            webPreferences: {
                preload: path.join(__dirname, 'preload.js'),
                contextIsolation: false,
                enableRemoteModule: true,//来打开remote模块，使得渲染进程中可以调用主进程的方法
                nodeIntegration: true,//渲染进程也可以使用node模块，electron引入vue，vue所在窗口与主线程交互必须要true
            },
        })

        // AND LOAD THE INDEX.HTML OF THE APP.
        if (process.env.NODE_ENV === 'development') {
            mainWindow.loadURL(winURL)
        } else {
            mainWindow.loadFile(winURL)
        }

        // Open the DevTools.//FIXME
        mainWindow.webContents.openDevTools();
    }


    // 登录页面
    let loginWindow

    function createLoginWindow() {
        // 解决electron console里提示安全warning的问题
        process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';

        const loginWinURL = process.env.NODE_ENV === 'development'
            ? `http://localhost:8080/#/login`
            : `${__dirname}/vue/dist/index.html/#/login`
        // Create the browser window.
        loginWindow = new BrowserWindow({
            width: 400,
            height: 400,
            frame: false,
            center: true,
            webPreferences: {
                preload: path.join(__dirname, 'preload.js'),
                contextIsolation: false,
                enableRemoteModule: true,//来打开remote模块，使得渲染进程中可以调用主进程的方法
                nodeIntegration: true,//渲染进程也可以使用node模块，electron引入vue，vue所在窗口与主线程交互必须要true
            },
        })

        // AND LOAD THE INDEX.HTML OF THE APP.
        if (process.env.NODE_ENV === 'development') {
            loginWindow.loadURL(loginWinURL)
        } else {
            loginWindow.loadFile(loginWinURL)
        }

        // Open the DevTools.//FIXME
        loginWindow.webContents.openDevTools();
    }

    // 单机db存储
    let store = level('config');


    return {
        "store": store,
        "createMainWindow": createMainWindow,
        "createLoginWindow": createLoginWindow
    }

}();


module.exports = hook;