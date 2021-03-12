// vue.config.js
module.exports = {
    publicPath: './',
    assetsDir: './',
    css: {
        loaderOptions: {
            less: {
                modifyVars: {
                    // 直接覆盖变量
                    'sidebar-selected-border-color': '@red',
                    'sidebar-width':'63px',
                    'popover-action-width':'80px',
                    // 'popover-action-height':'22px',
                    // 'popover-action-font-size':'10px',
                },
            },
        },
    },
}
