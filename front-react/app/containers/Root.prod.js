import React from 'react';
import { Provider } from 'react-redux';

import App from '../components/Router';

const root = React.createClass({
    render(){
        return (
            <Provider store={this.props.store}>
                <App/>
            </Provider>
        );
    }
});

export default root;