import React from 'react';
import {Router, Route, IndexRoute, browserHistory} from 'react-router';


import Layout from './Layout'
import Home from './Home'
import {Index as Notices} from './Notice'
import NoMatch from './NoMatch'
import OauthGoogle from './oauth/google'

import Dashboard from './dashboard/Router'
import Cms from './cms/Router'

export default React.createClass({
    render: function () {
        return (
            <Router history={browserHistory}>
                <Route path="/" component={Layout}>
                    {Cms}
                    {Dashboard}
                    <Route path="notices" component={Notices}/>
                    <Route path="oauth/google" component={OauthGoogle}/>
                    <Route path="home" component={Home}/>
                    <Route path="*" component={NoMatch}/>
                    <IndexRoute component={Home}/>
                </Route>
            </Router>
        );
    }
});