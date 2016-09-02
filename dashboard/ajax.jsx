import i18n from 'i18next'

import {TOKEN} from './constants'

export function get(url, done, fail) {
    call("GET", url, null, done, fail);
}

export function post(url, form, done, fail) {
    call("POST", url, form, done, fail);
}

function call(method, url, data, done, fail) {
    if(!done){
        done = function (rst) {
            console.log(rst);
        }
    }

    if (!fail) {
        fail = function (err) {
            console.log(err);
        }
    }

    fetch(`${process.env.CONFIG.backend}/${i18n.language}/api${url}`,
        {
            headers:{
                'Authorization':'Bearer ' + window.sessionStorage.getItem(TOKEN),
                //'Access-Control-Allow-Origin': 'http://localhost:4200',
            },
            method: method,
            body: data,
            mode: 'cors',
        })
        .then(res => res.json())
        .then(done).catch(fail);
}