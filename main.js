// Modules to control application life and create native browser window
const {app, ipcMain, BrowserWindow} = require('electron')
const path = require('path')
_ = require('./main_logic/hook_handler') // 引入主进程逻辑

process.env.NODE_ENV = "development" // FIXME 修改生产环境

let mainWindow
// 如果开发环境直接获取8080，如果生产环境直接读取index.html
const winURL = process.env.NODE_ENV === 'development'
    ? `http://localhost:8080`
    : `${__dirname}/vue/dist/index.html`

function createMainWindow() {
    // 解决electron console里提示安全warning的问题
    process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';

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
function createLoginWindow() {
    // 解决electron console里提示安全warning的问题
    process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';

    const loginWinURL = process.env.NODE_ENV === 'development'
        ? `http://localhost:8080/#/login`
        : `${__dirname}/vue/dist/index.html`
    // Create the browser window.
    let loginWindow = new BrowserWindow({
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


// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
    createLoginWindow()
    // createMainWindow()
    app.on('activate', function () {
        // On macOS it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (BrowserWindow.getAllWindows().length === 0) createMainWindow()
    })
})

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit()
})

