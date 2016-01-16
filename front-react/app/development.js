console.log('IT-PACKAGE Reading: ' + process.env.NODE_ENV + '(' + VERSION + ')');

import {hashHistory} from 'react-router'
import Main from './main'
Main({
    history: hashHistory
});

