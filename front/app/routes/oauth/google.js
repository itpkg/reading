import Ember from 'ember';

export default Ember.Route.extend({
  urlParser: Ember.inject.service(),
  auth: Ember.inject.service(),
  ajax: Ember.inject.service(),
  redirect(){
    var self = this;
    this.get('ajax').post(
      '/oauth/sign_in',{
        data: {
          type:'google',
          code:this.get('urlParser').parameter("code")
        }
      }).then(
        function(tkn){
          sessionStorage.setItem('token',tkn);            
          self.get('auth').refresh();
          self.transitionTo('index');
        }
    );
  }
});
