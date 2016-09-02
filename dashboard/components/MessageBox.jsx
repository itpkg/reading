import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

import Snackbar from 'material-ui/Snackbar';

import {showMessage} from '../engines/platform/actions'

const Widget = ({box, onClose}) => (
    <Snackbar
        open={box.show}
        message={box.message}
        autoHideDuration={10000}
        onRequestClose={onClose}
    />
);

Widget.propTypes = {
    box: PropTypes.object.isRequired,
    onClose: PropTypes.func.isRequired
};

export default connect(
    state => ({box: state.messageBox}),
    dispatch => ({
        onClose: function (e) {
            dispatch(showMessage());
        }
    })
)(Widget)

