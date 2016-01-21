import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux';
import {Alert} from 'react-bootstrap'
import {routeActions} from 'redux-simple-router'
import i18next from 'i18next/lib';
import parse from 'url-parse'

import {POST} from '../../ajax'
import {signIn} from '../../actions/user'


class Google extends Component {
    componentDidMount() {
        const {onSignIn} = this.props;
        onSignIn();
    }

    render() {
        return (<div className="col-md-offset-1 col-md-10">
            <br/>
            <Alert bsStyle="success">
                <strong>{i18next.t("messages.waiting")}</strong>{new Date().toLocaleString()}
            </Alert>
        </div>)
    }
}


Google.propTypes = {
    onSignIn: PropTypes.func.isRequired
};

export default connect(
    state => ({}),
    dispatch => ({
        onSignIn: function () {
            POST(
                '/oauth/sign_in',
                {
                    type: 'google',
                    code: parse(location.href, true).query.code
                },
                function (tkn) {
                    dispatch(signIn(tkn));
                    dispatch(routeActions.push('/home'))
                })
        }
    })
)(Google);

