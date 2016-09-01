export const AUTH_SIGN_IN = "platform.auth.sign_in";
export const AUTH_SIGN_OUT = "platform.auth.sign_out";

export const REFRESH_SITE_INFO = "platform.site.refresh";

export const TOGGLE_NAV_BAR = "platform.nav_bar.toggle";

export function signIn(token) {
    return {type: AUTH_SIGN_IN, token}
}

export function signOut() {
    return {type: AUTH_SIGN_OUT}
}

export function refresh(info) {
    return {type: REFRESH_SITE_INFO, info}
}


export function toggle() {
    return {type: TOGGLE_NAV_BAR}
}

