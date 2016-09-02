import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import i18n from 'i18next'
import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

import {call} from '../ajax'
import {checkResult} from '../engines/platform/actions'


const Widget = React.createClass({
    getInitialState: function () {
        const {fields} = this.props;
        return fields.reduce(function (obj, fld) {
            obj[fld.id] = fld.value;
            return obj;
        }, {});

    },
    handleChange: function (e) {
        var tmp = this.state;
        tmp[e.target.id] = e.target.value;
        this.setState(tmp);
    },
    handleSubmit: function (e) {
        e.preventDefault();
        const {method, action, onCheck, fields} = this.props;
        var data = new FormData();
        fields.forEach(function (fld) {
            data.append(fld.id, this.state[fld.id]);
        }.bind(this));

        call(method, action, data, function (rst) {
            onCheck(rst);
            if (!rst.errors) {
                var tmp = fields.reduce(function (obj, fld) {
                    obj[fld.id] = fld.value;
                    return obj;
                }, {});
                this.setState(tmp);
            }
        }.bind(this));
    },
    render: function () {
        const {title, fields} = this.props;
        return (            <fieldset className="form">
            <legend>{title}</legend>
            {
                fields.map(function (fld) {
                    switch (fld.type) {
                        case "text":
                            return (<TextField id={fld.id} key={fld.id}
                                               value={this.state[fld.id]}
                                               onChange={this.handleChange}
                                               floatingLabelText={fld.label}
                            />);
                            break;
                        case "password":
                            return (<TextField id={fld.id} key={fld.id}
                                               value={this.state[fld.id]}
                                               onChange={this.handleChange}
                                               floatingLabelText={fld.label}
                                               type="password"
                            />);
                            break;
                        case "textarea":
                            return (<TextField id={fld.id} key={fld.id}
                                               value={this.state[fld.id]}
                                               onChange={this.handleChange}
                                               floatingLabelText={fld.label}
                                               multiLine={true}
                            />);
                            break;
                        default:
                            return (<input type="hidden"
                                           key={fld.id} id={fld.id}
                                           value={this.state[fld.id]}/>);
                    }
                }.bind(this))
            }
            <FlatButton onClick={this.handleSubmit} label={i18n.t("buttons.submit")} primary={true}/>
        </fieldset>            )
    }
});


Widget.propTypes = {
    title: PropTypes.string.isRequired,
    method: PropTypes.string.isRequired,
    action: PropTypes.string.isRequired,
    fields: PropTypes.array.isRequired,
    onCheck: PropTypes.func.isRequired
};


export default connect(
    state => ({}),
    dispatch => ({
        onCheck: function (rst) {
            dispatch(checkResult(rst));
        }
    })
)(Widget);