import React from 'react';
import {Alert} from 'react-bootstrap'
import {Link} from 'react-router'
import parse from 'url-parse'
import i18next from 'i18next/lib'

export const Index = React.createClass({
    render(){ //todo
        return (<div>
            index
            <Link to={`/home`}>Home</Link>
            <br/>
            <Link to={`/cms/articles/about`}>about</Link>
            <br/>
            <Link to={`/cms/articles/faq`}>about</Link>
        </div>)
    }
});


export const NoMatch = React.createClass({
    render(){
        return (<div className="col-md-offset-1 col-md-10">
            <br/>
            <Alert bsStyle="danger" onDismiss={this.handleAlertDismiss}>
                <strong>{i18next.t("no_match")}</strong>
            </Alert>
        </div>)
    }
});


export const Message = React.createClass({
    getInitialState: function () {
        var query = parse(window.location.href, true).query;
        return {ok: query.ok == 'true', message: query.msg}
    },
    render(){
        return (
            <div className="col-md-offset-1 col-md-10">
                <Alert bsStyle={ this.state.ok? "success":"danger"}>
                    <strong>{this.state.message}</strong>
                </Alert>
            </div>
        )
    }
});
