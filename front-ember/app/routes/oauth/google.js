import Ember from 'ember';

export default Ember.Route.extend({
  urlParser: Ember.inject.service(),  
  session: Ember.inject.service(),
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
           self.get('session').update(tkn);
           window.location.href='/';
        }
    );
  }
});
