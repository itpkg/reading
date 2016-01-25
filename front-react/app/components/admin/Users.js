import React, {PropTypes} from 'react';
import {Table} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

import {GET} from '../../ajax'

const Widget = React.createClass({
    getInitialState() {
        return {items: []}
    },
    componentDidMount() {
        GET(
            '/admin/users',
            function (rst) {
                this.setState({items: rst});
            }.bind(this)
        );
    },
    render(){
        var showPermission = function (p, i) {
            return (
                <tr key={i}>
                    <td>{p.role}</td>
                    <td>{p.begin}</td>
                    <td>{p.end}</td>
                </tr>
            );
        };
        return (
            <div>
                {this.state.items.map(function (u, i) {
                    return (<blockquote key={i}>
                        <Table striped bordered condensed hover>
                            <thead>
                            <tr>
                                <th>{i18next.t('models.permission.role')}</th>
                                <th>{i18next.t('models.permission.begin')}</th>
                                <th>{i18next.t('models.permission.end')}</th>
                            </tr>
                            </thead>
                            <tbody>
                            {u.permissions.map(showPermission)}
                            </tbody>
                        </Table>
                        <footer>{u.label} <cite>{u.lastSignIn}</cite></footer>
                    </blockquote>);
                })}
            </div>
        )
    }
});
export default Widget