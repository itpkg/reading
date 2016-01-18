import Ember from 'ember';

export default Ember.Route.extend({
  model(params) {
    // multiple models
    // return Ember.RSVP.hash({
    //   songs: this.store.findAll('song'),
    //   albums: this.store.findAll('album')
    // });

    return this.store.findRecord('book', params.id);
  }

  // render customized template
  // renderTemplate() {
  //   this.render('book');
  // }
});
