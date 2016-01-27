require('react-select/dist/react-select.css');

import React,{PropTypes} from 'react';
import {connect} from 'react-redux';
import SelectInput from 'react-select'
import i18next from 'i18next/lib';

import {GET} from '../../ajax'

//todo
export const Select = React.createClass({
    getInitialState() {
        return {
            options: [
                {value: 1, label: 'One'},
                {value: 2, label: 'Two'},
                {value: 11, label: 'a1'},
                {value: 12, label: 'a2'},
                {value: 13, label: 'a3'},
                {value: 14, label: 'a4'}
            ]
        }
    },
    getOptions(input){
        console.log(input);
        GET(
            "/cms/tags",
            function (rst) {
                return {
                    options: [
                        {value: input, label: input}
                    ]
                }
            }.bind(this),
            function (e) {
                return {
                    options: [
                        {value: 'default', label: 'default'}
                    ]
                }
            }.bind(this)
        );
    },
    render(){
        return (
            <div className="form-group">
                <label>{i18next.t('models.cms.article.tags')}</label>
                <SelectInput.Async multi={true} id={this.props.id} loadOptions={this.getOptions}/>
            </div>

        )
    }
});

export const Show = React.createClass({
    render(){
        return <div>tag {this.props.params.id}</div>;
    }
});

const IndexW = React.createClass({
    render(){
        return <div>tags</div>;
    }
});

IndexW.propTypes = {
    user: PropTypes.object.isRequired
};

export const Index = connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(IndexW);




