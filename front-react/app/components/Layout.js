import React, { Component, PropTypes } from 'react';
import {connect} from 'react-redux';

import {refresh} from '../actions/site'
import {GET} from '../ajax'
import Footer from './widgets/Footer'
import Header from './widgets/Header'


class Layout extends Component {
    componentDidMount() {
        const {onRefresh} = this.props;
        onRefresh();
    }

    render() {
        const {children} = this.props;
        return (
            <div>
                <Header />
                <div className="container-fluid">
                    <div className="row">
                        {children}
                    </div>
                    <hr/>
                    <Footer/>
                </div>
            </div>
        );
    }
}

Layout.propTypes = {
    onRefresh: PropTypes.func.isRequired
};


export default connect(
    state => ({}),
    dispatch => ({
        onRefresh: function () {
            GET('/site/info', function (rst) {
                dispatch(refresh(rst));
            })
        }
    })
)(Layout);
