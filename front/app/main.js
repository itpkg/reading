require('./main.css');

import React from 'react'
import ReactDOM from 'react-dom'

import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import RootRouter from './components/router'


function main(options) {
    i18next
        .use(XHR)
        .use(LanguageDetector)
        .init({
            fallbackLng: 'en-US',
            backend: {
                loadPath: '/locales/{{lng}}.json'
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',
                caches: ['localStorage', 'cookie']
            }
        }, (err, t)=> {
            ReactDOM.render(RootRouter, document.getElementById('root'));
        });
}

export default main