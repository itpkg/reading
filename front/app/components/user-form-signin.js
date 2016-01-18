import Ember from 'ember';

export default Ember.Component.extend({
  session: Ember.inject.service(),
  loginFailure: false,

  actions: {
    authenticate: function () {
      var userForm = this;

      let credentials = userForm.getProperties('identification', 'password');
      userForm.get('session').authenticate('authenticator:jwt', credentials)
        .then(function () {
          userForm.sendAction('redirectTo', '/');
        }, function () {
          userForm.set('loginFailure', true);
        });
    }
  }
});
