// pages/event_list/event_list.js
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    slideButtons: [{
      type: "warn",
      text: "删除"
    }],
    eventList: [],
    frequencyOptions: app.globalData.frequencyOptions,
    frequencyDetailWeekOptions: app.globalData.frequencyDetailWeekOptions,
    frequencyDetailMonthOptions: app.globalData.frequencyDetailMonthOptions,
  },
  onTapSlideButton(data) {
    console.log(this.data.eventList[data.detail.index].id)
    wx.request({
      url: app.globalData.serverHost + '/notice/v1/del', //仅为示例，并非真实的接口地址
      method: "POST",
      data: {
        id: this.data.eventList[data.detail.index].id,
        openid: this.data.eventList[data.detail.index].openid
      },
      header: {
        'content-type': 'application/json' // 默认值
      },
      success: (res) => {
        this.getList()
      },
      fail(res) {
        console.log(res)
      },
    })
  },
  getList() {
    wx.request({
      url: app.globalData.serverHost + '/notice/v1/list', //仅为示例，并非真实的接口地址
      method: "POST",
      data: {
        openid: app.globalData.openid,
      },
      header: {
        'content-type': 'application/json' // 默认值
      },
      success: (res) => {
        var list = [];
        res.data.noticeInfo.forEach(element => {
          var frequencyDetailDisplay = element.frequencyDetail
          switch (element.frequency) {
            case "0":
              frequencyDetailDisplay = element.frequencyDetail
              if (frequencyDetailDisplay == "0") {
                frequencyDetailDisplay = "今天"
              }
              break
            case "2":
              frequencyDetailDisplay = this.data.frequencyDetailWeekOptions[element.frequencyDetail]
              break
            case "3":
              frequencyDetailDisplay = this.data.frequencyDetailMonthOptions[element.frequencyDetail]
              break
            default:
              frequencyDetailDisplay = element.frequencyDetail
              break
          }
          list.push({
            content: element.content,
            frequency: element.frequency,
            frequencyDetail: element.frequencyDetail,
            frequencyDetailDisplay: frequencyDetailDisplay,
            id: element.id,
            openid: element.openid,
            time: element.time
          })
        });

        this.setData({
          eventList: list
        })
      },
      fail: (res) => {
        console.log(res)
      },
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.getList()
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})