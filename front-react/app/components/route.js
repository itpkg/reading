import React from 'react';
import {Route, Link, IndexRoute } from 'react-router'

import {Layout} from './widgets'
import {Index, NoMatch, Message} from './home'
import {Article} from './cms'


const route = (
    <Route path="/" component={Layout}>
        <IndexRoute component={Index}/>

        <Route path="home" component={Index}/>
        <Route path="message" component={Message}/>
        <Route path="/cms/articles/:aid" component={Article}/>

        <Route path="*" component={NoMatch}/>
    </Route>
);

export default route