import React, {PropTypes} from 'react';
import {Table} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

const Widget = React.createClass({
    render(){
        return (
            <div>
                <br/>
                <Table striped bordered condensed hover>
                    <thead>
                    <tr>
                        <th>{i18next.t('models.permission.user')}</th>
                        <th>{i18next.t('models.permission.role')}</th>
                        <th>{i18next.t('models.permission.begin')}</th>
                        <th>{i18next.t('models.permission.end')}</th>
                    </tr>
                    </thead>
                    <tbody>
                    {this.props.items.map(function (item, idx) {
                        return (
                            <tr key={idx}>
                                <td>{item.user}</td>
                                <td>{item.role}</td>
                                <td>{item.begin}</td>
                                <td>{item.end}</td>
                            </tr>
                        );
                    })}
                    </tbody>
                </Table>
            </div>
        )
    }
});
export default Widget