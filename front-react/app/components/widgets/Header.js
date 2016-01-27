import React, {PropTypes} from 'react';
import {Navbar, Nav, NavItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {Link} from 'react-router';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

import PersonalBar from './PersonalBar'
import LangBar from './LangBar'

function Header({title, topNavBar}) {
    var links = topNavBar.split('\n').filter(function (u) {
        return u != ''
    }).map(function (u) {
        u = u.trim();
        return {
            label: "links" + u.replace(/\//g, '.'),
            href: u
        };
    });

    return (
        <Navbar inverse fixedTop fluid>
            <Navbar.Header>
                <Navbar.Brand>
                    <Link to="/home">{title}</Link>
                </Navbar.Brand>
                <Navbar.Toggle />
            </Navbar.Header>
            <Navbar.Collapse>
                <Nav>
                    <LinkContainer to="/home">
                        <NavItem eventKey={1}>{i18next.t("links.home")}</NavItem>
                    </LinkContainer>
                    {links.map(function (l, i) {
                        return (
                            <LinkContainer key={i} to={l.href}>
                                <NavItem eventKey={i+2}>{i18next.t(l.label)}</NavItem>
                            </LinkContainer>
                        )
                    })}
                    <PersonalBar />
                </Nav>
                <LangBar />
            </Navbar.Collapse>
        </Navbar>
    )
}


Header.propTypes = {
    title: PropTypes.string.isRequired,
    topNavBar: PropTypes.string.isRequired
};

export default Header
