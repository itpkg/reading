import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function () {
  this.route('books', function () {
    this.route('show', {path: '/:id'});
  });
  this.route('session', function () {
    this.route('new', {path: '/sign_in'});
    this.route('destroy');
  });
  this.route('users', function () {
    this.route('new', {path: '/sign_up'});
  });
});

export default Router;
