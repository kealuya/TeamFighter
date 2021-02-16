module.exports = {
    presets: [
        '@vue/cli-plugin-babel/preset',
    ],
    // 引入vant自动按需加载组件plugin
    plugins: [
        ["import", {
            "libraryName": "vant",
            "libraryDirectory": "es",
            // 指定样式路径
            style: (name) => `${name}/style/less`,
        }, 'vant',
        ],
    ]
}
