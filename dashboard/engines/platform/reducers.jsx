import jwtDecode from 'jwt-decode'

import {
    AUTH_SIGN_IN, AUTH_SIGN_OUT,
    REFRESH_SITE_INFO,
TOGGLE_MESSAGE_BOX,
    TOGGLE_NAV_BAR
} from './actions'

import {TOKEN} from '../../constants'

function parse(tkn) {
    try {
        return jwtDecode(tkn);
    } catch (e) {
        return {}
    }
}

const initCurrentUserState = parse(sessionStorage.getItem(TOKEN));

function currentUser(state = initCurrentUserState, action) {
    switch (action.type) {
        case AUTH_SIGN_IN:
            sessionStorage.setItem(TOKEN, action.token);
            return parse(action.token);
        case AUTH_SIGN_OUT:
            sessionStorage.removeItem(TOKEN);
            return {};
        default:
            return state
    }
}

function siteInfo(state = {}, action) {
    switch (action.type) {
        case REFRESH_SITE_INFO:
            return action.info;
        default:
            return state
    }
}

function navBar(state = {}, action) {
    switch (action.type) {
        case TOGGLE_NAV_BAR:
            return !action.open;
        default:
            return state
    }
}

function messageBox(state={show:false, message:''}, action) {
    switch (action.type) {
        case TOGGLE_MESSAGE_BOX:
            return action.data;
        default:
            return state
    }
}

const reducers = {currentUser, siteInfo, navBar, messageBox};

export default reducers