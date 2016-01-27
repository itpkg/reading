import {combineReducers} from 'redux'
import {browserHistory} from 'react-router'
import {syncHistory, routeReducer} from 'react-router-redux'

import reducers from '../reducers'

export const middleware = syncHistory(browserHistory);
export const reducer = combineReducers(Object.assign({}, reducers, {
    routing: routeReducer
}));