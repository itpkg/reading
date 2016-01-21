import React, {Component, PropTypes} from 'react';
import {Tabs, Tab, Alert} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import $ from 'jquery';

import SiteInfo from '../admin/site/Info'
import SiteSeo from '../admin/site/Seo'
import Roles from '../admin/Roles'
import Logs from './Logs'

import {GET} from '../../ajax'
import NoMatch from '../NoMatch'

const Profile = React.createClass({
    getInitialState() {
        const {user} = this.props;
        return {
            user: user,
            key: 'users.logs',
            logs: [],
            alert: {
                style: 'success',
                message: '',
                show: false
            }
        };
    },
    handleSelect(key) {
        this.setState({alert: {show: false}});
        this.setState({key});
        switch (key) {
            case 'users.logs':
                GET(
                    '/users/logs',
                    function (rst) {
                        this.setState({logs: rst});
                    }.bind(this),
                    this.handleFail
                );
        }
    },
    handleFail(data){
        this.setState({
            alert: {
                style: 'danger',
                message: data.responseText,
                show: true
            }
        })
    },
    componentDidMount() {
        this.handleSelect('users.logs');
    },
    render() {
        const {user} = this.props;
        if ($.isEmptyObject(user)) {
            return (<NoMatch/>)
        } else {
            var alert = function (m) {
                if (m.show) {
                    return (
                        <div>
                            <br/>
                            <Alert bsStyle={m.style}>
                                <strong>{m.message}</strong> {new Date().toLocaleString()}
                            </Alert>
                        </div>
                    )
                } else {
                    return <br/>
                }
            };
            return (
                <div className="col-md-offset-1 col-md-10">
                    {alert(this.state.alert)}
                    <Tabs activeKey={this.state.key} onSelect={this.handleSelect}>
                        <Tab eventKey={'users.logs'} title={i18next.t('users.logs')}>
                            <Logs logs={this.state.logs}/>
                        </Tab>
                        <Tab eventKey={'roles'} title={i18next.t('admin.roles')}>
                            <Roles />
                        </Tab>
                        <Tab eventKey={'site.info'} title={i18next.t('admin.site.info')}>
                            <SiteInfo />
                        </Tab>
                        <Tab eventKey={'site.seo'} title={i18next.t('admin.site.seo')}>
                            <SiteSeo />
                        </Tab>
                    </Tabs>
                </div>
            )
        }
    }
});

Profile.propTypes = {
    user: PropTypes.object.isRequired
};

export default connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(Profile);