import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

const Widget = React.createClass({
    getInitialState() {

        return {};
    },
    componentDidMount(){
        console.log('mount roles');
    },
    render(){
        return (
            <p>
                roles
            </p>
        )
    }
});
export default Widget