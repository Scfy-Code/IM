function checkEmail(obj, pat) {
    var email = $(obj).val();
    if (pat.search(email)) {
        $(obj).removeClass("is-invalid");
        $(obj).addClass("is-valid");
        $(obj).next().removeClass("invalid-feedback").addClass("valid-feedback").text("验证通过");
        return true;
    } else {
        $(obj).removeClass("is-valid");
        $(obj).addClass("is-invalid");
        $(obj).next().removeClass("valid-feedback").addClass("invalid-feedback").text("请输入有效的邮箱");
        return false;
    }
}
function checkPassword(obj, pat) {
    var password = $(obj).val();
    if (pat.search(password)) {
        $(obj).removeClass("is-invalid");
        $(obj).addClass("is-valid");
        $(obj).next().removeClass("invalid-feedback").removeClass("valid-feedback").addClass("valid-feedback").text("验证通过");
        return true;
    } else {
        $(obj).removeClass("is-valid");
        $(obj).addClass("is-invalid");
        $(obj).next().removeClass("valid-feedback").addClass("invalid-feedback").text("请输入有效的密码");
        return false;
    }
}
function Valid(pat1, pat2) {
    return checkEmail($("#email"), pat1) && checkPassword($("#password"), pat2);
}