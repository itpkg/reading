import React from 'react'
import {Link} from 'react-router'

const Widget = ({children}) => (
    <div>
        layout:
        <Link to={`/`}>Index</Link>
        &nbsp;
        <Link to={`/auth/sign-in`}>Sign in</Link>
        &nbsp;
        <Link to={`/auth/sign-up`}>Sign up</Link>
        <br />
        {children}
    </div>
);


export default Widget
