import Ember from 'ember';

export default Ember.Component.extend({
  actions: {
    removeAttachment(id) {
      if (confirm('Are you sure?')) {
        let attachmentList     = this,
            attachmentToRemove = attachmentList.get('attachments').filterBy('id', id)[0];

        attachmentToRemove.destroyRecord().then(function () {
          let remainedAttachments = attachmentList.get('attachments').filter(function (attachment) {
            return attachment.id !== id;
          });

          attachmentList.set('attachments', remainedAttachments);
        }, function () {
          alert('Error! Record cannot be deleted!');
        });
      }
    }
  }
});
