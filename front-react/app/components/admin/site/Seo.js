import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';
import LinkedStateMixin from 'react-addons-linked-state-mixin'
import {Input, ListGroupItem, ListGroup} from 'react-bootstrap'

import {GET, POST, response, failed} from '../../../ajax';
import Buttons from '../../widgets/FormButtons'

const Widget = React.createClass({
    mixins: [LinkedStateMixin],
    getInitialState() {
        return {}
    },
    componentDidMount() {
        this.handleReset();
    },
    handleSubmit(e){
        e.preventDefault();
        POST(
            "/admin/site/seo",
            this.state,
            response(this.handleReset)
        )
    },
    handleReset(e){
        GET('/admin/site/seo', function (rst) {
            this.setState(rst);
        }.bind(this))
    },
    render(){
        var baidu = this.state.baiduVerify;
        var google = this.state.googleVerify;
        return (<div>

            <form method='POST' onSubmit={this.handleSubmit}>

                <Input type="textarea" id="robotsTxt" label={i18next.t('forms.siteSeo.robotsTxt')}
                       valueLink={this.linkState('robotsTxt')}/>
                <Input type="text" id="googleVerify" label={i18next.t('forms.siteSeo.googleVerify')}
                       valueLink={this.linkState('googleVerify')}/>
                <Input type="text" id="baiduVerify" label={i18next.t('forms.siteSeo.baiduVerify')}
                       valueLink={this.linkState('baiduVerify')}/>
                <Buttons onReset={this.handleReset}/>
            </form>
            <ListGroup>
                <ListGroupItem href={API_HOST+'/robots.txt'} target="_blank">robots.txt</ListGroupItem>
                <ListGroupItem href={API_HOST+'/google'+google+'.html'}
                               target="_blank">{'google' + google + '.html'}</ListGroupItem>
                <ListGroupItem href={API_HOST+'/baidu_verify_'+baidu+'.html'}
                               target="_blank">{'baidu_verify_' + baidu + '.html'}</ListGroupItem>
                <ListGroupItem href={API_HOST+'/sitemap.xml'} target="_blank">sitemap.xml</ListGroupItem>
                <ListGroupItem href={API_HOST+'/rss.atom'} target="_blank">rss.atom</ListGroupItem>
            </ListGroup>
            <ul>
                <li>
                    <a href="http://www.robotstxt.org/robotstxt.html" target="_blank">
                        http://www.robotstxt.org/robotstxt.html
                    </a>
                </li>
            </ul>
        </div>)
    }
});
export default Widget;