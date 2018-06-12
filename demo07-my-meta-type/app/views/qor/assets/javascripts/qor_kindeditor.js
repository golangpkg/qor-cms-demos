
(function(factory) {
    if (typeof define === 'function' && define.amd) {
        // AMD. Register as anonymous module.
        define(['jquery'], factory);
    } else if (typeof exports === 'object') {
        // Node / CommonJS
        factory(require('jquery'));
    } else {
        // Browser globals.
        factory(jQuery);
    }
})(function($) {
    'use strict';

    let componentHandler = window.componentHandler,
        NAMESPACE = 'qor.kindeditor',
        EVENT_ENABLE = 'enable.' + NAMESPACE,
        EVENT_DISABLE = 'disable.' + NAMESPACE,
        EVENT_UPDATE = 'update.' + NAMESPACE,
        SELECTOR_COMPONENT = '[class*="mdl-js"],[class*="mdl-tooltip"]';

    function enable(target) {
        //console.log("################### kindeditor enable ###################");
        /*jshint undef:false */
        var editor = KindEditor.create('#kindeditor-id',{
            uploadJson : '/common/kindeditor/upload?dir=image',
            allowFileManager : false,
            afterBlur: function(){this.sync();}
        });
    }

    function disable(target) {
        //console.log("################### kindeditor disable ###################");
    }

    $(function() {
        $(document)
            .on(EVENT_ENABLE, function(e) {
                enable(e.target);
            })
            .on(EVENT_DISABLE, function(e) {
                disable(e.target);
            })
            .on(EVENT_UPDATE, function(e) {
                disable(e.target);
                enable(e.target);
            });
    });
});

//console.log("################### kindeditor load finish ###################");