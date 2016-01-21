import React from 'react';
import {Alert} from 'react-bootstrap'
import i18next from 'i18next/lib';

export default React.createClass({
    render(){
        return (<div className="col-md-offset-1 col-md-10">
            <br/>
            <Alert bsStyle="danger" onDismiss={this.handleAlertDismiss}>
                <strong>{i18next.t("no_match")}</strong>
            </Alert>
        </div>)
    }
});
