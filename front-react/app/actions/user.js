import {SIGN_IN, SIGN_OUT} from '../constants'

export function signIn(token) {
    return {
        type: SIGN_IN,
        token: token
    };
}

export function signOut() {
    return {
        type: SIGN_OUT
    };
}


