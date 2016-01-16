import {createStore, compose, applyMiddleware} from 'redux'

import {reducer, reduxRouterMiddleware} from './vars'
import DevTools from '../containers/DevTools';

const createStoreWithMiddleware = compose(
    applyMiddleware(reduxRouterMiddleware),
    DevTools.instrument() //dev-tools
)(createStore);

const store = createStoreWithMiddleware(reducer);
reduxRouterMiddleware.listenForReplays(store); //dev-tools

export default store;
