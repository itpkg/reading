import i18n from 'i18next';

export const AUTH_SIGN_IN = "platform.auth.sign_in";
export const AUTH_SIGN_OUT = "platform.auth.sign_out";

export const REFRESH_SITE_INFO = "platform.site.refresh";

export const TOGGLE_NAV_BAR = "platform.nav_bar.toggle"; //fixme

export const TOGGLE_MESSAGE_BOX = "platform.message_box.toggle";

export function signIn(token) {
    return {type: AUTH_SIGN_IN, token}
}

export function signOut() {
    return {type: AUTH_SIGN_OUT}
}

export function refresh(info) {
    return {type: REFRESH_SITE_INFO, info}
}


export function showNavBar() {
    return {type: TOGGLE_NAV_BAR}
}


export function showMessage(msg) {
    var data = msg ? {show:true, message:msg} : {show:false, message:''};
    return {type: TOGGLE_MESSAGE_BOX, data}
}


export function checkResult(rst) {
    var data = {show: true, message: i18n.t('notices.success')};
    if (rst.errors) {
        data.message = i18n.t('notices.fail')+JSON.stringify(rst.errors);
    }

    return {type: TOGGLE_MESSAGE_BOX, data}

}

