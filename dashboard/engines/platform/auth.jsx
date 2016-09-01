import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {browserHistory} from 'react-router'

import {List, ListItem} from 'material-ui/List';

import ActionSignIn from 'material-ui/svg-icons/action/fingerprint';
import ActionUnlock from 'material-ui/svg-icons/action/lock-open';
import ActionSignUp from 'material-ui/svg-icons/action/account-box';
import ActionConfirm from 'material-ui/svg-icons/action/assignment-turned-in';
import ActionForgotPassword from 'material-ui/svg-icons/action/find-replace';
import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

import i18n from 'i18next'

export const LeaveWord = ({}) => (
    <fieldset className="form">
        <legend>{i18n.t('platform.auth.leave_a_message')}</legend>
        <TextField
            id='message'
            multiLine={true}
        />
        <br/>
        <FlatButton label={i18n.t("buttons.save")} primary={true}/>
    </fieldset>

);

const LinksW = ({onGoto}) => (
    <List>
        <ListItem onClick={()=>onGoto('/auth/sign-in')}
                  primaryText={i18n.t("platform.auth.sign-in")} leftIcon={<ActionSignIn />}/>
        <ListItem onClick={()=>onGoto('/auth/sign-up')}
                  primaryText={i18n.t("platform.auth.sign-up")} leftIcon={<ActionSignUp />}/>
        <ListItem onClick={()=>onGoto('/auth/confirm')}
                  primaryText={i18n.t("platform.auth.confirm")} leftIcon={<ActionConfirm />}/>
        <ListItem onClick={()=>onGoto('/auth/unlock')}
                  primaryText={i18n.t("platform.auth.unlock")} leftIcon={<ActionUnlock />}/>
        <ListItem onClick={()=>onGoto('/auth/forgot-password')}
                  primaryText={i18n.t("platform.auth.forgot-password")} leftIcon={<ActionForgotPassword />}/>
    </List>
);

LinksW.propTypes = {
    onGoto: PropTypes.func.isRequired
};

export const Links = connect(
    state => ({}),
    dispatch => ({
        onGoto: function (url) {
            browserHistory.push(url);
        }
    })
)(LinksW);

export const SignIn = ({}) => (
    <div>
        <fieldset className="form">
            <legend>{i18n.t('platform.auth.sign-in')}</legend>
            <TextField id="email"
                       floatingLabelText={i18n.t('platform.user.email')}
            />
            <TextField id="password"
                       floatingLabelText={i18n.t('platform.user.password')}
                       type="password"
            />
            <FlatButton label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>
        <br/>
        <Links/>
    </div>
);

export const SignUp = ({}) => (
    <div>
        <fieldset className="form">
            <legend>{i18n.t('platform.auth.sign-up')}</legend>
            <TextField id="email"
                       floatingLabelText={i18n.t('platform.user.email')}
            />
            <TextField id="password"
                       floatingLabelText={i18n.t('platform.user.password')}
                       type="password"
            />
            <TextField id="password_confirmation"
                       floatingLabelText={i18n.t('platform.user.password_confirmation')}
                       type="password"
            />
            <FlatButton label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>
        <br/>
        <Links/>
    </div>
);

export const Confirm = ({}) => (
    <div>
        <fieldset className="form">
            <legend>{i18n.t('platform.auth.confirm')}</legend>
            <TextField id="email"
                       floatingLabelText={i18n.t('platform.user.email')}
            />
            <FlatButton label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>
        <br/>
        <Links/>
    </div>
);

export const Unlock = ({}) => (
    <div>
        <fieldset className="form">
            <legend>{i18n.t('platform.auth.unlock')}</legend>
            <TextField id="email"
                       floatingLabelText={i18n.t('platform.user.email')}
            />
            <FlatButton label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>
        <br/>
        <Links/>
    </div>
);

//todo
export const Profile = ({}) => (
    <div>
        profile
    </div>
);


export const ForgotPassword = ({}) => (
    <div>
        <fieldset className="form">
            <legend>{i18n.t('platform.auth.forgot-password')}</legend>
            <TextField id="email"
                       floatingLabelText={i18n.t('platform.user.email')}
            />
            <FlatButton label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>
        <br/>
        <Links/>
    </div>
);






