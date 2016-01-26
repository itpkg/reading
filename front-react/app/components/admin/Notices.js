import React from 'react';
import {ListGroup, ListGroupItem, Input} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

import {GET, POST, DELETE, response, failed} from '../../ajax';
import Buttons from '../widgets/FormButtons';

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState() {
        return {
            data: {
                items: []
            }
        }
    },
    componentDidMount() {
        this.reloadNotices();
    },
    reloadNotices(){
        GET(
            '/notices',
            function (rst) {
                this.setState({data: rst});
            }.bind(this)
        );
    },
    handleReset(){
        this.setState({id: null, content: ''});
        this.reloadNotices();
    },
    handleDelete(e){
        e.preventDefault();
        var id = this.state.id;
        if (!id) {
            failed();
            return
        }
        DELETE("/admin/notice/" + id, response(this.handleReset))
    },
    handleSubmit(e){
        e.preventDefault();
        var data = {id: this.state.id, content: this.state.content};
        if (data.content === '') {
            failed();
            return
        }
        POST(
            "/admin/notices",
            data,
            response(this.handleReset)
        )
    },
    render(){
        var self = this;
        var handleClick = function (id) {
            var item = self.state.data.items[id];
            self.setState(item);
        };
        return (
            <div>

                <form method='POST' onSubmit={this.handleSubmit}>
                    <Input type="hidden" id="id" valueLink={this.linkState('id')}/>
                    <Input id='content'
                           type='textarea' label='forms.notice.content'
                           valueLink={this.linkState('content')}
                    />
                    <Buttons onReset={this.handleReset} onDelete={this.handleDelete}/>
                </form>
                <br/>
                <ListGroup>
                    {this.state.data.items.map(function (item, idx) {
                        return (
                            <ListGroupItem onClick={handleClick.bind(this, idx)} key={item.id}>
                                {item.created_at}: {item.content}
                            </ListGroupItem>
                        )
                    })}
                </ListGroup>
            </div>
        )
    }
});
export default Widget;