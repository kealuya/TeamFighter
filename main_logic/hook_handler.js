const {app, ipcMain, BrowserWindow} = require('electron')
let hook = function () {
    // 所有方法必须异步返回，不能同步，不然可能发生主线程阻塞
    console.log('hook_handler init')

    // getVersion 获取程序版本 来自package.json
    ipcMain.on("get_version", (event, arg) => {
        event.sender.send('get_version_reply', app.getVersion());
    })

}();


