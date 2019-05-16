const webpack        = require('webpack');
const path           = require('path');
const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

module.exports = {
  module: {
    rules: [
      {
        include: [path.resolve(__dirname, 'src')],
        loader: 'babel-loader',

        options: {
          plugins: ['syntax-dynamic-import'],

          presets: [
            [
              '@babel/preset-env',
              {
                modules: false
              }
            ]
          ]
        },

        test: /\.js$/
      }
    ]
  },

  entry: {
    moment: path.resolve(__dirname, '') + '/moment.export.js',
    jquery: path.resolve(__dirname, '') + '/jquery.export.js'
  },

  output: {
    path: path.resolve(__dirname, 'libs/'),
    libraryTarget: "this",
    filename: '[name].inc.js'
  },

  // mode: 'development',
  mode: 'production',

  optimization: {
    splitChunks: {
      cacheGroups: {
        vendors: {
          priority: -10,
          test: /[\\/]node_modules[\\/]/
        }
      },

      chunks: 'async',
      minChunks: 1,
      minSize: 30000,
      name: true
    }
  }
};
