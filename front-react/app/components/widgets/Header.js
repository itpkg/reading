import React, {PropTypes} from 'react';
import {Navbar, Nav, NavItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {Link} from 'react-router';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

import PersonalBar from './PersonalBar'
import LangBar from './LangBar'

function Header({title}) {
    return (
        <Navbar inverse fixedTop fluid>
            <Navbar.Header>
                <Navbar.Brand>
                    <Link to="home">{title}</Link>
                </Navbar.Brand>
                <Navbar.Toggle />
            </Navbar.Header>
            <Navbar.Collapse>
                <Nav>
                    <LinkContainer to="/home">
                        <NavItem eventKey={1}>{i18next.t("nav_bar.home")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/cms/articles'}>
                        <NavItem eventKey={2}>{i18next.t("nav_bar.articles")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/books'}>
                        <NavItem eventKey={3}>{i18next.t("nav_bar.books")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/video/items'}>
                        <NavItem eventKey={4}>{i18next.t("nav_bar.videos")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/cms/articles/faq'}>
                        <NavItem eventKey={5}>{i18next.t("nav_bar.faq")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/cms/articles/about'}>
                        <NavItem eventKey={6}>{i18next.t("nav_bar.about")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'/cms/articles/contact'}>
                        <NavItem eventKey={7}>{i18next.t("nav_bar.contact")}</NavItem>
                    </LinkContainer>
                    <PersonalBar />
                </Nav>
                <LangBar />
            </Navbar.Collapse>
        </Navbar>
    )
}


Header.propTypes = {
    title: PropTypes.string.isRequired
};


export default connect(
    state => ({
        title: state.site_info.subTitle
    }),
    dispatch => ({})
)(Header);
