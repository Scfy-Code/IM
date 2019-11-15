function index() {
    window.location.href = "/";
}
function user() {

}
//删除用户
function del_talker(bindID) {
    $.ajax({
        url: "/delete_talker.action",
        data: {
            bindID: bindID,
        },
        type: "POST",
        dataType: "json",
        async: false,
        success: function (data) {
            if (data.status == "success") {
                $("#" + bindID).remove();
            }
            if (data.status == "failure") {

            }
        }
    });
}
// 退出群组
function quit_team(bindID) {
    $.ajax({
        url: "/quit_team.action",
        data: {
            bindID: bindID,
        },
        type: "POST",
        dataType: "json",
        async: false,
        success: function (data) {
            if (data.status == "success") {
                $("#" + bindID).remove();
            }
            if (data.status == "failure") {

            }
        }
    })
}
//切换聊天对象
function ex_talker(bindID) {
    var bd = $("#receiver").val()
    if (bd == "") {

    } else {
        $("#" + bd).attr("class", "talker-inline");
    }
    $("#receiver").val(bindID);
    $("#" + bindID).attr("class", "talker-selected");
}
//切换聊天对象
function ex_team(bindID) {
    var bd = $("#receiver").val()
    if (bd == "") {

    } else {
        $("#" + bd).attr("class", "team-inline");
    }
    $("#receiver").val(bindID);
    $("#" + bindID).attr("class", "team-selected");
}
//展示聊天对象信息
function show_talkerInfo(obj) {

}
//发送聊天消息
function send_message() {
    var receiverID = $("#receiver").val();
    var message = $("#textmessage").val();
    $.ajax({
        url: "",
        data: {
            receiverID: receiverID,
            message: message
        },
        type: "POST",
        dataType: "json",
        async: false,
        success: function (data) {

        }
    })
    return false
}
$(function () {
    $('[data-toggle="tab"]').tooltip();
})
