const {app, ipcMain, BrowserView, BrowserWindow, screen} = require('electron')
let hook = function () {
    // 所有方法必须异步返回，不能同步，不然可能发生主线程阻塞
    console.log('hook_handler init')

    // getVersion 获取程序版本 来自package.json
    ipcMain.on("get_version", (event, arg) => {
        event.sender.send('get_version_reply', app.getVersion());
    })

    // 扩展功能
    ipcMain.on("expand", (event, arg) => {
        console.log(arg)

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
    })


    // 消息通知
    ipcMain.on("notify", (event, arg) => {
        const {width, height} = screen.getPrimaryDisplay().workAreaSize

        const win = new BrowserWindow({
            width: 500, height: 120,
            title: "TeamFighter -- 通知",
            resizable: false,
            x: width - 500,
            y: height - 120,
            alwaysOnTop: true,
            frame: false,
            webPreferences: {
                contextIsolation: false,
                enableRemoteModule: true,//来打开remote模块，使得渲染进程中可以调用主进程的方法
                nodeIntegration: true,//渲染进程也可以使用node模块，electron引入vue，vue所在窗口与主线程交互必须要true
            },
        })
        // win.loadURL("https://img01.yzcdn.cn/vant/logo.png")
        win.loadFile(`${__dirname}/page_file/notify.html`)
        win.webContents.on('did-finish-load', function () {
            win.webContents.send('param', {"content": "我是消息哦"});
        });
        // win.webContents.openDevTools();


        event.sender.send('notify_reply', {"msg": "通知发送成功", "success": true});
    })

}();


