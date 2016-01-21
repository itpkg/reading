import $ from 'jquery';
import i18next from 'i18next/lib';

export function GET(url, done, fail) {
    $.ajax({
        method: 'GET',
        url: API_HOST + url,
        data: {locale: i18next.language},
        crossDomain: true,
        beforeSend: function (xhr) {
            var token = sessionStorage.getItem('token');
            if (token) {
                //xhr.setRequestHeader('Access-Control-Allow-Origin', '*');
                //xhr.setRequestHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
                //xhr.setRequestHeader('Access-Control-Allow-Headers', 'Authorization');
                xhr.setRequestHeader('Authorization', 'Bearer ' + token);
            }

        },
        dataType: 'json'
    }).done(done).fail(fail);

}
export function POST(url, data, done, fail) {
    data.locale = i18next.language;
    $.ajax({
        method: 'POST',
        url: API_HOST + url,
        data: data,
        dataType: 'json'
    }).done(done).fail(fail);

}
function PATCH() {
    console.log('todo')
}
function DELETE() {
    console.log('todo')
}
