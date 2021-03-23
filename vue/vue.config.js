// vue.config.js
let all_color = "#29b7cb"//蔚蓝
module.exports = {
    publicPath: './',
    assetsDir: './',
    css: {
        loaderOptions: {
            less: {
                modifyVars: {
                    // 直接覆盖变量
                    //========颜色========
                    //侧边导航选择后的边颜色
                    'sidebar-selected-border-color': all_color,
                    'button-primary-background-color': all_color,
                    'button-primary-border-color': all_color,
                    'rate-icon-full-color': all_color,
                    'tabs-line-color': all_color,
                    'tabs-default-color': all_color,
                    'slider-active-background-color': all_color,
                    'tag-primary-color': all_color,
                    'badge-background-color': all_color,
                    //================================
                    'sidebar-width': '63px',
                    'popover-action-width': '80px',
                    // 'popover-action-height':'22px',
                    // 'popover-action-font-size':'10px',
                },
            },
        },
    },
}
