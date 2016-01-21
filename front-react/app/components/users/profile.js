import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

function Profile({copyright}) {
    return (
        <footer>
            <p>
                profile
            </p>
        </footer>
    )

}

Profile.propTypes = {
};

export default connect(
    state => ({}),
    dispatch => ({})
)(Profile);