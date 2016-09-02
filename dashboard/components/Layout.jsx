import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

import i18n from 'i18next'

import {refresh} from '../engines/platform/actions'
import {get} from '../ajax'

import Header from './Header'
import Footer from './Footer'
import MessageBox from './MessageBox'


const Widget = React.createClass({
    getInitialState: function () {
        const {onRefresh} = this.props;
        onRefresh();
        return {};
    },
    render: function () {
        const {children} = this.props;

        //fixme hidden listitem's text
        //<MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>

        return (
            <MuiThemeProvider>
                <div>
                    <Header/>
                    <div className="container">
                        {children}
                    </div>
                    <MessageBox/>
                    <hr className="flaired"/>
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

