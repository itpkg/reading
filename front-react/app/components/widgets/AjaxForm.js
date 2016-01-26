import React, {PropTypes} from 'react';
import {Input, Button, ButtonToolbar, Alert} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

import {GET,POST} from '../../ajax'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState(){
        return {
            _form_: {
                fields: []
            },
            _alert_: {}
        }
    },
    handleSubmit(e){
        e.preventDefault();
        var data = Object.assign({}, this.state);
        delete data._form_;
        delete data._alert_;

        var fm = this.state._form_;
        switch (fm.method) {
            default:
                POST(
                    fm.action,
                    data, function (rst) {
                        this.setState({
                            _alert_: {
                                show: true,
                                style: rst.ok ? 'success' : 'danger',
                                head: 'messages.' + (rst.ok ? 'success' : 'failed'),
                                body: rst.messages ? rst.messages : []
                            }
                        });
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
    handleReset(){
        const {url} = this.props;

        GET(url, function (fm) {
            var val = {_form_: fm};
            fm.fields.forEach(function (fld) {
                val[fld.id] = fld.value;
            });
            this.setState(val);
        }.bind(this))

    },
    componentDidMount() {
        this.handleReset();
    },
    handleAlertDismiss(){
        this.setState({_alert_: {}});
    },
    render(){
        var fm = this.state._form_;
        var alt = this.state._alert_;
        return (
            <div>
                {alt.show ? (<Alert bsStyle={alt.style} onDismiss={this.handleAlertDismiss}>
                    <h4>{i18next.t(alt.head)}[{new Date().toLocaleString()}]</h4>
                    <ul>
                        {alt.body.map(function (itm, idx) {
                            return <li key={idx}>{itm}</li>
                        })}
                    </ul>
                </Alert>) : <br/>}
                <form method={fm.method ? fm.method : 'POST'} action={fm.action} onSubmit={this.handleSubmit}>
                    {
                        fm.fields.map(function (fld, idx) {
                            var label = i18next.t("forms." + fm.id + "." + fld.id);
                            switch (fld.type) {
                                case "email":
                                case "text":
                                case "password":
                                case "textarea":
                                    return <Input key={idx} id={fld.id}
                                                  type={fld.type} label={label}
                                                  valueLink={this.linkState(fld.id)}
                                                  placeholder={fld.placeholder}/>;
                                default:
                                    console.log("unknown type" + fld.type);
                                    return <Input key={idx} type="hidden"/>
                            }
                        }.bind(this))
                    }
                    <div className="form-group">
                        <ButtonToolbar>
                            <Button type="submit" bsStyle="primary">{i18next.t('buttons.submit')}</Button>
                            <Button type="reset"
                                    onClick={this.handleReset}>{i18next.t('buttons.reset')}</Button>
                        </ButtonToolbar>
                    </div>
                </form>
            </div>
        )
    }
});

export default Widget;