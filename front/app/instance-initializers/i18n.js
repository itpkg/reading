export function initialize( applicationInstance ) {
  applicationInstance.lookup('service:i18n').set('locale', calculateLocale());
}

function calculateLocale() {
  //console.log(navigator.language +" "+ navigator.userLanguage+" "+localStorage.getItem("locale"));
  return localStorage.getItem("locale") || navigator.language || navigator.userLanguage || 'en-US';
}

export default {
  name: 'i18n',
  initialize
};
