import $ from 'jquery';
import i18next from 'i18next/lib';

export function GET(url, success) {
    $.ajax({
        method: 'GET',
        url: API_HOST + url,
        data: {locale: i18next.language},
        dataType: 'json',
        success: success
    });

}
export function POST(url, data, success) {
    data.locale = i18next.language;
    $.ajax({
        method: 'POST',
        url: API_HOST + url,
        data: data,
        dataType: 'json',
        success: success
    });

}
function PATCH() {
    console.log('todo')
}
function DELETE() {
    console.log('todo')
}
