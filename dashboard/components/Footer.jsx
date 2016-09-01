import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

const Widget = ({info}) => (
    <div className="footer">
        {info.copyright}
    </div>
);

Widget.propTypes = {
    info: PropTypes.object.isRequired
};

export default connect(
    state => ({info: state.siteInfo})
)(Widget)

