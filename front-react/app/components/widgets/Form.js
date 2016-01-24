import React, {PropTypes} from 'react';
import {Input, Button, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState(){
        return {}
    },
    componentDidMount() {
        const {fields} = this.props;
        var val = {};
        fields.forEach(function (fld) {
            val[fld.id] = fld.value;
        });
        this.setState(val);
    },
    handleInputChange(e){
        var val = {};
        val[e.target.id] = e.target.value;
        this.setState(val);
    },
    handleSubmit(e){
        e.preventDefault();
        console.log(this.props.fields);
        console.log(this.state);
        //console.log($(this).serialize());
        console.log(e);
    },
    render(){
        const {title, method, action, fields} = this.props;

        var handler = this.handleInputChange;
        var link = this.linkState;

        return (
            <div>
                {title ? (<div><h3>{title}</h3>
                    <hr/>
                </div>) : <br/>}
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
                            <Button type="reset">{i18next.t('buttons.reset')}</Button>
                        </ButtonToolbar>
                    </div>
                </form>
            </div>
        )
    }
});
export default Widget;