import { combineReducers } from 'redux'

import { SIGN_IN, SIGN_OUT, SITE_INFO } from './actions'

function site_info(state={}, action){
    switch (action.type){
        case SITE_INFO:
            return action.info;
        default:
            return state;
    }
}

function session(state = {}, action) {
    switch (action.type) {
        case SIGN_IN:
            return {
                ...state,
                current_user: action.user
            };
        case SIGN_OUT:
            return {
                ...state,
                current_user: null
            };
        default:
            return state;
    }
}


const reducers = combineReducers({
    session,
    site_info
});

export default reducers