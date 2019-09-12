var socket = null;
function getMessage(data) {
    var message = JSON.parse(data);
    // 1、判断当前会话用户是否是消息的发送者：如果是解析消息展现在会话框中？如果不是则在会话列表中添加标记
    var span = $("#talk" + message.id).children();
    if (span.length > 1) {
        var messageCount = parseInt($(span[1]).text()) + 1;
        if (messageCount >= 100) {
            $(span[1]).text('99+')
        } else {
            $(span[1]).text(messageCount)
        }
    } else {
        $(span[0]).after('<span class="badge badge-pill badge-danger float-right">1</span>')
    }
    alert(message.content)
}
$(function () {
    socket = new WebSocket("ws://127.0.0.1:8080/talk.action")
    socket.onopen = function (evt) {
        console.log("Connection open ...");
    };

    socket.onmessage = function (evt) {
        getMessage(evt.data);
    };

    socket.onclose = function (evt) {
        console.log("Connection closed.");
    };

    $('#message').bind('input propertychange', function () {
        var message = $("#message").val();
        var talkerID = $("#talker").attr("talkerID");
        if (message = null || "" == message || talkerID == null || talkerID == "") {
            $("#send").attr("disabled", "disabled");
        } else {
            $("#send").removeAttr("disabled");
        }
    });
    $('[data-toggle="popover"]').popover();
});
//发送文本消息
function sendTextMesage() {
    var receiver = $("#reciver").attr("reciver");
    var content = $("#textMsg").val();
    var msgType = "TEXT";
    $.ajax({
        url: "/msg/textMessage.action",
        type: "POST",
        data: {
            receiver: receiver,
            content: content,
        },
        async: false,
        cache: false,
        dataType: "json",
        success: function (data) {
            alert(data)
        },
    });
}
function exTalker(obj) {
    //1、判断点击的用户是否是当前聊天用户
    var talkerID = $(obj).attr("id");
    if (talkerID == $("talker").attr(talkerID)) {
        //选中的用户是当前的聊天用户，不执行任何操作
        return;
    } else {
        //选中的用户不是当前的聊天用户，切换聊天用户
        var talkerImage = $(obj).children("img").attr("src");
        var talkerRemarkName = $(obj).attr("name");
        var talkerSignature = "";
        var talkerID = $(obj).attr("id");
        $("#talkerImage").attr("src", talkerImage);
        $("#talkerRemarkName").text(talkerRemarkName);
        $("#talkerSignature").text(talkerSignature)
        $("#talker").attr("talkerID", talkerID);
    }
    //请求聊天部分聊天记录和最新的会话内容
    $.ajax({
        url: "",
        type: "POST",
        data: {},
        async: true,
        cache: false,
        dataType: "json",
        success: function () {

        },
    });
}