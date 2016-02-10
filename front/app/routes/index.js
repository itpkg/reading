import Ember from 'ember';

export default Ember.Route.extend({
  model() {
    return Ember.RSVP.hash({
      notices: this.store.findAll('notice')
    });
  }
});
