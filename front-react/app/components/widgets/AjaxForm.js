import React, {PropTypes} from 'react';
import {Input, Button, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

import {GET} from '../../ajax'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState(){
        return {
            _form_: {
                fields: []
            }
        }
    },
    //handleReset(){
    //    var val = {};
    //    this.state._form_.fields.forEach(function(fld){
    //        val[fld.id] = fld.value;
    //    });
    //    this.setState(val);
    //},
    handleSubmit(e){
        e.preventDefault();
        console.log(this.state);
    },
    componentDidMount() {
        const {url} = this.props;

        GET(url, function (fm) {
            var val = {_form_: fm};
            fm.fields.forEach(function (fld) {
                val[fld.id] = fld.value;
            });
            this.setState(val);
        }.bind(this))
    },
    render(){
        var fm = this.state._form_;
        return (
            <div>
                <br/>
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
                                    onClick={this.componentDidMount.bind(this)}>{i18next.t('buttons.reset')}</Button>
                        </ButtonToolbar>
                    </div>
                </form>
            </div>
        )
    }
});

export default Widget;