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


import {post} from '../../ajax'
import {checkResult} from './actions'
import Form from '../../components/Form'


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

const SignInW = ({}) => (
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

SignInW.propTypes = {
    onCheck: PropTypes.func.isRequired
};


export const SignIn = connect(
    state => ({}),
    dispatch => ({
        onCheck: function (rst) {
            dispatch(checkResult(rst));
        }
    })
)(SignInW);

export const SignUp = () =>(<div>
    <Form
        fields={[
            {
                id: 'email',
                type: 'text',
                label: i18n.t('platform.user.email'),
                value: '',
            },
            {
                id: 'password',
                type: 'password',
                label: i18n.t('platform.user.password'),
                value: '',
            },
            {
                id: 'password_confirmation',
                type: 'password',
                label: i18n.t('platform.user.password_confirmation'),
                value: '',
            },
            {
                id:'confirm_success_url',
                type:'hidden',
                value:`${process.env.CONFIG.host}/flash`,
            }
        ]}
        title={i18n.t('platform.auth.sign-up')}
        method="post"
        action="/auth"/>
    <br/>
    <Links/>
</div>);

//todo 暂时不支持重新发送激活邮件
export const Confirm = ({}) => (
    <div>
        <Form
            fields={[
                {
                    id: 'email',
                    type: 'text',
                    label: i18n.t('platform.user.email'),
                    value: '',
                },
                {
                    id:'confirm_success_url',
                    type:'hidden',
                    value:`${process.env.CONFIG.host}/flash`,
                }
            ]}
            title={i18n.t('platform.auth.confirm')}
            method="post"
            action="/auth/confirmation"/>
        <br/>
        <Links/>
    </div>
);

//todo
export const Unlock = ({}) => (
    <div>
        <Form
            fields={[
                {
                    id: 'email',
                    type: 'text',
                    label: i18n.t('platform.user.email'),
                    value: '',
                },
                {
                    id:'confirm_success_url',
                    type:'hidden',
                    value:`${process.env.CONFIG.host}/flash`,
                }
            ]}
            title={i18n.t('platform.auth.unlock')}
            method="post"
            action="/auth/unlock"/>
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

//todo
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






