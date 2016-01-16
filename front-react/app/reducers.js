import { combineReducers } from 'redux'

import { SIGN_IN, SIGN_OUT } from './actions'

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
    session
});

export default reducers