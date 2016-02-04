import Ember from 'ember';
import $ from 'jquery';

export default Ember.Component.extend({
  url: '/',
  addedFile: function (file) {
    this.set('fileName', file.name);
  },

  didRender() {
    let dropZoneComponent = this,
        dropZone          = $('#drag_drop');

    if (!dropZone[0].dropzone) {
      $('#drag_drop').dropzone({
        autoProcessQueue: false,
        maxFiles: 1,
        dictDefaultMessage: 'Try dropping some files here, or click to select files to upload.',
        init: function () {
          var submitBtn = $('#drag_drop_upload'),
              resetBtn  = $('#drag_drop_reset'),
              dropZone  = this;

          submitBtn.on('click', function () {
            dropZone.processQueue();
          });

          resetBtn.on('click', function () {
            dropZone.removeAllFiles();
            dropZoneComponent.set('fileName', '');
          });

          this.on('addedfile', dropZoneComponent.addedFile.bind(dropZoneComponent));
        }
      });
    }
  }
});
