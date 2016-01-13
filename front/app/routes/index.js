import Ember from 'ember';

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  i18n: Ember.inject.service(),
  model() {    
    return this.get('ajax').request('/info?locale='+this.get('i18n').get("locale"));
  },
  afterModel: function(model) {
    document.title = model.title+"-"+model.subTitle;
  }
});
