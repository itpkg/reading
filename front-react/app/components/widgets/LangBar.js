import React, {PropTypes} from 'react';
import {Nav, NavItem} from 'react-bootstrap'
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

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


LangBar.propTypes = {
    onSwitchLang: PropTypes.func.isRequired
};


export default connect(
    state => ({}),
    dispatch => ({
        onSwitchLang: function (lang) {
            localStorage.setItem("locale", lang);
            location.reload();
        }
    })
)(LangBar);