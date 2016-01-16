import React from 'react';
import {Router, Route, Link, IndexRoute,browserHistory} from 'react-router'

import {Layout} from './widgets'
import {Index, NoMatch, Message} from './home'
import {Article} from './cms'
import {Google} from './oauth'


const router = React.createClass({
    render(){
        return (
            <Router history={browserHistory}>
                <Route path="/" component={Layout}>
                    <IndexRoute component={Index}/>

                    <Route path="home" component={Index}/>
                    <Route path="oauth/google" component={Google}/>

                    <Route path="message" component={Message}/>
                    <Route path="cms/articles/:aid" component={Article}/>

                    <Route path="*" component={NoMatch}/>
                </Route>
            </Router>
        );
    }
});

export default router
