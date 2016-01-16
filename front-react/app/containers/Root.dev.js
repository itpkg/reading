import React from 'react';
import { Provider } from 'react-redux';

import App from '../components/Router';
import DevTools from './DevTools';

const root = React.createClass({
    render(){
        return (
            <Provider store={this.props.store}>
                <div>
                    <App/>
                    <DevTools />
                </div>
            </Provider>
        );
    }
});

export default root;