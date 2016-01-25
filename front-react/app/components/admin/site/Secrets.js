import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

import AjaxForm from '../../widgets/AjaxForm'

const Widget = React.createClass({
    render(){
        return (<AjaxForm url="/admin/site/secrets"/>)
    }
});
export default Widget;