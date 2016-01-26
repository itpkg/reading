import React, {PropTypes} from 'react';
import i18next from 'i18next/lib';

import {GET} from '../../../ajax';
import Form from '../../widgets/Form';
import {List as Notices} from '../../Notice'

const Widget = React.createClass({
    getInitialState() {
        return {
            data: {
                items: []
            }
        }
    },
    componentDidMount() {
        this.loadNotices();
    },
    loadNotices(){
        GET(
            '/notices',
            function (rst) {
                this.setState({data: rst});
            }.bind(this)
        );
    },
    render(){
        return (
            <div>
                <Form title={i18next.t('buttons.create')} action="/admin/notices" fields={[
                {
                    id:"content",
                    type:"textarea"
                }
            ]} onSubmit={this.loadNotices}/>
                <br/>
                <Notices notices={this.state.data.items} onRefresh={this.loadNotices}/>
            </div>
        )
    }
});
export default Widget;