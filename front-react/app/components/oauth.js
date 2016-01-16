import React from 'react';
import {Alert} from 'react-bootstrap'
import {Route } from 'react-router'
import i18next from 'i18next/lib';
import parse from 'url-parse'

import {AjaxMixin} from '../mixins/ajax';

export const Google = React.createClass({
    mixins: [AjaxMixin],
    componentDidMount(){
        var query = parse(window.location.href, true).query;
        console.log(query);
        //return {ok: query.ok == 'true', message: query.msg}
        //this.GET('/oauth/sign_in', function (rst) {
        //    this.setState(rst);
        //});
    },
    render(){
        return (
            <div className="col-md-offset-1 col-md-10">
                <br/>
                <Alert bsStyle="success">
                    <strong>{i18next.t('messages.waiting')}</strong>
                </Alert>
            </div>
        )
    }
});