function checkEmail(pat) {
    var email = $("#email").val();
    $("#msg").text("");
    if (pat.test(email)) {
        $("#email").addClass("is-valid").removeClass("is-invalid");
        $("#emailMsg").removeClass("invalid-feedback").addClass("valid-feedback").text("验证通过");
        return true;
    } else {
        $("#email").addClass("is-invalid").removeClass("is-valid");
        $("#emailMsg").removeClass("valid-feedback").addClass("invalid-feedback").text("请输入有效的邮箱");
        return false;
    }
}
function checkPassword(pat) {
    var password = $("#password").val();
    $("#msg").text("");
    if (pat.test(password)) {
        $("#password").addClass("is-valid").removeClass("is-invalid");
        $("#passwordMsg").removeClass("invalid-feedback").removeClass("valid-feedback").addClass("valid-feedback").text("验证通过");
        return true;
    } else {
        $("#password").addClass("is-invalid").removeClass("is-valid");
        $("#passwordMsg").removeClass("valid-feedback").addClass("invalid-feedback").text("请输入有效的密码");
        return false;
    }
}
function Valid(pat1, pat2) {
    return checkEmail($("#email"), pat1) && checkPassword($("#password"), pat2);
}
