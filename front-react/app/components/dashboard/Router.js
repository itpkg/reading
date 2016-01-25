import React,{PropTypes} from 'react';
import {Route} from 'react-router';

import UsersLogs from '../users/Logs'
import SiteInfo from '../admin/site/Info'
import SiteSeo from '../admin/site/Seo'
import SiteSecrets from '../admin/site/Secrets'
import Roles from '../admin/Roles'
import Layout from './Layout'

const Router = (
    <Route path="dashboard" component={Layout}>
        <Route path="admin/site/info" component={SiteInfo}/>
        <Route path="admin/site/seo" component={SiteSeo}/>
        <Route path="admin/site/secrets" component={SiteSecrets}/>
        <Route path="admin/roles" component={Roles}/>
        <Route path="users/logs" component={UsersLogs}/>
    </Route>
);

export default Router