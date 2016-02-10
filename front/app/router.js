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
  this.route('notices', function () {
  });
  this.route('dashboard', function () {
    this.route('cms', function () {
      this.route('articles', function () {
        this.route('new');
        this.route('edit', {path: '/:id'});
      })
    });
    this.route('personal', function () {
      this.route('logs');
    });
    this.route('attachments', function () {
      this.route('create');
      this.route('destroy');
    });
  });
});

export default Router;
