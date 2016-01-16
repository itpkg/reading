import {createStore, applyMiddleware} from 'redux'

import {reducer, reduxRouterMiddleware} from './vars'

const createStoreWithMiddleware = applyMiddleware(reduxRouterMiddleware)(createStore);
const store = createStoreWithMiddleware(reducer);

export default store
