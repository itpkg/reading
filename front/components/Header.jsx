import React, {PropTypes} from 'react'
import { connect } from 'react-redux'

import AppBar from 'material-ui/AppBar';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';

import i18n from 'i18next'


const Widget = ({info, onToggle}) => (
    <AppBar
        title={info.title}
        onLeftIconButtonTouchTap={onToggle}
        iconElementRight={
            <IconMenu
                iconButtonElement={
                    <IconButton><MoreVertIcon /></IconButton>
                }
                targetOrigin={{horizontal: 'right', vertical: 'top'}}
                anchorOrigin={{horizontal: 'right', vertical: 'top'}}
            >
                <MenuItem primaryText="Refresh"/>
                <MenuItem primaryText="Help"/>
                <MenuItem primaryText="Sign out"/>
            </IconMenu>
        }
    />
);

Widget.propTypes = {
    onToggle: PropTypes.func.isRequired,
    info: PropTypes.object.isRequired
};

export default connect(
    state => ({ info: state.siteInfo }),
    dispatch => ({
        onToggle: function(){
            console.log("on toggle");
        }
    })
)(Widget)

