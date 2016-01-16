import {combineReducers} from 'redux';
import {syncHistory, routeReducer} from 'redux-simple-router'
import {browserHistory} from 'react-router'

import reducers from '../reducers';


export const reducer = combineReducers(Object.assign({}, reducers, {
    routing: routeReducer
}));

export const reduxRouterMiddleware = syncHistory(browserHistory);



