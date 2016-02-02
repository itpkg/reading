import Ember from 'ember';

export function timeAgo(time) {
  return moment((new Date(time)).toISOString()).fromNow();
}

export default Ember.Helper.helper(timeAgo);
