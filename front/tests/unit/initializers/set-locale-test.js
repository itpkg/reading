import Ember from 'ember';
import SetLocaleInitializer from '../../../initializers/set-locale';
import { module, test } from 'qunit';

let application;

module('Unit | Initializer | set locale', {
  beforeEach() {
    Ember.run(function() {
      application = Ember.Application.create();
      application.deferReadiness();
    });
  }
});

// Replace this with your real tests.
test('it works', function(assert) {
  SetLocaleInitializer.initialize(application);

  // you would normally confirm the results of the initializer here
  assert.ok(true);
});
