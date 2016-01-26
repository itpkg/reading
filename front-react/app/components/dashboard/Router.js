import React,{PropTypes} from 'react';
import {Route} from 'react-router';

import PersonalLogs from '../personal/Logs'
import SiteInfo from '../admin/site/Info'
import SiteSeo from '../admin/site/Seo'
import SiteSecrets from '../admin/site/Secrets'
import Notices from '../admin/Notices'
import Locales from '../admin/Locales'
import Users from '../admin/Users'
import Attachments from '../Attachments'
import Layout from './Layout'

const Router = (
    <Route path="dashboard" component={Layout}>
        <Route path="admin/site/info" component={SiteInfo}/>
        <Route path="admin/site/seo" component={SiteSeo}/>
        <Route path="admin/site/secrets" component={SiteSecrets}/>
        <Route path="admin/notices" component={Notices}/>
        <Route path="admin/locales" component={Locales}/>
        <Route path="admin/users" component={Users}/>
        <Route path="personal/logs" component={PersonalLogs}/>
        <Route path="attachments" component={Attachments}/>
    </Route>
);

export default Router