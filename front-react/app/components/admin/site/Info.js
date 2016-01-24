import React, {PropTypes} from 'react';
import i18next from 'i18next/lib';

import Form from '../../widgets/Form'

const Widget = React.createClass({
    render(){
        const {item} = this.props;

        return (
            <Form action="/admin/site/info"
                  fields={[
                    {
                        id:"title",
                        type:"text",
                        label: i18next.t('models.site.title'),
                        value:item.title
                    },
                    {
                        id:"subTitle",
                        type: "text",
                        label: i18next.t('models.site.subTitle'),
                        value:item.subTitle
                    },
                    {
                        id:"authorName",
                        type: "text",
                        label: i18next.t('models.site.author.name'),
                        value:item.author.name
                    },
                    {
                        id:"authorEmail",
                        type: "text",
                        label: i18next.t('models.site.author.email'),
                        value:item.author.email
                    },
                    {
                        id:"keywords",
                        type: "text",
                        label: i18next.t('models.site.keywords'),
                        value:item.keywords
                    },
                    {
                        id:"description",
                        type: "textarea",
                        label: i18next.t('models.site.description'),
                        value:item.description
                    },
                    {
                        id:"copyright",
                        type: "text",
                        label: i18next.t('models.site.copyright'),
                        value:item.copyright
                    }
                ]}
            />
        )
    }
});
export default Widget;