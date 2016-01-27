import React from 'react';
import {Route} from 'react-router';

import {Show as ShowArticle, Index as ListArticle} from './Article'
import {Show as ShowTag} from './Tag'
import Layout from './Layout'

const Router = (
    <Route path="cms" component={Layout}>
        <Route path="article/:aid" component={ShowArticle}/>
        <Route path="articles" component={ListArticle}/>

        <Route path="tag/:name" component={ShowTag}/>
    </Route>
);

export default Router