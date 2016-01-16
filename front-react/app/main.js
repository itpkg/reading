require('bootstrap/dist/css/bootstrap.css');
require('bootstrap/dist/css/bootstrap-theme.css');
require('./main.css');

import React from 'react'
import { render } from 'react-dom'
import { compose, createStore, combineReducers, applyMiddleware } from 'redux';
import { Provider } from 'react-redux'
import { Router } from 'react-router'

import { syncHistory, routeReducer } from 'redux-simple-router'
import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import Route from './components/route'
import reducers from './reducers'


function main(options) {
    i18next
        .use(XHR)
        .use(LanguageDetector)
        .init({
            fallbackLng: 'en-US',
            backend: {
                loadPath: API_HOST + '/locales/{{lng}}'
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',
                caches: ['localStorage', 'cookie']
            }
        }, (err, t)=> {
            const reducer = combineReducers(Object.assign({}, reducers, {
                routing: routeReducer
            }));
            const reduxRouterMiddleware = syncHistory(options.history);
            const createStoreWithMiddleware = applyMiddleware(reduxRouterMiddleware)(createStore);
            const store = createStoreWithMiddleware(reducer);


            render(
                <Provider store={store}>
                    <Router history={options.history}>
                        {Route}
                    </Router>
                </Provider>,
                document.getElementById('root')
            );
        });
}

export default main