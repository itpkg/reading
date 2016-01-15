import Ember from 'ember';

export default Ember.Component.extend({
  site: Ember.inject.service('site-info')
});
