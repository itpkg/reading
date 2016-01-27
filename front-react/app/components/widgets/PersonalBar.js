import React, {Component, PropTypes} from 'react';
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {connect} from 'react-redux';
import {routeActions} from 'redux-simple-router'
import i18next from 'i18next/lib';

import {GET} from '../../ajax'
import {refresh} from '../../actions/auth'
import {signOut} from '../../actions/user'
import {CurrentUser} from '../../mixins'

const PersonalBar = React.createClass({
    mixins: [CurrentUser],
    componentDidMount() {
        const {onOauth} = this.props;
        onOauth();
    },
    render() {
        const {user,oauth, onSignOut} = this.props;
        if (this.isSignIn()) {
            return (
                <NavDropdown eventKey={'welcome'} title={i18next.t("users.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <LinkContainer to={'/cms/article/'+Math.random().toString(36).substring(2)}>
                        <MenuItem eventKey={'article.new'}>{i18next.t("dashboard.article")}</MenuItem>
                    </LinkContainer>
                    <LinkContainer to={'/dashboard/personal/logs'}>
                        <MenuItem eventKey={'dashboard'}>{i18next.t("dashboard.index")}</MenuItem>
                    </LinkContainer>
                    <MenuItem divider/>
                    <MenuItem eventKey={'sign_out'} onClick={() => onSignOut()}>{i18next.t("users.sign_out")}</MenuItem>
                </NavDropdown>
            )
        } else {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.sign_up_or_in")} id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1}
                              href={oauth.google}>
                        {i18next.t("users.sign_in_with.google")}
                    </MenuItem>
                </NavDropdown>
            )
        }
    }
});


PersonalBar.propTypes = {
    user: PropTypes.object.isRequired,
    oauth: PropTypes.object.isRequired,
    onOauth: PropTypes.func.isRequired,
    onSignOut: PropTypes.func.isRequired
};


export default connect(
    state => ({
        user: state.current_user,
        oauth: state.oauth
    }),
    dispatch => ({
        onOauth: function () {
            GET('/oauth/sign_in', function (rst) {
                dispatch(refresh(rst));
            })
        },
        onSignOut: function () {
            dispatch(signOut());
            dispatch(routeActions.push('/home'));
        }
    })
)(PersonalBar);
