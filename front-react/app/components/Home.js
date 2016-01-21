import React from 'react'
import {connect} from 'react-redux'

import {signIn, signOut} from '../actions/user'

function Home({}) {
    return (
        <div>
            home
        </div>
    );
}

export default connect(
    state => ({user: state.current_user, title: state.site_info.title}),
    {signIn, signOut}
)(Home);
