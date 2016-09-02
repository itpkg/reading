import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

import {Links as NonSignInLinks} from '../engines/platform/auth'

import {NewFm as LeaveWord} from '../engines/platform/leave_words'

const Widget = (user) => user.id ?
    (<div>index</div>) :
    (<div>
        <LeaveWord/>
        <br/>
        <NonSignInLinks/>
    </div>);


Widget.propTypes = {
    user: PropTypes.object.isRequired
};

export default connect(
    state=>({user: state.currentUser}),
    dispatch => ({})
)(Widget);
