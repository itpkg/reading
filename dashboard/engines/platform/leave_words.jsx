import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {browserHistory} from 'react-router'

import i18n from 'i18next'

import Form from '../../components/Form'

export const NewFm = ()=>(<Form
    fields={[{
        id: 'content',
        type:'textarea',
        label: '',
        value: '',
    }]}
    title={i18n.t('platform.auth.leave_a_message')}
    method="post"
    action="/leave_words"/>);

