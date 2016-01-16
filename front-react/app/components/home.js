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
                <h4>{i18next.t("no_match")}</h4>
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
        return (<div className="row">
                <div className="col-md-offset-1 col-md-10">
                    <Alert bsStyle={ this.state.ok? "success":"danger"}>
                        <h4>{this.state.message}</h4>
                    </Alert>
                </div>
            </div>
        )
    }
});
