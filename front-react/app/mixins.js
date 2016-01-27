import $ from 'jquery'
import i18next from 'i18next/lib';

export const CurrentUser = {
    isSignIn(){
        const {user} = this.props;
        return !$.isEmptyObject(user);
    }
};
