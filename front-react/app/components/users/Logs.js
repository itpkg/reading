import React, {PropTypes} from 'react';
import {ListGroup, ListGroupItem} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import TimeAgo from 'react-timeago';

const Widget = React.createClass({
    render(){
        return (
            <div>
                <br/>
                <ListGroup>
                    {this.props.items.map(function (item) {
                        return (<ListGroupItem key={item.id}>
                            <TimeAgo date={item.created_at}/>: {item.message}
                        </ListGroupItem>)
                    })}
                </ListGroup>
            </div>
        )
    }
});
export default Widget