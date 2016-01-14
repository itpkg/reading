import Ember from 'ember';

export default Ember.Service.extend({
  i18n: Ember.inject.service(),
  ajax: Ember.inject.service(),
  subTitle:null,
  copyright:null,
  init(){
    var self = this;
    this.get('ajax').request('/site/info?locale='+this.get('i18n.locale')).then(function(rst){
      document.title = rst.subTitle+'-'+rst.title;
      self.set('subTitle',rst.subTitle);
      self.set('copyright', rst.copyright);
    });
  }
});
