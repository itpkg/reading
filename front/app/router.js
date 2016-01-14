import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('not-match', {path: '/*wildcard'});

  this.route('cms', function() {
    this.route('articles', function() {
      this.route('show', { path: '/show/:aid' });
      this.route('edit', { path: '/edit/:id' });
    });
  });

  this.route('video', function() {
    this.route('items', function() {
      this.route('show', { path: '/show/:id' });
    });
  });

  this.route('books', function() {
    this.route('show', { path: '/show/:id' });
  });
});

export default Router;
