module.exports = require("./make-webpack-config")({
    env: 'production',
    apiHost: '/api/v1',
    css:false,
    prerender: true,
    minimize: true
});