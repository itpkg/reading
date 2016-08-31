import React from 'react'
import {createStore, combineReducers} from 'redux'
import {Provider} from 'react-redux'
import {Router, IndexRoute, Route, browserHistory} from 'react-router'
import {syncHistoryWithStore, routerReducer} from 'react-router-redux'

import root from './engines'
import Layout from './components/Layout'
import NoMatch from './components/NoMatch'
import Index from './components/Index'


const reducers = root.reducers();
const store = createStore(
    combineReducers({
        ...reducers,
        routing: routerReducer
    })
);
const history = syncHistoryWithStore(browserHistory, store);

const Widget = () => (
    <Provider store={store}>
        <Router history={history}>
            <Route path="/" component={Layout}>
                <IndexRoute component={Index}/>
                {root.routes()}
                <Route path="*" component={NoMatch}/>
            </Route>
        </Router>
    </Provider>
);

export default Widget