import {applyMiddleware, compose, createStore} from 'redux'

import {middleware, reducer} from './vars'

const finalCreateStore = compose(
    applyMiddleware(middleware)
)(createStore);
const store = finalCreateStore(reducer);

export default store;