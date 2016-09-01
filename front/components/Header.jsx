import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';
import AppsIcon from 'material-ui/svg-icons/navigation/apps';
import FlatButton from 'material-ui/FlatButton';


import {browserHistory} from 'react-router'
import {signOut} from '../engines/platform/actions'

import i18n from 'i18next'


const Widget = ({info, onGoto, onSignOut, user}) => (
    <AppBar
        title={info.title}
        onTitleTouchTap={()=>onGoto('/')}
        iconElementLeft={<IconButton onClick={()=>onGoto('/')}><AppsIcon /></IconButton>}
        iconElementRight={ user.uid ?
            <FlatButton label={i18n.t('platform.auth.sign-out')} onClick={onSignOut}/> :
            <FlatButton label={i18n.t('platform.auth.sign-in-or-up')} onClick={()=>onGoto('/auth/sign-in')}/>}
        /*
         iconElementRight={
         <IconMenu
         iconButtonElement={
         <IconButton><MoreVertIcon /></IconButton>
         }
         targetOrigin={{horizontal: 'right', vertical: 'top'}}
         anchorOrigin={{horizontal: 'right', vertical: 'top'}}
         >
         {(user.uid ?
         ["aaa", "bbb"] :
         [
         "sign-in",
         "sign-up",
         "confirm",
         "unlock",
         "forgot-password"
         ]).map(k => <MenuItem key={k} onClick={()=>onGoto(`/auth/${k}`)} primaryText={i18n.t(`platform.auth.${k}`)}/>)}
         </IconMenu>
         }
         */

    />
);

Widget.propTypes = {
    onSignOut: PropTypes.func.isRequired,
    onGoto: PropTypes.func.isRequired,
    info: PropTypes.object.isRequired,
    user: PropTypes.object.isRequired
};

export default connect(
    state => ({
        info: state.siteInfo,
        user: state.currentUser
    }),
    dispatch => ({
        onGoto: function (url) {
            browserHistory.push(url);
        },
        onSignOut: function (e) {
            dispatch(signOut);
        }
    })
)(Widget)

