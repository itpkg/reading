import React from 'react';
import {Provider} from 'react-redux'

import Router from '../components/Router'

export default React.createClass({
    render: function () {
        return (
            <Provider store={this.props.store}>
                <Router/>
            </Provider>
        );
    }
});