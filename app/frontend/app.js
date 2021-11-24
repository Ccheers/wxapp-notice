const serverHost = "https://wxapi.ericcai.fun"
// app.js
App({
  onLaunch() {
    // 初始化月份选项
    this.initvarFrequencyDetailMonthOptions()

    var openid = wx.getStorageSync('openid')
    if (openid) {
      this.globalData.openid = openid
      this.globalData.launched = true
      this.globalData.launchedCallback4Index()
    } else {
      this.login()
    }
    
    // 引入 color ui
    wx.getSystemInfo({
      success: e => {
        this.globalData.StatusBar = e.statusBarHeight;
        let custom = wx.getMenuButtonBoundingClientRect();
        this.globalData.Custom = custom;  
        this.globalData.CustomBar = custom.bottom + custom.top - e.statusBarHeight;
      }
    })
  },
  login() {
    // 登录
    wx.login({
      success: (res) => {
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
            wx.setStorageSync('openid', this.globalData.openid)
          },
          fail(res) {
            console.log(res)
          },
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
  },
  initvarFrequencyDetailMonthOptions(){
    var frequencyDetailMonthOptions = ["请选择"]
    for (var i = 1; i < 32; i++) {
      frequencyDetailMonthOptions.push(i.toString() + "号")
    }
    this.globalData.frequencyDetailMonthOptions = frequencyDetailMonthOptions
  },
  globalData: {
    userInfo: null,
    openid: null,
    serverHost: serverHost,
    launched: false,
    launchedCallback4Index: () => {},
    frequencyOptions: [
      "一次",
      "工作日",
      "每周",
      "每月",
      "每天"
    ],
    frequencyDetailWeekOptions: [
      "周日",
      "周一",
      "周二",
      "周三",
      "周四",
      "周五",
      "周六",
    ],
    frequencyDetailMonthOptions: []
  }
})