require('react-select/dist/react-select.css');

import React,{PropTypes} from 'react';
import {connect} from 'react-redux';
import SelectInput from 'react-select'
import i18next from 'i18next/lib';
import {TagCloud, DefaultRenderer} from "react-tagcloud";
import {routeActions} from 'react-router-redux'
import {Row, Col, Thumbnail} from 'react-bootstrap'
import {LinkContainer} from 'react-router-bootstrap';
import {Link} from 'react-router'
import $ from 'jquery'

import {GET} from '../../ajax'
import {randomColor} from '../../utils'

//const renderer = new DefaultRenderer({
//    props: {
//        onClick: (e) => routeActions.push('/cms/tags/'+$(e.target).text())
//    }
//});
const renderer = (tag, size, key) => {
    return (
        <span key={key}>
            &nbsp;
            <Link style={randomColor()} to={'/cms/tag/'+tag.value}>{tag.value}</Link>
            &nbsp;
        </span>
    );
};

export const Cloud = React.createClass({
    getInitialState() {
        return {
            items: []
        }
    },
    componentDidMount(){
        GET('/cms/tags', function (tags) {
            var items = tags.items.map(function (t) {
                return {value: t.name}
            });
            this.setState({items: items})
        }.bind(this))
    },
    render(){
        return (
            <div>
                <h4>{i18next.t('bars.hot_tags')}</h4>
                <hr/>
                <TagCloud tags={this.state.items} minSize={12} maxSize={35} renderer={renderer}/>
            </div>
        )
    }
});
//todo
export const Select = React.createClass({
    getInitialState() {
        return {
            options: [
                {value: 1, label: 'One'},
                {value: 2, label: 'Two'},
                {value: 11, label: 'a1'},
                {value: 12, label: 'a2'},
                {value: 13, label: 'a3'},
                {value: 14, label: 'a4'}
            ]
        }
    },
    getOptions(input){
        console.log(input);
        GET(
            "/cms/tags",
            function (rst) {
                return {
                    options: [
                        {value: input, label: input}
                    ]
                }
            }.bind(this),
            function (e) {
                return {
                    options: [
                        {value: 'default', label: 'default'}
                    ]
                }
            }.bind(this)
        );
    },
    render(){
        return (
            <div className="form-group">
                <label>{i18next.t('models.cms.article.tags')}</label>
                <SelectInput.Async multi={true} id={this.props.id} loadOptions={this.getOptions}/>
            </div>

        )
    }
});

const ShowW = React.createClass({
    getInitialState() {
        return {
            articles: []
        }
    },
    componentWillReceiveProps(state){
        this.getTag(state.name);
    },
    componentDidMount(){
        const {name} = this.props;
        this.getTag(name);

    },
    getTag(name){
        GET(
            "/cms/tag/" + name,
            function (tag) {
                this.setState(tag);
            }.bind(this)
        )
    },
    render(){
        return (
            <div>
                <h3>{this.state.name}</h3>
                <hr/>

                <Row>
                    {this.state.articles.map(function (a, i) {
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
                </Row>

            </div>
        );
    }
});


ShowW.propTypes = {
    name: PropTypes.string.isRequired
};

export const Show = connect(
    (state, ownProps) => ({
        name: ownProps.params.name
    }),
    dispatch => ({})
)(ShowW);




