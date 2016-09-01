import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';


import i18n from 'i18next'

import {refresh} from '../engines/platform/actions'

import Header from './Header'
import NavBar from './NavBar'
import Footer from './Footer'


const Widget = React.createClass({
    getInitialState: function() {
        const {onRefresh} = this.props;
        onRefresh();
        return {};
    },
    render: function () {
        const {children} = this.props;
        return (
            <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
                <div>
                    <Header/>
                    <NavBar/>
                    {children}
                    <hr/>
                    <Footer/>
                </div>
            </MuiThemeProvider>
        )

    }
});

Widget.propTypes = {
    children: PropTypes.object.isRequired,
    onRefresh: PropTypes.func.isRequired
};

export default connect(
    state=>({}),
    dispatch => ({
        onRefresh: function () {
            fetch('http://localhost:3000/api/site/info')
                .then(res => res.json())
                .then(rst => dispatch(refresh(rst)));
            // ajax("get", "/site/info", null, function(ifo){
            //   dispatch(refresh(ifo));
            //   document.documentElement.lang = ifo.lang;
            //   document.title = ifo.title;
            // });
        }
    })
)(Widget);

