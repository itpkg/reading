import React, {PropTypes} from 'react';
import {connect} from 'react-redux';
import i18next from 'i18next/lib';

function Footer({copyright}) {
    return (
        <footer>
            <p>
                {copyright}
                &nbsp;
                <span
                    dangerouslySetInnerHTML={{__html: i18next.t('build_using', {url:'https://github.com/itpkg/reading'})}
}/>
            </p>
        </footer>
    )

}

Footer.propTypes = {
    copyright: PropTypes.string.isRequired
};

export default connect(
    state => ({
        copyright: state.site_info.copyright
    }),
    dispatch => ({})
)(Footer);