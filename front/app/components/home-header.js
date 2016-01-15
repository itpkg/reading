import Ember from 'ember';

export default Ember.Component.extend({
  i18n: Ember.inject.service(),
  auth: Ember.inject.service(),
  site: Ember.inject.service('site-info'),
  locales: Ember.computed('i18n.locale', 'i18n.locales', function() {
    const i18n = this.get('i18n');
    return this.get('i18n.locales').map(function (loc) {
      return { id: loc, text: i18n.t('locales.' + loc) };
    });
  }),
  actions: {
    setLocale(locale) {
      localStorage.setItem("locale", locale);
      this.set('i18n.locale', locale);
      this.get('site').init();
    },
    signOut(){
      this.get('auth').sign_out();      
    }
  }
});
