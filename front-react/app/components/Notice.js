import React,{PropTypes} from 'react';
import {connect} from 'react-redux';
import TimeAgo from 'react-timeago';

import {GET} from  '../ajax'
import RemoveButton from './widgets/RemoveButton'
import Markdown from './widgets/Markdown'

export const ShowW = React.createClass({
    render(){
        const {user, notice, onRefresh} = this.props;

        return (<blockquote>
            <p>
                <Markdown body={notice.content}/>
            </p>
            <footer>
                {user.isAdmin ?
                    <RemoveButton onRefresh={onRefresh} action={"/admin/notice/"+notice.id} size="xsmall"/> : ""}
                &nbsp;
                <cite><TimeAgo date={notice.created_at}/></cite>
            </footer>
        </blockquote>)
    }
});


ShowW.propTypes = {
    user: PropTypes.object.isRequired,
    notice: PropTypes.object.isRequired
};

export const Show = connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(ShowW);

export const List = React.createClass({
    render(){
        const {notices, onRefresh} = this.props;
        return (
            <div>
                {notices.map(function (n, i) {
                    return <Show key={i} notice={n} onRefresh={onRefresh}/>
                })}
            </div>
        )
    }
});

export const Index = React.createClass({
    getInitialState() {
        return {
            data: {
                items: []
            }
        }
    },
    componentDidMount() {
        GET(
            '/notices',
            function (rst) {
                this.setState({data: rst});
            }.bind(this)
        );
    },
    render(){
        return (
            <div className="row">
                <div className="col-md-offset-1 col-md-10">
                    <br/>
                    {this.state.data.items.map(function (n, i) {
                        return <Show key={i} notice={n}/>
                    })}
                </div>
            </div>
        )
    }
});

//
//List.propTypes = {
//    notices: PropTypes.arrayOf(PropTypes.shape({
//        id: PropTypes.number.isRequired,
//        content: PropTypes.string.isRequired,
//        created_at: PropTypes.string.isRequired
//    }).isRequired).isRequired
//};


