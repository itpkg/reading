import {FETCH_SITE_INFO} from '../constants'

export function refresh(info) {
    return {
        type: FETCH_SITE_INFO,
        info
    };
}