import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {browserHistory} from 'react-router'

import TextField from 'material-ui/TextField';
import FlatButton from 'material-ui/FlatButton';

import i18n from 'i18next'

import {post} from '../../ajax'
import {checkResult} from './actions'

const NewFmW = React.createClass({
    getInitialState: function () {
        return {
            content: ''
        };
    },
    handleChange: function (e) {
        var tmp = this.state;
        tmp[e.target.id] = e.target.value;
        this.setState(tmp);
    },
    handleSubmit: function (e) {
        e.preventDefault();
        const {onCheck} = this.props;
        var data = new FormData();
        data.append( "content", this.state.content );
        post('/leave_words', data, function (rst) {
            onCheck(rst);
            if(!rst.errors){
                this.setState({content:''});
            }

        }.bind(this));
    },
    render: function () {
        return (<fieldset className="form">
            <legend>{i18n.t('platform.auth.leave_a_message')}</legend>
            <TextField
                id='content'
                value={this.state.content}
                onChange={this.handleChange}
                multiLine={true}
            />
            <br/>
            <FlatButton onClick={this.handleSubmit} label={i18n.t("buttons.save")} primary={true}/>
        </fieldset>)
    }
});


NewFmW.propTypes = {
    onCheck: PropTypes.func.isRequired
};


export const NewFm = connect(
    state => ({}),
    dispatch => ({
        onCheck: function (rst) {
            dispatch(checkResult(rst));
        }
    })
)(NewFmW);