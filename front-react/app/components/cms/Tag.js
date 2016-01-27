import React,{PropTypes} from 'react';
import {connect} from 'react-redux';

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




