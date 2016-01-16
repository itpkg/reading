console.log('ENV: development');

import {hashHistory} from 'react-router'
import Main from './main'
Main({
    apiUrl: '/',
    history: hashHistory
});

