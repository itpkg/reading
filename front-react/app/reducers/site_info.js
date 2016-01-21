import {FETCH_SITE_INFO} from '../constants'

export function site_info(state = {}, action) {
    switch (action.type) {
        case FETCH_SITE_INFO:
            return action.info;
        default:
            return state;
    }
}

