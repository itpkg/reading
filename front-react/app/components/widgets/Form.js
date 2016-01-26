import React, {PropTypes} from 'react';
import {Alert, Input, Button, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

import {POST} from  '../../ajax'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState(){
        return {
            _alert_: {}
        }
    },
    handleReset(){
        const {fields} = this.props;
        var val = {};
        fields.forEach(function (fld) {
            val[fld.id] = fld.value;
        }.bind(this));
        this.setState(val);
    },
    componentDidMount() {
        this.handleReset();
    },
    handleAlertDismiss(){
        this.setState({_alert_: {}});
    },
    handleSubmit(e){
        e.preventDefault();
        var data = Object.assign({}, this.state);
        delete data._alert_;

        const {method, action, onSubmit} = this.props;
        switch (method) {
            default:
                POST(
                    action,
                    data, function (rst) {
                        this.setState({
                            _alert_: {
                                show: true,
                                style: rst.ok ? 'success' : 'danger',
                                head: 'messages.' + (rst.ok ? 'success' : 'failed'),
                                body: rst.messages ? rst.messages : []
                            }
                        });
                        if (onSubmit) {
                            onSubmit();
                        }
                    }.bind(this),
                    function (e) {
                        this.setState({
                            _alert_: {
                                show: true,
                                style: 'danger',
                                head: 'messages.failed',
                                body: [e.responseText]
                            }
                        });
                    }.bind(this)
                );
                break;
        }
    },
    render(){
        const {title, method, action, fields} = this.props;
        var link = this.linkState;
        var alt = this.state._alert_;
        return (
            <div>
                {title ? (<div><h3>{title}</h3>
                    <hr/>
                </div>) : <br/>}
                {alt.show ? (<Alert bsStyle={alt.style} onDismiss={this.handleAlertDismiss}>
                    <h4>{i18next.t(alt.head)}[{new Date().toLocaleString()}]</h4>
                    <ul>
                        {alt.body.map(function (itm, idx) {
                            return <li key={idx}>{itm}</li>
                        })}
                    </ul>
                </Alert>) : <br/>}
                <form method={method ? method : 'POST'} action={action} onSubmit={this.handleSubmit}>
                    {
                        fields.map(function (fld) {
                            switch (fld.type) {
                                case "email":
                                case "text":
                                case "password":
                                case "textarea":
                                    return <Input key={fld.id} id={fld.id}
                                                  type={fld.type} label={fld.label}
                                                  valueLink={link(fld.id)}
                                                  placeholder={fld.placeholder}/>;
                                default:
                                    console.log("unknown type" + fld.type);
                                    return <Input key={idx} type="hidden"/>
                            }
                        })
                    }
                    <div className="form-group">
                        <ButtonToolbar>
                            <Button type="submit" bsStyle="primary">{i18next.t('buttons.submit')}</Button>
                            <Button onClick={this.handleReset}>{i18next.t('buttons.reset')}</Button>
                        </ButtonToolbar>
                    </div>
                </form>
            </div>
        )
    }
});
export default Widget;