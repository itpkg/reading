import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import {Link} from 'react-router'
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {refresh} from '../engines/platform/actions'

import i18n from 'i18next'

import Header from './Header'
import NavBar from './NavBar'
import Footer from './Footer'

const Widget = ({children}) => (
    <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
        <div>
            <Header/>
            <NavBar/>
            {children}
            <hr/>
            <Footer/>
        </div>
    </MuiThemeProvider>
);

Widget.propTypes = {
    onRefresh: PropTypes.func.isRequired,
    children: PropTypes.object.isRequired
};

export default connect(
  state=>({}),
  dispatch => ({
    onRefresh: function(){
        console.log("on refresh");
      // ajax("get", "/site/info", null, function(ifo){
      //   dispatch(refresh(ifo));
      //   document.documentElement.lang = ifo.lang;
      //   document.title = ifo.title;
      // });
    }
  })
)(Widget);

