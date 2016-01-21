import React, {Component, PropTypes} from 'react';
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {connect} from 'react-redux';
import {routeActions} from 'redux-simple-router'
import i18next from 'i18next/lib';
import $ from 'jquery';

import {GET} from '../../ajax'
import {refresh} from '../../actions/auth'
import {signOut} from '../../actions/user'


class PersonalBar extends Component {
    componentDidMount() {
        const {onOauth} = this.props;
        onOauth();
    }

    render() {
        const {user,oauth, onSignOut} = this.props;
        if ($.isEmptyObject(user)) {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.sign_up_or_in")} id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1}
                              href={oauth.google}>
                        {i18next.t("users.sign_in_with.google")}
                    </MenuItem>
                </NavDropdown>
            )
        } else {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <LinkContainer to={'/users/profile'}>
                        <MenuItem eventKey={3.1}>{i18next.t("users.profile")}</MenuItem>
                    </LinkContainer>
                    <MenuItem divider/>
                    <MenuItem eventKey={3.3} onClick={() => onSignOut()}>{i18next.t("users.sign_out")}</MenuItem>
                </NavDropdown>
            )
        }
    }
}


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
