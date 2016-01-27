import React, {PropTypes} from 'react';
import {Alert, Input, Button, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';


const Widget = React.createClass({
    render(){
        const {onReset, onDelete} = this.props;
        return (
            <div className="form-group">
                <ButtonToolbar>
                    <Button type="submit" bsStyle="primary">{i18next.t('buttons.submit')}</Button>
                    <Button onClick={onReset}>{i18next.t("buttons.reset")}</Button>
                    {
                        onDelete ?
                            <Button bsStyle="danger" onClick={onDelete}>{i18next.t("buttons.remove")}</Button> :
                            ""
                    }
                </ButtonToolbar>
            </div>
        )
    }
});

export default Widget;