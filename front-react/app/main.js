require('./main.css');
import React from 'react'
import { render } from 'react-dom'
import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';


import store from './store';
import Root from './containers/Root'


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
        render(<Root store={store}/>, document.getElementById('root'));
    });

