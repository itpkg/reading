import React,{PropTypes} from 'react';
import {Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import $ from 'jquery';

import NoMatch from '../NoMatch'

const Layout = React.createClass({
    render() {
        const {children, user} = this.props;
        if ($.isEmptyObject(user)) {
            return (<NoMatch/>)
        } else {
            var items = [];
            items.push(
                <LinkContainer key="users" to="/dashboard/users/logs">
                    <NavItem eventKey={1}>{i18next.t('dashboard.users.logs')}</NavItem>
                </LinkContainer>
            );
            if (user.isAdmin) {
                items.push(
                    <NavDropdown key="admin.site" title={i18next.t('dashboard.admin.site.index')} id="dashboard-admin">
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
                        <LinkContainer to="/dashboard/admin/roles">
                            <MenuItem>{i18next.t('dashboard.admin.roles')}</MenuItem>
                        </LinkContainer>
                    </NavDropdown>
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
    }
});

Layout.propTypes = {
    user: PropTypes.object.isRequired
};

export  default connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(Layout);

