import React from 'react';
import {Router, Route, IndexRoute, browserHistory} from 'react-router';

import CmsArticle from './cms/Article'
import Layout from './Layout'
import Home from './Home'
import {Index as Notices} from './Notice'
import NoMatch from './NoMatch'
import OauthGoogle from './oauth/google'

import Dashboard from './dashboard/Router'

export default React.createClass({
    render: function () {
        return (
            <Router history={browserHistory}>
                <Route path="/" component={Layout}>


                    <Route path="cms/articles/:aid" component={CmsArticle}/>


                    <Route path="notices" component={Notices}/>
                    {Dashboard}
                    <Route path="oauth/google" component={OauthGoogle}/>
                    <Route path="*" component={NoMatch}/>
                    <IndexRoute component={Home}/>
                </Route>
            </Router>
        );
    }
});