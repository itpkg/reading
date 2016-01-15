import Ember from 'ember';

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
  sign_out(){
    sessionStorage.removeItem('token');
    this.set('current_user', null);
  },
  refresh(){
    var tkn = sessionStorage.getItem('token');
    if(tkn){
      try{
        this.set("current_user", JSON.parse(Base64.decode(tkn.split('.')[1])));
      }catch(e){
        this.sign_out();
      }
    }
  }
});
