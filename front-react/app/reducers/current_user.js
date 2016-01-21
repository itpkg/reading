import {SIGN_IN,SIGN_OUT} from '../constants'

const key = 'token';

export function current_user(state = {}, action) {
    switch (action.type) {
        case SIGN_IN:
            //todo 解析token
            sessionStorage.setItem(key, action.token);
            return action.user;
        case SIGN_OUT:
            sessionStorage.removeItem(key);
            return {};
        default:
            return state;
    }
}

