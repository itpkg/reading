import Ember from 'ember';

export default Ember.Route.extend({
  actions: {
    redirectTo() {
      this.transitionTo('/');
    }
  }
})
;
