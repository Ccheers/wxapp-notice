// index.js
// 获取应用实例
const app = getApp()
const subTmplID = "TPxDR87VAsidwvX6UKdlVPbKVdwPl4qnhxRV8qFoAYA"
Page({
  data: {
    can_sub: true,
    loading: true,
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    canIUseGetUserProfile: false,
    canIUseOpenData: wx.canIUse('open-data.type.userAvatarUrl') && wx.canIUse('open-data.type.userNickName') // 如需尝试获取用户信息可改为false
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },
  // 事件处理函数
  toEventList() {
    wx.navigateTo({
      url: '../event_list/event_list'
    })
  },
  // 事件处理函数
  toNewEvent() {
    wx.navigateTo({
      url: '../new_event/new_event'
    })
  },
  onLoad() {
    if (wx.getUserProfile) {
      this.setData({
        canIUseGetUserProfile: true
      })
    }

    if (!app.globalData.launched) {

      app.globalData.launchedCallback4Index = () => {
        this.checkSub()
        this.setData({
          loading: false
        })
      }
      
    }else{
      this.checkSub()
        this.setData({
          loading: false
        })
    }

  },
  checkSub() {
    var subed = true
    wx.getSetting({
      withSubscriptions: true,
      success(res) {
        if (res['subscriptionsSetting'][subTmplID] == "accept") {
          subed = false
        }
      },
    })
    if (!subed) {
      this.setData({
        can_sub: true
      })
      return
    }

    wx.request({
      url: app.globalData.serverHost + '/notice/v1/check', //仅为示例，并非真实的接口地址
      method: "POST",
      data: {
        openid: app.globalData.openid,
      },
      header: {
        'content-type': 'application/json' // 默认值
      },
      success: (res) => {
        console.log(res.data)

        this.setData({
          can_sub: !res.data.subscribed
        })
      },
      fail: (res) => {
        this.setData({
          can_sub: true
        })
        console.log(res)
      },
    })
  },
  subNotice() {
    wx.requestSubscribeMessage({
      tmplIds: [subTmplID],
      success: (e) => {
        if (e[subTmplID] == "accept") {
          wx.request({
            url: app.globalData.serverHost + '/notice/v1/sub', //仅为示例，并非真实的接口地址
            method: "POST",
            data: {
              openid: app.globalData.openid
            },
            header: {
              'content-type': 'application/json' // 默认值
            },
            success: (res) => {
              this.checkSub()
              console.log(res.data)
            },
            fail(res) {
              console.log(res)
            },
          })
        }
      },
      fail() {},
      complete() {},
    })

  },
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },
})