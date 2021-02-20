const {app, ipcMain, BrowserView, BrowserWindow} = require('electron')
let hook = function () {
    // 所有方法必须异步返回，不能同步，不然可能发生主线程阻塞
    console.log('hook_handler init')

    // getVersion 获取程序版本 来自package.json
    ipcMain.on("get_version", (event, arg) => {
        event.sender.send('get_version_reply', app.getVersion());
    })

    // getVersion 获取程序版本 来自package.json
    ipcMain.on("expand", (event, arg) => {
        console.log(arg)

        const win = new BrowserWindow({
            width: 400, height: 750,
            title: "企业微信消息推送",
            resizable: false,
            frame:false,
        })

        const view = new BrowserView()
        win.setBrowserView(view)
        view.setBounds({x: 0, y: 0, width: 400, height: 750})
        view.webContents.loadURL('http://pioneer.szhtkj.com.cn:8082/analyse')


        event.sender.send('expand_reply', app.getVersion());
    })

}();


