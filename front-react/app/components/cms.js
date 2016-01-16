import React from 'react';


export const Article = React.createClass({
    render(){ //todo
        return (<div>
            {this.props.params.aid}
        </div>)
    }
});