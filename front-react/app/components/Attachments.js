import React, {PropTypes} from 'react';
import DropZone from 'react-dropzone';
import {ListGroup, ListGroupItem, Button, ButtonToolbar} from 'react-bootstrap'
import i18next from 'i18next/lib';
import TimeAgo from 'react-timeago';
import $ from 'jquery';


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
        if (files.size == 0) {
            failed();
            return
        }
        var data = new FormData();
        $.each(files, function (k, v) {
            data.append(k, v);
        });
        UPLOAD("/attachments", response(this.handleReset))
    },
    handleReset(){
        this.setState({files: []})
    },
    render(){
        return (
            <div>
                <DropZone onDrop={this.handleDrop}>
                    <div>{i18next.t('placeholders.upload')}</div>
                </DropZone>
                <br/>
                <ButtonToolbar>
                    <Button bsStyle="primary" onClick={this.handleUpload}>{i18next.t('buttons.upload')}</Button>
                    <Button onClick={this.handleReset}>{i18next.t('buttons.reset')}</Button>
                </ButtonToolbar>
                <br/>
                <ListGroup>
                    {this.state.files.filter(function (item) {
                        return item.type.startsWith('image/');
                    }).map(function (item, idx) {
                        return (<ListGroupItem key={idx}>
                            <img src={item.preview}/>
                        </ListGroupItem>)
                    })}
                </ListGroup>
            </div>
        )
    }
});

export default Widget;