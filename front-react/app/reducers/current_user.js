import jwtDecode from 'jwt-decode'
import {SIGN_IN,SIGN_OUT} from '../constants'

const key = 'token';

function parse(tkn){
    try{
        return jwtDecode(tkn);
    }catch (e){
        console.log(e);
        return {}
    }
}

const initState =parse(sessionStorage.getItem(key));

export function current_user(state = initState, action) {
    switch (action.type) {
        case SIGN_IN:
            sessionStorage.setItem(key, action.token);
            return parse(action.token);
        case SIGN_OUT:
            sessionStorage.removeItem(key);
            return {};
        default:
            return state;
    }
}

