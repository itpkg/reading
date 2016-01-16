import {createStore, combineReducers, applyMiddleware} from 'redux';
import {syncHistory, routeReducer} from 'redux-simple-router'
import {browserHistory} from 'react-router'

import reducers from '../reducers';


const reducer = combineReducers(Object.assign({}, reducers, {
    routing: routeReducer
}));
const reduxRouterMiddleware = syncHistory(browserHistory);
const createStoreWithMiddleware = applyMiddleware(reduxRouterMiddleware)(createStore);
const store = createStoreWithMiddleware(reducer);

export default store;
