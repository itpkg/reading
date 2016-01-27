import React, {PropTypes} from 'react';
import {ListGroup, ListGroupItem} from 'react-bootstrap'
import i18next from 'i18next/lib';
import TimeAgo from 'react-timeago';

import {GET} from '../../ajax'

const Widget = React.createClass({
    getInitialState() {
        return {items: []}
    },
    componentDidMount() {
        GET(
            '/users/logs',
            function (rst) {
                this.setState({items: rst});
            }.bind(this)
        );
    },
    render(){
        return (
            <ListGroup>
                {this.state.items.map(function (item) {
                    return (<ListGroupItem key={item.id}>
                        <TimeAgo date={item.created_at}/>: {item.message}
                    </ListGroupItem>)
                })}
            </ListGroup>
        )
    }
});
export default Widget