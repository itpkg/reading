import React, {PropTypes} from 'react';
import DropZone from 'react-dropzone';
import {ListGroup, ListGroupItem, Button, Image, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';
import TimeAgo from 'react-timeago';
import {connect} from 'react-redux';
import $ from 'jquery';

import RemoveButton from './widgets/RemoveButton'
import {GET, UPLOAD, DELETE, failed, response} from '../ajax';

const Widget = React.createClass({
    getInitialState() {
        return {
            files: [],
            data: {
                items: []
            }
        }
    },
    componentDidMount() {
        this.reloadAttachments();
    },
    reloadAttachments(){
        GET(
            '/attachments',
            function (rst) {
                this.setState({data: rst});
            }.bind(this)
        );
    },
    handleDrop(files){
        this.setState({files: files});
    },
    handleUpload(e){
        e.preventDefault();
        var files = this.state.files;
        if (files.length === 0) {
            failed();
            return
        }
        var data = new FormData();
        $.each(files, function (k, v) {
            data.append(k, v);
        });

        UPLOAD("/attachments", data, response(this.handleReset))
    },
    handleReset(){
        this.setState({files: []});
        this.reloadAttachments();
    },
    render(){
        const {user}=this.props;
        var self = this;
        return (
            <div>
                <div className="row">
                    <div className="col-md-3">
                        <DropZone onDrop={this.handleDrop}>
                            <div>{i18next.t('placeholders.upload')}</div>
                        </DropZone>
                        <br/>
                        <ButtonToolbar>
                            <Button bsStyle="primary" onClick={this.handleUpload}>{i18next.t('buttons.upload')}</Button>
                            <Button onClick={this.handleReset}>{i18next.t('buttons.reset')}</Button>
                        </ButtonToolbar>
                    </div>


                    {this.state.files.map(function (item, idx) {
                        return (<div className="col-md-3" key={idx}>
                            {
                                item.type.startsWith('image/') ?
                                    <Image src={item.preview} thumbnail/> :
                                    item.name
                            }
                        </div>)
                    })}
                </div>
                <hr/>
                <ListGroup>
                    {this.state.data.items.map(function (item, idx) {
                        return (
                            <blockquote key={idx}>
                                <a target="_blank" href={item.url}>{item.title}</a>
                                <footer>
                                    <TimeAgo date={item.created_at}/>
                                    &nbsp;
                                    <cite>
                                        <RemoveButton size='xsmall' action={"/attachment/"+item.id}
                                                      onRefresh={self.reloadAttachments}/>
                                    </cite>
                                </footer>
                            </blockquote>
                        );
                    })}
                </ListGroup>
            </div>
        )
    }
});


Widget.propTypes = {
    user: PropTypes.object.isRequired
};

export  default connect(
    state => ({user: state.current_user}),
    dispatch => ({})
)(Widget);