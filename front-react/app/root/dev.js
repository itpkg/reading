import React from 'react';
import {Provider} from 'react-redux'

import Router from '../components/Router'
import DevTools from './DevTools'

export default React.createClass({
    render: function () {
        return (
            <Provider store={this.props.store}>
                <div>
                    <Router/>
                    <DevTools />
                </div>
            </Provider>
        );
    }
});