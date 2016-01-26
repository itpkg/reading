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
