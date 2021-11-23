const serverHost = "https://wxapi.ericcai.fun"
// app.js
App({
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        console.log(res)
        wx.request({
          url: this.globalData.serverHost + '/sns/jscode2session', //仅为示例，并非真实的接口地址
          data: {
            js_code: res.code,
          },
          header: {
            'content-type': 'application/json' // 默认值
          },
          success: (res) => {
            this.globalData.openid = res.data.openid
            this.globalData.launched = true
            this.globalData.launchedCallback4Index()
          },
          fail(res) {
            console.log(res)
          },
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
  },
  globalData: {
    userInfo: null,
    openid: null,
    serverHost: serverHost,
    launched: false,
    launchedCallback4Index: () => {}
  }
})