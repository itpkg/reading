import React from 'react';
import {Button} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import i18next from 'i18next/lib';
import $ from 'jquery';

const Layout = React.createClass({
    render() {
        const {children} = this.props;
        return (
            <div className="row">
                <br/>
                <div className="col-md-9">
                    {children}
                </div>
                <div className="col-md-3">
                </div>
            </div>
        )
    }
});

export default Layout;
