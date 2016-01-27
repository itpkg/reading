import React,{PropTypes} from 'react';
import {Link} from 'react-router'
import {Input, Alert, Button, ListGroup, ListGroupItem, Row, Col, Thumbnail} from 'react-bootstrap'
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'
import {connect} from 'react-redux';
import TimeAgo from 'react-timeago';
import {LinkContainer} from 'react-router-bootstrap';

import {GET, POST, DELETE, response, failed, error} from '../../ajax';
import Buttons from '../widgets/FormButtons';
import {CurrentUser} from '../../mixins'
import Markdown from '../widgets/Markdown'
import NoMatch from '../NoMatch'
import {Select} from './Tag'
import {randomColor} from '../../utils'

const EditW = React.createClass({
    mixins: [LinkedStateMixin, CurrentUser],
    getInitialState() {
        const {aid} = this.props;
        return {aid: aid, tags: []}
    },
    componentWillReceiveProps(state){
        this.getArticle(state.aid);
    },
    componentDidMount() {
        const {aid} = this.props;
        this.getArticle(aid);
    },
    handleReset(){
        const {aid} = this.props;
        this.getArticle(aid);
    },
    handleSubmit(e){
        e.preventDefault();
        var data = this.state;
        if (data.body === '' || data.aid === '' || data.title === '') {
            failed();
            return
        }
        data.tags = data.tags_.split(',').map(function (t) {
            return t.trim()
        });
        POST(
            "/cms/articles",
            data,
            response(function () {
                const {aid} = this.props;
                this.getArticle(aid);
            }.bind(this)),
            error
        );
    },
    getArticle(aid){
        GET(
            "/cms/article/" + aid,
            function (article) {
                article.tags_ = article.tags.map(function (t) {
                    return t.name
                }).join(',');
                this.setState(article);
            }.bind(this),
            function (e) {
                this.setState({
                    id: null,
                    aid: aid,
                    title: e.responseText,
                    summary: '',
                    body: '',
                    tags: [],
                    tags_: ''
                });
            }.bind(this)
        )
    },
    render(){
        const {aid} = this.props;
        if (this.isSignIn()) {
            return (<form onSubmit={this.handleSubmit}>
                <Input id="aid"
                       type="text" label={i18next.t('models.cms.article.aid')}
                       valueLink={this.linkState('aid')}/>
                <Input id="title"
                       type="text" label={i18next.t('models.cms.article.title')}
                       valueLink={this.linkState('title')}/>
                <Input id='summary'
                       type='textarea' label={i18next.t('models.cms.article.summary')}
                       valueLink={this.linkState('summary')}/>

                <Input id='tags'
                       type='text' label={i18next.t('models.cms.article.tags')}
                       valueLink={this.linkState('tags_')}/>

                <Input id='body'
                       type='textarea' label={i18next.t('models.cms.article.body')}
                       valueLink={this.linkState('body')}
                />
                <Buttons onReset={this.handleReset} onDelete={this.handleDelete}/>
            </form>)
        } else {
            return <NoMatch/>
        }
    }
});

EditW.propTypes = {
    user: PropTypes.object.isRequired,
    aid: PropTypes.string.isRequired
};

export const Edit = connect(
    (state, ownProps) => ({
        user: state.current_user,
        aid: ownProps.params.aid
    }),
    dispatch => ({})
)(EditW);


const ShowW = React.createClass({
    mixins: [CurrentUser],
    getInitialState() {
        return {
            item: {
                tags: [],
                comments: []
            }
        }
    },
    componentWillReceiveProps(state){
        this.getArticle(state.aid);
    },
    componentDidMount(){
        const {aid} = this.props;
        this.getArticle(aid);
    },
    getArticle(aid){
        GET(
            "/cms/article/" + aid,
            function (article) {
                this.setState({item: article, error: null});
            }.bind(this),
            function (e) {
                this.setState({
                    item: {id: null},
                    error: e.responseText
                });
            }.bind(this)
        )
    },
    render(){
        const {aid} = this.props;
        var item = this.state.item;
        if (item.id) {
            return (
                <div>
                    <h3><Link to={`/cms/article/${aid}`}>{item.title}</Link></h3>
                    <hr/>
                    <blockquote>
                        {i18next.t('models.cms.article.summary')}: {item.summary}
                        <br/>
                        {i18next.t('models.cms.article.tags')}:
                        {item.tags.map(function (t, i) {
                            return (<span key={i}>
                                &nbsp;
                                <Link to={"/cms/tag/"+t.name} style={randomColor()}>
                                    {t.name}
                                </Link>
                                &nbsp;
                                </span>)
                        })}
                        <footer>
                            {i18next.t('models.cms.article.updated_at')}:
                            &nbsp;
                            <TimeAgo date={item.updated_at}/>
                        </footer>
                    </blockquote>
                    <Markdown body={item.body}/>
                </div>
            );
        } else {
            return (
                <Alert bsStyle="danger">
                    <strong>{aid}</strong>
                    &nbsp;
                    {this.state.error}
                    &nbsp;
                    {new Date().toLocaleString()}
                    &nbsp;
                    {
                        this.isSignIn() ?
                            <Link className="btn btn-primary"
                                  to={`/dashboard/cms/article/${aid}/edit`}>{i18next.t('dashboard.article.new')}</Link> :
                            <Link className="btn btn-primary" to={`/`}>{i18next.t('links.back_to_home')}</Link>
                    }
                </Alert>
            )
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
        aid: ownProps.params.aid
    }),
    dispatch => ({})
)(ShowW);


export const Index = React.createClass({
    getInitialState() {
        return {
            items: []
        }
    },
    componentDidMount(){
        GET('/cms/articles', function (rst) {
            this.setState(rst)
        }.bind(this))
    },
    render(){
        return (<Row>
            {this.state.items.map(function (a, i) {
                return <Col key={i} md={4}>
                    <Thumbnail>
                        <h4>{a.title}</h4>
                        <p>{a.summary}</p>
                        <p>
                            <Link to={"/cms/article/"+a.aid} className="btn btn-primary">
                                {i18next.t('buttons.show')}
                            </Link>
                        </p>
                    </Thumbnail>
                </Col>

            })}
        </Row>);
    }
});

export const LatestBar = React.createClass({
    getInitialState() {
        return {
            items: []
        }
    },
    componentDidMount(){
        const {size}=this.props;
        GET('/cms/articles?size=' + size, function (rst) {
            this.setState(rst)
        }.bind(this))
    },
    render(){
        return (<div>
            <h4>{i18next.t('bars.latest_articles')}</h4>
            <hr/>
            <ul>
                {this.state.items.map(function (a, i) {
                    return (<li key={i}>
                        <Link to={"/cms/article/"+a.aid}>{a.title}</Link>
                    </li>)
                })}
            </ul>
        </div>)
    }
});

LatestBar.propTypes = {
    size: PropTypes.number.isRequired
};

export const Manage = React.createClass({
    getInitialState() {
        return {
            items: []
        }
    },
    componentDidMount(){
        GET('/cms/articles/self', function (rst) {
            this.setState(rst)
        }.bind(this))
    },
    render(){
        return (
            <ListGroup>
                {this.state.items.map(function (item, idx) {
                    return (
                        <LinkContainer key={idx} to={'/dashboard/cms/article/'+item.aid+'/edit'}>
                            <ListGroupItem>
                                {item.title}
                            </ListGroupItem>
                        </LinkContainer>
                    )
                })}
            </ListGroup>
        );
    }
});


