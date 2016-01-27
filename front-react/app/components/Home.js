import React from 'react'

import {Index as Articles} from './cms/Article'
import {Cloud as TagCloud} from './cms/Tag'
import {Bar as NoticeBar} from './Notice'


const Widget = React.createClass({
    render(){
        return (
            <div className="row">
                <div className="col-md-9">
                    <br/>
                    <Articles />
                </div>
                <div className="col-md-3">
                    <TagCloud />
                    <br/>
                    <NoticeBar size={3}/>
                </div>
            </div>
        );
    }
});

export default Widget;
