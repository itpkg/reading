import Ember from 'ember';
//import Base64 from '/bower_components/js-base64/base64';

//var Base64 = require('bower_components/js-base64/base64.js').Base64;

export default Ember.Service.extend({
  ajax: Ember.inject.service(),
  current_user: null,
  oauth: null,
  init(){
    var self = this;
    this.get('ajax').request('/oauth/sign_in').then(function(rst){
      self.set('oauth', rst);
    });
    this.refresh();
  },
  refresh(){
    var tkn = sessionStorage.getItem('token');
    console.log(tkn);
    console.log(Base64.decode(tkn));
    //this.set('token', user);
  }
});
