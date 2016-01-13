require('./main.css');
require('bootstrap/dist/css/bootstrap.css');

import React from 'react'
import ReactDOM from 'react-dom'

import {applyMiddleware, compose, createStore, combineReducers} from 'redux';
import {Provider} from 'react-redux'
import {Router, Route} from 'react-router'
import createHistory from  'history/lib/createHashHistory'
import {syncReduxAndRouter, routeReducer} from  'redux-simple-router'


import i18next from 'i18next/lib';
import XHR from 'i18next-xhr-backend/lib';
import LanguageDetector from 'i18next-browser-languagedetector/lib';

import RootRoute from './components/router'
import {current_user} from './reducers'

const reducer = combineReducers(Object.assign({}, current_user, {
    routing: routeReducer
}));
const history = createHistory();
const finalCreateStore = compose(
)(createStore);
const store = finalCreateStore(reducer);



function main(options) {
    i18next
        .use(XHR)
        .use(LanguageDetector)
        .init({
            fallbackLng: 'en-US',
            backend: {
                loadPath: '/locales/{{lng}}'
            },
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: 'locale',
                lookupCookie: 'locale',
                lookupLocalStorage: 'locale',
                caches: ['localStorage', 'cookie']
            }
        }, (err, t)=> {
            syncReduxAndRouter(history, store);

            ReactDOM.render(<Provider store={store}>
                <Router history={history}>
                    {RootRoute}
                </Router>
            </Provider>, document.getElementById('root'));
        });
}

export default main