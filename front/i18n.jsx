import i18n from 'i18next'
import LanguageDetector from 'i18next-browser-languagedetector'

import {LOCALE} from './constants'
import resources from './locales'

console.log(resources);

i18n
    .use(LanguageDetector)
    .init({
            fallbackLng: 'en',
            resources: resources,
            detection: {
                order: ['querystring', 'localStorage', 'cookie', 'navigator'],
                lookupQuerystring: LOCALE,
                lookupCookie: LOCALE,
                lookupLocalStorage: LOCALE,

                caches: ['localStorage', 'cookie'],
                cookieMinutes: 365 * 24 * 60
            }
        }
    );

export default i18n;