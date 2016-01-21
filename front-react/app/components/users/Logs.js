import React, {PropTypes} from 'react';
import {ListGroup, ListGroupItem} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

const Widget = React.createClass({
    render(){
        return (
            <ListGroup>
                {this.props.logs.forEach(function (item) {
                    return <ListGroupItem>{item.created}: {item.message}</ListGroupItem>
                })}
            </ListGroup>
        )
    }
});
export default Widget