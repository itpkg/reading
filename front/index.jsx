import injectTapEventPlugin from 'react-tap-event-plugin'
// Needed for onTouchTap
// http://stackoverflow.com/a/34015469/988941
injectTapEventPlugin();

console.log("react version: " + React.version);
console.log("front version: " + process.env.CONFIG.version);


import React from 'react'
import ReactDOM from 'react-dom'

import App from './App'


ReactDOM.render(<App />, document.getElementById('root'));
