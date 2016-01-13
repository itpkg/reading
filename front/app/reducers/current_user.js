import  {SIGN_IN, SIGN_OUT} from '../constants';

const initialState = {
    token:sessionStorage.getItem("token")
};

function current_user(state=initialState, action){
    if(action.type === SIGN_IN){
        return {token:action.token}
    }else if(action.type === SIGN_OUT){
        return {token:undefined}
    }
    return state
}

module.exports = current_user;