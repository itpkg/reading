import {createStore, compose, combineReducers, applyMiddleware} from 'redux';
import {syncHistory, routeReducer} from 'redux-simple-router'
import {browserHistory} from 'react-router'

import reducers from '../reducers';
import DevTools from '../containers/DevTools';


const reducer = combineReducers(Object.assign({}, reducers, {
    routing: routeReducer
}));
const reduxRouterMiddleware = syncHistory(browserHistory);
const createStoreWithMiddleware = compose(
    applyMiddleware(reduxRouterMiddleware),
    DevTools.instrument() //dev-tool
)(createStore);
const store = createStoreWithMiddleware(reducer);

reduxRouterMiddleware.listenForReplays(store); //dev-tool

export default store;
