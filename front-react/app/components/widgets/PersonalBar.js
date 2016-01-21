import React, {Component, PropTypes} from 'react';
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import $ from 'jquery';

import {GET} from '../../ajax'
import {refresh} from '../../actions/auth'


class PersonalBar extends Component {
    componentDidMount() {
        const {user, onOauth} = this.props;
        if ($.isEmptyObject(user)) {
            onOauth();
        }
    }

    render() {
        const {user,oauth} = this.props;
        if ($.isEmptyObject(user)) {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.sign_up_or_in")} id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1}
                              href={oauth.google}>{i18next.t("users.sign_in_with.google")}</MenuItem>
                </NavDropdown>
            )
        } else {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1} href="/#/personal/profile">{i18next.t("users.titles.profile")}</MenuItem>
                    <MenuItem divider/>
                    <MenuItem eventKey={3.3}>{i18next.t("users.titles.sign_out")}</MenuItem>
                </NavDropdown>
            )
        }
    }
}


PersonalBar.propTypes = {
    user: PropTypes.object.isRequired,
    oauth: PropTypes.object.isRequired,
    onOauth: PropTypes.func.isRequired
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
        }
    })
)(PersonalBar);
