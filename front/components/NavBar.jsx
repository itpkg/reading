import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import Divider from 'material-ui/Divider';

import i18n from 'i18next'


const Widget = ({onClose, open}) => (

    <Drawer docked={false} open={open}>
        <MenuItem onTouchTap={onClose}>Menu Item</MenuItem>
        <Divider />
        <MenuItem onTouchTap={onClose}>Menu Item 2</MenuItem>
    </Drawer>
);

Widget.propTypes = {
    onClose: PropTypes.func.isRequired,
    open: PropTypes.bool.isRequired
};
export default connect(
    state => ({open: false}),
    dispatch => ({
        onClose: function () {
            console.log("on close");
        }
    })
)(Widget)

