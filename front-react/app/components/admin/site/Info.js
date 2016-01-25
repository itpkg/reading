import React, {PropTypes} from 'react';
import i18next from 'i18next/lib';

import AjaxForm from '../../widgets/AjaxForm'

const Widget = React.createClass({
    render(){
        return (<AjaxForm url="/admin/site/info"/>)
    }
});
export default Widget;