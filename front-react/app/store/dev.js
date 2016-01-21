import {applyMiddleware, compose, createStore} from 'redux'

import DevTools from '../root/DevTools'
import {middleware, reducer} from './vars'

const finalCreateStore = compose(
    applyMiddleware(middleware),
    DevTools.instrument() //dev
)(createStore);
const store = finalCreateStore(reducer);
middleware.listenForReplays(store); //dev

export default store;