import {FETCH_OAUTH} from '../constants'

export function oauth(state = {}, action) {
    switch (action.type) {
        case FETCH_OAUTH:
            return action.oauth;
        default:
            return state;
    }
}

