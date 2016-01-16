import React from 'react';
import {IndexLink, Link} from 'react-router';
import {LinkContainer} from 'react-router-bootstrap';
import {Input, ButtonInput,Navbar, Nav, NavItem, NavDropdown, MenuItem, Alert} from 'react-bootstrap'
import i18next from 'i18next/lib';

import {AjaxMixin} from '../mixins/ajax'

export const Header = React.createClass({
    onSignOut: function () {
//todo
    },
    personalBar: function () {
        var user = null;//todo this.state.current_user;
        if (user) {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.titles.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1} href="/#/personal/profile">{i18next.t("users.titles.profile")}</MenuItem>
                    <MenuItem divider/>
                    <MenuItem eventKey={3.3} onclick={this.onSignOut}>{i18next.t("users.titles.sign_out")}</MenuItem>
                </NavDropdown>)
        } else {
            return (<NavDropdown eventKey={3} title={i18next.t("users.titles.sign_in_or_up")} id="basic-nav-dropdown">
                <MenuItem eventKey={3.1} href="/#/users/sign-in">{i18next.t("users.titles.sign_in")}</MenuItem>
                <MenuItem eventKey={3.2} href="/#/users/sign-up">{i18next.t("users.titles.sign_up")}</MenuItem>
                <MenuItem eventKey={3.3}
                          href="/#/users/forgot-password">{i18next.t("users.titles.forgot_your_password")}</MenuItem>
                <MenuItem eventKey={3.4}
                          href="/#/users/confirm">{i18next.t("users.titles.did_not_receive_confirmation_instructions")}</MenuItem>
                <MenuItem eventKey={3.5}
                          href="/#/users/unlock">{i18next.t("users.titles.did_not_receive_unlock_instructions")}</MenuItem>
            </NavDropdown>)
        }
    },
    switchLang: function (lang) {
        localStorage.setItem("locale", lang);
        location.reload();
    },
    render(){
        //todo info
        return (
            <Navbar inverse fixedTop fluid>
                <Navbar.Header>
                    <Navbar.Brand>
                        <Link to="home">{this.props.title}</Link>
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
                            <NavItem eventKey={2}>{i18next.t("nav_bar.books")}</NavItem>
                        </LinkContainer>
                        <LinkContainer to={'/video/items'}>
                            <NavItem eventKey={2}>{i18next.t("nav_bar.videos")}</NavItem>
                        </LinkContainer>
                        <LinkContainer to={'/cms/articles/faq'}>
                            <NavItem eventKey={2}>{i18next.t("nav_bar.faq")}</NavItem>
                        </LinkContainer>
                        <LinkContainer to={'/cms/articles/about'}>
                            <NavItem eventKey={2}>{i18next.t("nav_bar.about")}</NavItem>
                        </LinkContainer>
                        <LinkContainer to={'/cms/articles/contact'}>
                            <NavItem eventKey={2}>{i18next.t("nav_bar.contact")}</NavItem>
                        </LinkContainer>
                        {this.personalBar()}
                    </Nav>
                    <Nav pullRight>
                        <NavItem eventKey={1}
                                 onClick={this.switchLang.bind(this, "en-US")}>{i18next.t("locales.en_US")}</NavItem>
                        <NavItem eventKey={2}
                                 onClick={this.switchLang.bind(this, "zh-CN")}>{i18next.t("locales.zh_CN")}</NavItem>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        )
    }
});

export const Footer = React.createClass({
    render(){
        return (<footer>
            <p>
                {this.props.copyright}
                &nbsp;
                <span
                    dangerouslySetInnerHTML={{__html: i18next.t('build_using', {url:'https://github.com/itpkg/reading'})}}/>
            </p>
        </footer>)
    }
});


export const Form = React.createClass({
    getInitialState: function () {
        var sfs = {};
        this.props.fields.forEach(function (field) {
            sfs[field.id] = field.value;
        });

        return {fields: sfs}
    },
    handleChange: function (e) {
        var sfs = this.state.fields;
        sfs[e.target.id] = e.target.value;
        this.setState({fields: sfs});
    },
    handleAlertDismiss: function (e) {
        this.setState({result: undefined});
    },
    handleSubmit: function (e) {
        e.preventDefault();
        switch (this.props.method) {
            default:
                $.post(
                    this.props.action + "?locale=" + i18next.language,
                    this.state.fields,
                    function (rs) {
                        var submit = this.props.submit;
                        if (submit && rs.ok) {
                            submit(rs.data);
                        } else {
                            this.setState({result: rs});
                        }
                    }.bind(this));
        }
    },
    render: function () {
        var handleChange = this.handleChange;
        var resource = this.props.resource;

        var dialog = function (rs, dis) {
            if (rs) {
                var style = "danger";
                var data = rs.errors;
                if (rs.ok) {
                    style = "success";
                    data = rs.data;
                }
                return (<Alert bsStyle={style} onDismiss={dis}>
                    <h4>{data[0]}</h4>
                    <ul>
                        {data.slice(1).map(function (msg, idx) {
                            return (<li key={"item-"+idx}>{msg}</li>)
                        })}
                    </ul>
                </Alert>)
            } else {
                return <br/>
            }

        };
        var fields = this.props.fields.map(function (field) {
            var key = 'k-' + field.id;
            var label = i18next.t(resource + ".fields." + field.id);

            if (field.required) {
                label = "* " + label;
            }
            switch (field.type) {
                case "email":
                    return <Input id={field.id} key={key} onChange={handleChange} type="email" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-6"/>;
                case "text":
                    return <Input id={field.id} key={key} onChange={handleChange} type="text" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-10"/>;
                case "password":
                    return <Input id={field.id} key={key} onChange={handleChange} type="password" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-8"/>;
                case "checkbox":
                    return <Input id={field.id} key={key} onChange={handleChange} type="checkbox" label={label}
                                  wrapperClassName="col-xs-offset-2 col-xs-10"/>;
                default:
                    return <input id={field.id} key={key} type="hidden"/>;
            }
        });
        var method = this.props.method;
        if (!method) {
            method = 'post';
        }

        return (
            <fieldset>
                <legend>{this.props.title}</legend>
                {dialog(this.state.result, this.handleAlertDismiss)}
                <form method={method} action={this.props.action}
                      className="form-horizontal">
                    {fields}
                    <div className="form-group">
                        <div className="col-xs-offset-2 col-xs-10">
                            <button type="submit" onClick={this.handleSubmit}
                                    className="btn btn-primary">{i18next.t("buttons.submit")}</button>
                            &nbsp; &nbsp;
                            <button type="reset" className="btn btn-default">{i18next.t("buttons.reset")}</button>
                        </div>
                    </div>
                </form>
            </fieldset>)
    }
});


export const Layout = React.createClass({
    mixins: [AjaxMixin],
    getInitialState(){
        return {
            subTitle: 'Reading',
            title: 'IT-PACKAGE',
            copyright: '©2016 Company, Inc.'
        }
    },
    componentDidMount(){
        document.title = this.state.subTitle + '-' + this.state.title;
        this.GET('/site/info', function (rst) {
            this.setState(rst);
        });
    },
    render(){
        return (
            <div>
                <Header title={this.state.subTitle}/>
                <div className="container-fluid">
                    <div className="row">
                        {this.props.children}
                    </div>
                    <hr/>
                    <div>
                        <Footer copyright={this.state.copyright}/>
                    </div>
                </div>
            </div>
        )
    }
});
