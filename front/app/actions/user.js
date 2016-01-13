import {SIGN_IN, SIGN_OUT} from '../constants'

export function sign_in(token){
    return {
        type:SIGN_IN,
        token:token
    }
}

export function sign_out(){
    return {
        type:SIGN_OUT
    }
}