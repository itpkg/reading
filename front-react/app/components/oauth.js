import React from 'react';
import {Alert} from 'react-bootstrap'
import {Route } from 'react-router'

export const Google = React.createClass({
    render(){
        return (
            <Alert bsStyle="warning">
                <strong>Holy guacamole!</strong> Best check yo self, you're not looking too good.
            </Alert>
        )
    }
});