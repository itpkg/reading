import React from 'react';
import {Button} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import i18next from 'i18next/lib';

import {GET} from '../../ajax'

import {Cloud as TagCloud} from './Tag'
import {Bar as Notices} from '../Notice'
import {LatestBar as LatestArticlesBar} from './Article'

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
                    <Notices size={3}/>
                    <br/>
                    <LatestArticlesBar size={8}/>
                    <br/>
                    <TagCloud/>
                </div>
            </div>
        )
    }
});

export default Layout;
