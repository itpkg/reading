import i18n from 'i18next'

export function get(url, done, fail) {
    call("get", url, null, done, fail);
}

function call(method, url, data, done, fail) {

    if (!fail) {
        fail = function (e) {
            console.log(e);
        }
    }

    fetch(`${process.env.CONFIG.backend}/${i18n.language}/api${url}`,
        {
            method: method,
            data: data
        })
        .then(res => res.json())
        .then(done).catch(fail);
}