// Modules to control application life and create native browser window
const {app, ipcMain, BrowserWindow} = require('electron')
const hook = require('./hook_handler') // 引入主进程逻辑


process.env.NODE_ENV = "development" // FIXME 修改生产环境

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {

    hook.createLoginWindow()
    // createMainWindow()
    app.on('activate', function () {
        // On macOS it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (BrowserWindow.getAllWindows().length === 0) hook.createMainWindow()
    })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit()
})

