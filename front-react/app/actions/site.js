import {FETCH_SITE_INFO} from '../constants'

export function refresh(info) {
    document.title = info.subTitle + '-' + info.title;
    return {
        type: FETCH_SITE_INFO,
        info
    };
}