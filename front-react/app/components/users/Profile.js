import React, {Component, PropTypes} from 'react';
import ReactDOM from 'react-dom'
import {Tabs, Tab, Alert} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import $ from 'jquery';

import Roles from '../admin/Roles'
import Logs from './Logs'

import {GET} from '../../ajax'
import NoMatch from '../NoMatch'

import AjaxForm from '../widgets/AjaxForm'

const Profile = React.createClass({
    render() {
        const {user} = this.props;
        if ($.isEmptyObject(user)) {
            return (<NoMatch/>)
        } else {

            var tabs = [(
                <Tab key="users.logs" eventKey={'users.logs'} title={i18next.t('users.logs')}>
                    <Logs/>
                </Tab>
            )];
            if (user.isAdmin) {
                tabs.push(<Tab key="admin.roles" eventKey={'admin.roles'} title={i18next.t('admin.roles')}>
                    <Roles/>
                </Tab>);
                tabs.push(<Tab key="admin.site.info" eventKey={'admin.site.info'} title={i18next.t('admin.site.info')}>
                    <AjaxForm url="/admin/site/info"/>
                </Tab>);
                tabs.push(<Tab key="admin.site.seo" eventKey={'admin.site.seo'} title={i18next.t('admin.site.seo')}>
                    <AjaxForm url="/admin/site/seo"/>
                </Tab>);
                tabs.push(<Tab key="admin.site.secrets" eventKey={'admin.site.secrets'}
                               title={i18next.t('admin.site.secrets')}>
                    <AjaxForm url="/admin/site/secrets"/>
                </Tab>);
            }
            return (
                <div className="col-md-offset-1 col-md-10">
                    <br/>
                    <Tabs defaultActiveKey='users.logs' onSelect={this.handleSelect}>
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