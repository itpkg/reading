import React from 'react';
import {ListGroup, ListGroupItem, Input} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'

import {GET, POST, DELETE, failed, response} from '../../ajax';
import Buttons from '../widgets/FormButtons'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState() {
        return {
            items: []
        }
    },
    componentDidMount() {
        this.reloadLocales();
    },
    reloadLocales(){
        GET(
            '/admin/locales',
            function (rst) {
                this.setState({items: rst});
            }.bind(this)
        );
    },
    handleReset(){
        this.setState({id: null, code: null, message: ''});
        this.reloadLocales();
    },
    handleDelete(e){
        e.preventDefault();
        var id = this.state.id;
        if (!id) {
            failed();
            return
        }
        DELETE("/admin/locale/" + id, response(this.handleReset))
    },
    handleSubmit(e){
        e.preventDefault();
        var data = {id: this.state.id, code: this.state.code, message: this.state.message};
        if (!data.code || data.message === '') {
            failed();
            return
        }
        POST(
            "/admin/locales",
            data,
            response(this.handleReset)
        )
    },
    render(){
        var self = this;
        var handleClick = function (id) {
            var item = self.state.items[id];
            self.setState(item);
        };
        return (
            <div>
                <form method='POST' onSubmit={this.handleSubmit}>
                    <Input type="hidden" id="id" valueLink={this.linkState('id')}/>
                    <Input id='code'
                           type='text' label='forms.locale.code'
                           valueLink={this.linkState('code')}
                    />
                    <Input id='message'
                           type='textarea' label='forms.locale.message'
                           valueLink={this.linkState('message')}
                    />
                    <Buttons onReset={this.handleReset} onDelete={this.handleDelete}/>
                </form>

                <ListGroup>
                    {this.state.items.map(function (item, idx) {
                        return (<ListGroupItem onClick={handleClick.bind(this, idx)} key={item.id}>
                            {item.code} = {item.message}
                        </ListGroupItem>)
                    })}
                </ListGroup>
            </div>
        )
    }
});
export default Widget;