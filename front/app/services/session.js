import Ember from 'ember';

export default Ember.Service.extend({
  token:null,
  init(){
    this.set('token', sessionStorage.getItem('token'));
  },
  update(tkn){
    sessionStorage.setItem('token', tkn);
    this.set('token', tkn);
  },
  clear(){
    sessionStorage.removeItem('token');
    this.set('token', null);
  }

});
