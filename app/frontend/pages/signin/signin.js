// pages/signin/signin.js
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        formatter: (day) => {},
        msg: "",
        msgType: "",
        list: {},
        showList: false,
    },
    // 签到列表
    getSigninList() {
        wx.request({
            url: app.globalData.serverHost + '/signin/list',
            data: {
                openid: app.globalData.openid
            },
            success: (res) => {
                this.setData({
                    list: res.data.list,
                    showList: true,
                    formatter: function (day) {
                        var key = day.date.getFullYear() + "-" + (day.date.getMonth() + 1) + "-" + day.date.getDate()
                        if (res.data.list[key]) {
                            day.bottomInfo = "已签到"
                        }
                        return day
                    }
                })
            },
            fail: () => {
                this.alert("网络异常，请稍后再试", false)
            }
        })
        return this.data.list
    },
    // 签到
    signin() {

        wx.request({
            url: app.globalData.serverHost + '/signin',
            method: "POST",
            data: {
                openid: app.globalData.openid
            },
            success: (res) => {

            },
            fail: () => {
                this.alert("网络异常，请稍后再试", false)
            }
        })
    },
    alert(msg, isSuccess) {
        this.setData({
            msg: msg,
            msgType: (isSuccess ? "success" : "danger")
        })
        setTimeout(() => {

            this.setData({
                msg: "",
                msgType: ""
            })

        }, 2000)
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.signin()
        this.getSigninList()
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