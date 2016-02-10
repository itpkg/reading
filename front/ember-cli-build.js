var EmberApp = require('ember-cli/lib/broccoli/ember-app');

module.exports = function(defaults) {
  var app = new EmberApp(defaults, {
    // Add options here
  });

  app.import('bower_components/moment/moment.js');
  app.import('bower_components/dropzone/dist/dropzone.js');
  app.import('bower_components/dropzone/dist/dropzone.css');

  return app.toTree();
};
