import $ from 'jquery'

export const AjaxMixin = {
    GET: function (url, success) {
        $.get(API_HOST + url, success.bind(this));
    },
    POST: function () {
        console.log('todo post');
    },
    PUT: function () {
        console.log('todo put');
    },
    PATCH: function () {
        console.log('todo patch');
    },
    DELETE: function () {
        console.log('todo delete');
    }
};