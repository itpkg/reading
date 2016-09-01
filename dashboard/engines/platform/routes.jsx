import React from 'react'
import {Route} from 'react-router'

import {SignIn, SignUp, Confirm, Unlock, ForgotPassword, Profile} from './auth'

export default    [
    <Route key="platform.auth" path="auth">
        <Route path="sign-in" component={SignIn}/>
        <Route path="sign-up" component={SignUp}/>
        <Route path="confirm" component={Confirm}/>
        <Route path="unlock" component={Unlock}/>
        <Route path="forgot-password" component={ForgotPassword}/>
        <Route path="profile" component={Profile}/>
    </Route>
]

