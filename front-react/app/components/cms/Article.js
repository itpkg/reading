import React,{PropTypes} from 'react';
import {Link} from 'react-router'
import {Input, Alert} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'
import {connect} from 'react-redux';

import {GET, POST, DELETE, response, failed} from '../../ajax';
import Buttons from '../widgets/FormButtons';
import {CurrentUser} from '../../mixins'
import NoMatch from '../NoMatch'

const Edit = React.createClass({
    mixins: [LinkedStateMixin, CurrentUser],
    getInitialState() {
        return {
            comments: [],
            tags: []
        }
    },
    render(){
        const {item}=this.props;
        return (<form>edit{item.aid}</form>)
    }
});

const ShowW = React.createClass({
    mixins: [LinkedStateMixin, CurrentUser],
    getInitialState() {
        return {
            item: {
                tags: [],
                comments: []
            }
        }
    },
    componentWillReceiveProps(){
        this.reloadArticle();
    },
    componentDidMount() {
        this.reloadArticle();
    },
    reloadArticle(){
        const {aid, user} = this.props;
        GET(
            "/cms/article/" + aid,
            function (article) {
                this.setState({item: article});
            }.bind(this),
            function (e) {
                this.setState({item: {title: e.responseText}});
            }.bind(this)
        )
    },
    render(){
        const {aid, user} = this.props;
        var item = this.state.item;
        if (item.id) {
            return (
                <div>
                    <h3><Link to={`/cms/article/${aid}`}>{item.title}</Link></h3>
                    <hr/>
                    {item.body}
                </div>
            )
        } else {
            if (this.isSignIn(user)) {
                return (<Edit item={{aid:aid}}/>)
            } else {
                return <NoMatch/>
            }
        }

    }
});

ShowW.propTypes = {
    user: PropTypes.object.isRequired,
    aid: PropTypes.string.isRequired
};

export const Show = connect(
    (state, ownProps) => ({
        user: state.current_user,
        aid: ownProps.params.id,
        filter: ownProps.location.query.filter
    }),
    dispatch => ({})
)(ShowW);

export const Index = React.createClass({
    render(){
        return <div>articles </div>;
    }
});


