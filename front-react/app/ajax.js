import $ from 'jquery';
import i18next from 'i18next/lib';

const setToken = function (xhr) {
    var token = sessionStorage.getItem('token');
    if (token) {
        xhr.setRequestHeader('Authorization', 'Bearer ' + token);
    }

};

function appendLocale(url) {
    return API_HOST + url + (url.indexOf('?') === -1 ? '?' : '&') + 'locale=' + i18next.language;
}

export function failed() {
    alert(i18next.t('messages.failed'));
}

export function error(e) {
    return alert(e.responseText);
}
export function response(success) {
    return function (rst) {
        var msg = rst.messages ? '\n' + rst.messages.join('\n') : '';
        if (rst.ok) {
            alert(i18next.t('messages.success') + msg);
            if (success) {
                success();
            }
        } else {
            alert(i18next.t('messages.failed') + msg);
        }
    };
}

export function GET(url, done, fail) {
    $.ajax({
        method: 'GET',
        url: API_HOST + url,
        data: {locale: i18next.language},
        crossDomain: true,
        beforeSend: setToken,
        dataType: 'json'
    }).done(done).fail(fail);

}

export function POST(url, data, done, fail) {
    $.ajax({
        method: 'POST',
        url: appendLocale(url),
        data: data,
        beforeSend: setToken,
        dataType: 'json'
    }).done(done).fail(fail);

}

function PATCH() {
    console.log('todo')
}

export function DELETE(url, done, fail) {
    $.ajax({
        method: 'DELETE',
        url: appendLocale(url),
        beforeSend: setToken,
        dataType: 'json'
    }).done(done).fail(fail);
}

export function UPLOAD(url, data, done, fail) {
    $.ajax({
        method: 'POST',
        url: appendLocale(url),
        data: data,
        beforeSend: setToken,
        dataType: 'json',
        processData: false,
        contentType: false
    }).done(done).fail(fail);

}