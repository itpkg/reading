var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function (options) {
    var entry = {
        main: path.join(__dirname, 'app', options.env)
    };

    var loaders = [
        {
            test: /\.jsx?$/,
            exclude: /(node_modules)/,
            loader: 'babel',
            query: {
                presets: ['react', 'stage-0', 'es2015']
            }
        },
        {test: /\.css$/, loader: "style!css"},
        {test: /\.json$/, loader: "json"},
        {test: /\.(png|jpg|jpeg|gif|svg|ttf|woff|woff2|eot)$/, loader: "file-loader"}
    ];

    var plugins = [
        new webpack.ProvidePlugin({
            //fix 'jQuery is not defined' bug
            //$: "jquery",
            //jQuery: "jquery"
        })
    ];

    var htmlOptions = {
        title: 'reading',
        favicon: 'favicon.ico',
        inject: true,
        template: 'app/index.html'
    };

    if (options.minimize) {
        htmlOptions.minify = {
            collapseWhitespace: true,
            removeComments: true
        };

        plugins.push(new webpack.optimize.UglifyJsPlugin({
            output: {
                comments: false
            }
        }));

        plugins.push(new webpack.DefinePlugin({
            "process.env": {
                NODE_ENV: JSON.stringify("production")
            }
        }));
        plugins.push(new webpack.NoErrorsPlugin());

    }
    plugins.push(new HtmlWebpackPlugin(htmlOptions));
    plugins.push(new webpack.DefinePlugin({
        VERSION: JSON.stringify('v0.0.1'),
        API_HOST: JSON.stringify(options.apiHost),
        'process.env.NODE_ENV': JSON.stringify(options.env)
    }));

    var output = {
        path: path.join(__dirname, 'dist'),
        filename: options.prerender ? "[id]-[chunkhash].js" : '[name].js'
    };

    return {
        entry: entry,
        output: output,
        plugins: plugins,
        module: {
            loaders: loaders
        },
        devServer: {
            inline: true,
            port: 4200,
            proxy: {
                "*": "http://127.0.0.1:3000/"
            }
        }
    }
};
