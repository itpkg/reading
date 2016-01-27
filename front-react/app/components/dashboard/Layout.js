import React,{PropTypes} from 'react';
import {Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

import NoMatch from '../NoMatch'
import {CurrentUser} from '../../mixins'

const Layout = React.createClass({
    mixins: [CurrentUser],
    render() {
        const {children, user} = this.props;
        if (!this.isSignIn()) {
            return (<NoMatch/>)
        }
        var items = [];
        items.push(
            <LinkContainer key="logs" to="/dashboard/personal/logs">
                <NavItem eventKey='logs'>{i18next.t('dashboard.personal.logs')}</NavItem>
            </LinkContainer>
        );
        items.push(
            <LinkContainer key="attachments" to="/dashboard/attachments">
                <NavItem eventKey='attachments'>{i18next.t('dashboard.attachments')}</NavItem>
            </LinkContainer>
        );

        if (user.isAdmin) {
            items.push(
                <NavDropdown key="admin.site" title={i18next.t('dashboard.admin.site.index')} id="dashboard-admin">
                    <LinkContainer to="/dashboard/admin/site/top">
                        <MenuItem>{i18next.t('dashboard.admin.site.top')}</MenuItem>
                    </LinkContainer>
                    <MenuItem divider/>
                    <LinkContainer to="/dashboard/admin/site/info">
                        <MenuItem>{i18next.t('dashboard.admin.site.info')}</MenuItem>
                    </LinkContainer>
                    <LinkContainer to="/dashboard/admin/site/seo">
                        <MenuItem>{i18next.t('dashboard.admin.site.seo')}</MenuItem>
                    </LinkContainer>
                    <LinkContainer to="/dashboard/admin/site/secrets">
                        <MenuItem>{i18next.t('dashboard.admin.site.secrets')}</MenuItem>
                    </LinkContainer>
                    <MenuItem divider/>
                    <LinkContainer to="/dashboard/admin/notices">
                        <MenuItem>{i18next.t('dashboard.admin.notices')}</MenuItem>
                    </LinkContainer>
                    <MenuItem divider/>
                    <LinkContainer to="/dashboard/admin/locales">
                        <MenuItem>{i18next.t('dashboard.admin.locales')}</MenuItem>
                    </LinkContainer>
                </NavDropdown>
            );
            items.push(
                <LinkContainer key="users" to="/dashboard/admin/users">
                    <NavItem eventKey='users'>{i18next.t('dashboard.admin.users')}</NavItem>
                </LinkContainer>
            );
        }
        return (
            <div className="row">
                <br/>
                <div className="col-md-2">
                    <Nav bsStyle="pills" stacked activeKey={1}>
                        {items}
                    </Nav>
                </div>
                <div className="col-md-10">
                    {children}
                </div>
            </div>
        )
    }

});

Layout.propTypes = {
    user: PropTypes.object.isRequired
};

export  default connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(Layout);

