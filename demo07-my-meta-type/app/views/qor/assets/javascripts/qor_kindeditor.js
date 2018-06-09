$(function($) {
    var editor = KindEditor.create('#Content',{
        uploadJson : '/admin/kindeditor/upload',
        allowFileManager : false,
        afterBlur: function(){this.sync();}
    });
    var uploadEditor = KindEditor.editor({
        uploadJson : '/admin/kindeditor/upload?dir=image',
        allowFileManager : false
    });
});
console.log("################### kindeditor load finish ###################");