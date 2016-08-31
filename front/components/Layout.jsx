import React from 'react'
import {Link} from 'react-router'

import i18n from 'i18next'

const Widget = ({children}) => (
    <div>
        layout:
        <Link to={`/`}>{i18n.t('platform.home')}</Link>
        &nbsp;
        <Link to={`/auth/sign-in`}>Sign in</Link>
        &nbsp;
        <Link to={`/auth/sign-up`}>Sign up</Link>
        <br />
        {children}
    </div>
);


export default Widget
