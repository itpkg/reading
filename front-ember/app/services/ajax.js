import Ember from 'ember';
import AjaxService from 'ember-ajax/services/ajax';
import config from '../config/environment';

export default AjaxService.extend({
  host: config.apiHost,
  session: Ember.inject.service(),
  headers: Ember.computed({
    get(){
      let headers={};
      var tkn = this.get('session').token;
      if(tkn){        
        headers['Authorization'] = 'Bearer ' + tkn;
      }
      return headers;
    }
  })
});
