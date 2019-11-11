//删除用户
function del_user(talkerID) {
    $.ajax({
        url: "/delete_talker.action",
        data: {
            talkerID: talkerID,
        },
        type: "POST",
        dataType: "json",
        async: false,
        success: function (data) {
            if (data.status == "success") {
                $("#" + talkerID).remove();
            } else {

            }
        }
    })
}
//切换聊天对象
function ex_talker(talkerID) {
    $("#receiver").val(talkerID);
    $.ajax({
        url: "",
        data: {},
        type: "POST",
        dataType: "json",
        async: false,
        success: function (data) {

        }
    })
}
//展示聊天对象信息
function show_userInfo(obj) {

}
//发送聊天消息
function send_message() {
    var receiverID = $("#receiver").val();
    var message = $("#message").val();
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
