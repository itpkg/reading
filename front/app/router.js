import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('contact');
  this.route('about');
  //this.route('page-not-found', path: '/*wildcard');
  this.route('faq');
  this.route('not-match', {path: '/*wildcard'});
});

export default Router;
