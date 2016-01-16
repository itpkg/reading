import $ from 'jquery'
import i18next from 'i18next/lib';

function api(u) {
    return API_HOST + u + (u.indexOf('?') == -1 ? '?' : '&') + 'locale=' + i18next.language;
}

export const AjaxMixin = {
    GET: function (url, success) {
        $.get(api(url), success.bind(this));
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