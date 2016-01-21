import {FETCH_OAUTH} from '../constants'

export function refresh(oauth) {
    return {
        type: FETCH_OAUTH,
        oauth
    };
}