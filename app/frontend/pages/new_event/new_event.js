// pages/new_event/new_event.js
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
    frequencyOptions: [
      "一次",
      "工作日",
      "每周",
      "每月",
      "每天"
    ],
    frequencyDetailWeekOptions: [
      "周一",
      "周二",
      "周三",
      "周四",
      "周五",
      "周六",
      "周日",
    ],
    frequencyDetailMonthOptions: [],
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
  submitForm(){
    wx.request({
      url: 'example.php', //仅为示例，并非真实的接口地址
      data: this.data.form,
      header: {
        'content-type': 'application/json' // 默认值
      },
      success(res) {
        console.log(res.data)
      },
      fail(res) {
        console.log(res)
      },
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var frequencyDetailMonthOptions = []
    for (var i = 1; i < 32; i++) {
      frequencyDetailMonthOptions.push(i.toString() + "号")
    }
    this.setData({
      frequencyDetailMonthOptions: frequencyDetailMonthOptions
    })
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