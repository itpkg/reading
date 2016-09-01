import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';


import i18n from 'i18next'

import {refresh} from '../engines/platform/actions'
import {get} from '../ajax'

import Header from './Header'
import NavBar from './NavBar'
import Footer from './Footer'


const Widget = React.createClass({
    getInitialState: function () {
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
            get('/site/info', rst => {
                document.documentElement.lang = rst.lang;
                document.title = `${rst.subTitle}-${rst.title}`;
                dispatch(refresh(rst));
            });
        }
    })
)(Widget);

