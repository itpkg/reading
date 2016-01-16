export const SIGN_IN = 'sign in';
export const SIGN_OUT = 'sign out';

export function signIn(user) {
    return {
        type: SIGN_IN,
        user
    };
}

export function signOut() {
    return {type: SIGN_OUT}
}
