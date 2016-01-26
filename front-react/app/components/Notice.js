import React from 'react';
import TimeAgo from 'react-timeago';

import {GET} from  '../ajax'
import RemoveButton from './widgets/RemoveButton'
import Markdown from './widgets/Markdown'

export const Show = React.createClass({
    render(){
        const {notice} = this.props;

        return (<blockquote>
            <p>
                <Markdown body={notice.content}/>
            </p>
            <footer>
                <cite><TimeAgo date={notice.created_at}/></cite>
            </footer>
        </blockquote>)
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


