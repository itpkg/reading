import React, {PropTypes} from 'react';
import {Button} from 'react-bootstrap'
import i18next from 'i18next/lib';

import {DELETE} from '../../ajax'

const Widget = React.createClass({
    handleSubmit(e){
        e.preventDefault();
        if (confirm(i18next.t("messages.are_you_sure"))) {
            const {action, onRefresh} = this.props;
            DELETE(
                action,
                function (rst) {
                    alert(i18next.t('messages.' + (rst.ok ? 'success' : 'failed')) + ": " + (rst.messages ? rst.messages.join('\n') : ""));
                    if (onRefresh) {
                        onRefresh();
                    }
                },
                function (e) {
                    alert(e.responseText);
                }
            )
        }
    },
    render(){
        const {size} = this.props;
        return (
            <Button bsStyle="danger" bsSize={size} onClick={this.handleSubmit}>{i18next.t('buttons.remove')}</Button>)
    }
});

export default Widget;