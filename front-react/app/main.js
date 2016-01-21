require('bootstrap/dist/css/bootstrap.css');
require('bootstrap/dist/css/bootstrap-theme.css');
require('./main.css');

import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';


import React from 'react'
import ReactDOM from 'react-dom'

import Root from './root'
import Store from './store'

i18next
    .use(XHR)
    .use(LanguageDetector)
    .init({
        fallbackLng: 'en-US',
        backend: {
            loadPath: API_HOST + '/locales/{{lng}}',
            crossDomain: process.env.NODE_ENV !== 'production'
        },
        detection: {
            order: ['querystring', 'localStorage', 'cookie', 'navigator'],
            lookupQuerystring: 'locale',
            lookupCookie: 'locale',
            lookupLocalStorage: 'locale',
            caches: ['localStorage', 'cookie']
        }
    }, (err, t)=> {
        ReactDOM.render(<Root store={Store}/>, document.getElementById('root'));
    });



