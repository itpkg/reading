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
        const {children, siteInfo} = this.props;
        return (
            <div>
                <Header title={siteInfo.subTitle} />
                <div className="container-fluid">
                    <div className="row">
                        {children}
                    </div>
                    <hr/>
                    <Footer copyright={siteInfo.copyright}/>
                </div>
            </div>
        );
    }
}

Layout.propTypes = {
    siteInfo: PropTypes.object.isRequired,
    onRefresh: PropTypes.func.isRequired
};


export default connect(
    state => ({siteInfo:state.site_info}),
    dispatch => ({
        onRefresh: function () {
            GET('/site/info', function (rst) {
                dispatch(refresh(rst));
            })
        }
    })
)(Layout);
