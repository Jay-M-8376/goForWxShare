(function (window) {

    var shareTitle = '';
    var shareDesc = '';
    var shareImg = '';
    var shareUrl = '';

    function setConfig(appid, timestamp, randomStr, signature) {
        wx.config({
            debug: false,    //调试模式
            appId: appid,
            timestamp: timestamp,
            nonceStr: randomStr,
            signature: signature,
            jsApiList: [
                'checkJsApi',
                'onMenuShareTimeline',
                'onMenuShareAppMessage',
                'onMenuShareQQ'
            ]
        });
        wx.ready(function () {
            wx.checkJsApi({
                jsApiList: ['onMenuShareTimeline', 'onMenuShareAppMessage', 'onMenuShareQQ'],
                success: function (res) {
                    wx.onMenuShareTimeline({////分享朋友圈
                        title: shareDesc, // 分享标题
                        link: shareUrl, // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
                        imgUrl: shareImg, // 分享图标
                        success: function () {
                            // 用户确认分享后执行的回调函数
                            console.log('share success');
                        },
                        cancel: function () {
                            // 用户取消分享后执行的回调函数
                            console.log('cancle share behavior');
                        }
                    });

                    wx.onMenuShareAppMessage({
                        title: shareTitle, // 分享标题
                        desc: shareDesc, // 分享描述
                        link: shareUrl, // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
                        imgUrl: shareImg, // 分享图标
                        type: "link", // 分享类型,music、video或link，不填默认为link
                        dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
                        success: function () {
                            // 用户确认分享后执行的回调函数
                            console.log('share to appMessage success');
                        },
                        cancel: function () {
                            // 用户取消分享后执行的回调函数
                            console.log('cancle share to appMessage');
                        }
                    });

                    wx.onMenuShareQQ({
                        title: shareTitle, // 分享标题
                        desc: shareDesc, // 分享描述
                        link: shareUrl, // 分享链接
                        imgUrl: shareImg, // 分享图标
                        success: function () {
                            // 用户确认分享后执行的回调函数
                            console.log('share to qq success')
                        },
                        cancel: function () {
                            // 用户取消分享后执行的回调函数
                            console.log('cancle share to qq')
                        }
                    });
                }
            })
        });
        wx.error(function (res) {
            console.log(res);
        })
    }


    function getDatas() {
        var request = new XMLHttpRequest();
        request.open('post','http://project.mlord.cn:8999/api/share',true);
        //request.open('post', 'http://project.mlord.cn:8085/getIP/', true);
        //request.setRequestHeader("data-Type","json");
        request.onreadystatechange = function (req, e) {
            if (req.currentTarget.readyState == 4 && req.currentTarget.status == 200) {
                var data = JSON.parse(req.currentTarget.responseText);
                setConfig(data.AppId,data.Timestamp,data.NonceStr,data.Signature);
            } else {
                console.log('readyStatus : ' + req.currentTarget.readyState + ' \n status : ' + req.currentTarget.status);
            }
        }
        request.send(JSON.stringify({url:window.location.href}));
    }

    function setShare(title, desc, img, url) {
        shareTitle = title;
        shareDesc = desc;
        shareImg = img;
        shareUrl = url;
        getDatas();
    }

    window.setShare = setShare;

})(window);