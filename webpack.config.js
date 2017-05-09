var path = require('path');
var webpack = require('webpack');

module.exports = {
  entry: ['babel-polyfill', './ui/js/main.js'],
  output: {
    path: path.join(__dirname, '/static/js'),
    filename: '[name].bundle.js'
  },
  plugins: [
    new webpack.ProvidePlugin({
      'Promise': 'imports-loader?this=>global!exports-loader?global.Promise!es6-promise',
      'fetch': 'imports-loader?this=>global!exports-loader?global.fetch!whatwg-fetch'
    })
  ],
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: [/node_modules/],
        use: [{
          loader: 'babel-loader'
        }]
      }
    ]
  }
};
