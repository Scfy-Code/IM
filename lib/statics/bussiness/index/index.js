var socket = null;
$(function () {
    socket = new WebSocket("ws://127.0.0.1:8080/talk.action")
    socket.onopen = function (evt) {
        console.log("Connection open ...");
    };

    socket.onmessage = function (evt) {
        console.log("Received Message: " + evt.data);
    };

    socket.onclose = function (evt) {
        console.log("Connection closed.");
    };

    $('#message').bind('input propertychange', function () {
        var message = $("#message").val();
        if (message = null || "" == message) {
            $("#send").attr("disabled", true);
        } else {
            $("#send").removeAttr("disabled");
        }
    });
});
function sendMesage() {
    var message = $("#message").val();
    socket.send(message);
    $("#message").val("");
    $("#send").attr("disabled", true);
}