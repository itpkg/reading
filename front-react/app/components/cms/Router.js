import React from 'react';
import {Route} from 'react-router';

import {Show as ShowArticle, Index as ListArticle, Edit as EditArticle} from './Article'
import {Show as ShowTag, Index as ListTag} from './Tag'
import Layout from './Layout'

const Router = (
    <Route path="cms" component={Layout}>
        <Route path="article/:id" component={ShowArticle}/>
        <Route path="article/:id/edit" component={EditArticle}/>
        <Route path="articles" component={ListArticle}/>

        <Route path="tag/:id" component={ShowTag}/>
        <Route path="tags" component={ListTag}/>
    </Route>
);

export default Router