import React, {PropTypes} from 'react'
import {connect} from 'react-redux'

import usFlag from 'famfamfam-flags/dist/png/us.png'
import cnFlag from 'famfamfam-flags/dist/png/cn.png'

const Widget = ({info}) => (
    <div className="footer">
        {info.copyright}
        &nbsp;
        <a href="/?locale=en">
            <img src={usFlag} alt="English"/>
        </a>
        &nbsp;
        <a href="/?locale=zh-CN">
            <img src={cnFlag} alt="简体中文"/>
        </a>
    </div>
);

Widget.propTypes = {
    info: PropTypes.object.isRequired
};

export default connect(
    state => ({info: state.siteInfo})
)(Widget)

