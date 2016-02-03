import DS from 'ember-data';

export default DS.Model.extend({
  content: DS.attr(),
  createdAt: DS.attr('date')
});
