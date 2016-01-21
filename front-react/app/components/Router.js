import React from 'react';
import {Router, Route, IndexRoute, browserHistory} from 'react-router';

import CmsArticle from './cms/Article'
import Layout from './Layout'
import Home from './Home'
import Notice from './Notice'
import NoMatch from './NoMatch'
import OauthGoogle from './oauth/google'
import UsersProfile from './users/profile'

export default React.createClass({
    render: function () {
        return (<Router history={browserHistory}>
            <Route path="/" component={Layout}>
                <IndexRoute component={Home}/>
                <Route path="oauth/google" component={OauthGoogle}/>
                <Route path="users/profile" component={UsersProfile}/>

                <Route path="notices/:id" component={Notice}/>
                <Route path="cms/articles/:aid" component={CmsArticle}/>

                <Route path="*" component={NoMatch}/>
            </Route>
        </Router>);
    }
});