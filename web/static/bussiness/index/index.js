function del_user(obj) {
    $(obj).parent().parent().parent().parent().remove();
}
$(function () {
    $('[data-toggle="tab"]').tooltip();
})
