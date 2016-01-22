import React, {Component, PropTypes} from 'react';
import {Tabs, Tab, Alert} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import $ from 'jquery';

import SiteInfo from '../admin/site/Info'
import SiteSeo from '../admin/site/Seo'
import SiteSecrets from '../admin/site/Secrets'
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
            roles: [],
            site_seo: {},
            site_info: {},
            site_secrets: {},
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
            case 'admin.site.seo':
                GET(
                    '/admin/site/seo',
                    function (rst) {
                        this.setState({site_seo: rst});
                    }.bind(this),
                    this.handleFail
                );
                break;
            case 'admin.site.secrets':
                GET(
                    '/admin/site/secrets',
                    function (rst) {
                        this.setState({site_secrets: rst});
                    }.bind(this),
                    this.handleFail
                );
                break;
            case 'admin.site.info':
                GET(
                    '/admin/site/info',
                    function (rst) {
                        this.setState({site_info: rst});
                    }.bind(this),
                    this.handleFail
                );
                break;
            case 'admin.roles':
                GET(
                    '/admin/roles',
                    function (rst) {
                        this.setState({roles: rst});
                    }.bind(this),
                    this.handleFail
                );
                break;
            case 'users.logs':
                GET(
                    '/users/logs',
                    function (rst) {
                        this.setState({logs: rst});
                    }.bind(this),
                    this.handleFail
                );
                break;
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
            var tabs = [(
                <Tab key="users.logs" eventKey={'users.logs'} title={i18next.t('users.logs')}>
                    <Logs items={this.state.logs}/>
                </Tab>
            )];
            if (user.isAdmin) {
                tabs.push(<Tab key="admin.roles" eventKey={'admin.roles'} title={i18next.t('admin.roles')}>
                    <Roles items={this.state.roles}/>
                </Tab>);
                tabs.push(<Tab key="admin.site.info" eventKey={'admin.site.info'} title={i18next.t('admin.site.info')}>
                    <SiteInfo item={this.state.site_info}/>
                </Tab>);
                tabs.push(<Tab key="admin.site.seo" eventKey={'admin.site.seo'} title={i18next.t('admin.site.seo')}>
                    <SiteSeo item={this.state.site_seo}/>
                </Tab>);
                tabs.push(<Tab key="admin.site.secrets" eventKey={'admin.site.secrets'}
                               title={i18next.t('admin.site.secrets')}>
                    <SiteSecrets item={this.state.site_secrets}/>
                </Tab>);
            }
            return (
                <div className="col-md-offset-1 col-md-10">
                    {alert(this.state.alert)}
                    <Tabs activeKey={this.state.key} onSelect={this.handleSelect}>
                        {tabs}
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