module.exports = require("./make-webpack-config")({
    env: 'production',
    apiHost: '/api/v1',
    prerender: true,
    minimize: true
});