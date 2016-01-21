import React, { Component, PropTypes } from 'react';
import {Navbar, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {Link} from 'react-router';
import {connect} from 'react-redux';
import {routeActions} from 'redux-simple-router';
import i18next from 'i18next/lib';
import $ from 'jquery';


import {refresh as refreshSiteInfo} from '../actions/site'
import {refresh as refreshOauth} from '../actions/auth'
import {GET} from '../ajax'


function LangBar({onSwitchLang}) {
    return (
        <Nav pullRight>
            <NavItem eventKey={1}
                     onClick={()=>onSwitchLang("en-US")}>{i18next.t("locales.en_US")}</NavItem>
            <NavItem eventKey={2}
                     onClick={()=>onSwitchLang("zh-CN")}>{i18next.t("locales.zh_CN")}</NavItem>
        </Nav>
    )
}

class PersonalBar extends Component {
    componentDidMount() {
        const {user, onOauth} = this.props;
        if ($.isEmptyObject(user)) {
            onOauth();
        }
    }

    render() {
        const {user,oauth} = this.props;
        if ($.isEmptyObject(user)) {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.sign_up_or_in")} id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1}
                              href={oauth.google}>{i18next.t("users.sign_in_with.google")}</MenuItem>
                </NavDropdown>
            )
        } else {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1} href="/#/personal/profile">{i18next.t("users.titles.profile")}</MenuItem>
                    <MenuItem divider/>
                    <MenuItem eventKey={3.3}>{i18next.t("users.titles.sign_out")}</MenuItem>
                </NavDropdown>
            )
        }
    }
}

function Header({title, currentUser, oauth, onSwitchLang, onOauth}) {
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
                    <LinkContainer to="home">
                        <NavItem eventKey={1}>{i18next.t("nav_bar.home")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'cms/articles'}>
                        <NavItem eventKey={2}>{i18next.t("nav_bar.articles")}</NavItem>
                    </LinkContainer>
                    <LinkContainer to={'books'}>
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
                    <PersonalBar user={currentUser} oauth={oauth} onOauth={onOauth}/>
                </Nav>
                <LangBar onSwitchLang={onSwitchLang}/>
            </Navbar.Collapse>
        </Navbar>
    )
}


function Footer({copyright}) {
    return (
        <footer>
            <p>
                {copyright}
                &nbsp;
                <span
                    dangerouslySetInnerHTML={{__html: i18next.t('build_using', {url:'https://github.com/itpkg/reading'})}
}/>
            </p>
        </footer>
    )

}

class Layout extends Component {
    componentDidMount() {
        const {onRefresh} = this.props;
        onRefresh();
    }

    render() {
        const {siteInfo, children, currentUser, oauth, onSwitchLang, onOauth} = this.props;
        return (
            <div>
                <Header title={siteInfo.subTitle}
                        oauth={oauth}
                        currentUser={currentUser}
                        onOauth={onOauth}
                        onSwitchLang={onSwitchLang}/>
                <div className="container-fluid">
                    <div className="row">
                        {children}
                    </div>
                    <hr/>
                    <Footer copyright={siteInfo.copyright}/>
                </div>
            </div>
        );
    }
}

Layout.propTypes = {
    siteInfo: PropTypes.object.isRequired,
    currentUser: PropTypes.object.isRequired,
    oauth: PropTypes.object.isRequired,
    onRefresh: PropTypes.func.isRequired,
    onSwitchLang: PropTypes.func.isRequired,
    onOauth: PropTypes.func.isRequired
};


export default connect(
    state => ({
        siteInfo: state.site_info,
        currentUser: state.current_user,
        oauth: state.oauth
    }),
    dispatch => ({
        onRefresh: function () {
            GET('/site/info', function (rst) {
                dispatch(refreshSiteInfo(rst));
            })
        },
        onOauth: function () {
            GET('/oauth/sign_in', function (rst) {
                dispatch(refreshOauth(rst));
            })
        },
        onSwitchLang: function (lang) {
            localStorage.setItem("locale", lang);
            location.reload();
        }
    })
)(Layout);
