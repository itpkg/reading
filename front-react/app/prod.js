import {browserHistory} from 'react-router'

import Main from './main'

Main({
    apiUrl: '/api/v1',
    history: browserHistory
});