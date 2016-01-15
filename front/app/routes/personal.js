import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      tabs:[
        'aaa',
        'bbb',
        'ccc',
        'ddd'
      ]
    }
  }
});
