// pages/new_event/new_event.js
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    form: {
      frequency: 0,
      frequencyDetail: 0,
      time: "00:00",
      content: ""
    },
    btnDisabled: false,
    msg: "",
    msgType: "success",
    frequencyOptions: app.globalData.frequencyOptions,
    frequencyDetailWeekOptions: app.globalData.frequencyDetailWeekOptions,
    frequencyDetailMonthOptions: app.globalData.frequencyDetailMonthOptions,
  },
  onFrequencyPickerChange(e) {
    var form = this.data.form
    form.frequency = e.detail.value
    this.setData({
      form: form
    })
  },
  onFrequencyDetailPickerChange(e) {
    var form = this.data.form
    form.frequencyDetail = e.detail.value
    this.setData({
      form: form
    })
  },
  onTimePickerChange(e) {
    var form = this.data.form
    form.time = e.detail.value
    this.setData({
      form: form
    })
  },
  onInputChange(e) {
    var form = this.data.form
    form.content = e.detail.value
    this.setData({
      form: form
    })
  },
  submitForm() {
    this.setData({
      btnDisabled: true
    })
    wx.request({
      url: app.globalData.serverHost + '/notice/v1/put',
      method: "POST",
      data: {
        notice: {
          openid: app.globalData.openid,
          content: this.data.form.content,
          time: this.data.form.time,
          frequency: this.data.form.frequency,
          frequency_detail: this.data.form.frequencyDetail.toString(),
        }
      },
      header: {
        'content-type': 'application/json' // 默认值
      },
      success: (res) => {
        this.setData({
          msg: "操作成功",
          msgType: "success"
        })
        // 在C页面内 navigateBack，将返回A页面
        wx.navigateBack({
          delta: 1
        })
        console.log(res.data)
      },
      fail(res) {
        console.log(res)
        this.setData({
          msg: res.data.msg,
          msgType: "error"
        })
      },
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
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