Webpack is a build tool but not limited to that.

installing webpack

```
# node version manager
sudo npm install -g n

# to get latest node version
sudo n latest

yarn add webpack
```

References:
https://www.youtube.com/watch?v=GU-2T7k9NfI


### Using the webpack dev server
Disadvantages of File protocol:
* Doesn't send files on HTTP protocol
```
yarn add -D webpack-dev-server
```

### Webpack Core Concepts
4 core concepts when using Webpack

1) Entry point: Where should it start
2) Output: Where to store the bundle/bundles
3) Module Loaders:
4) Plugins:

config file
`webpack-config.js`

It should export JS object.
```
module.exports  = {

}
```

### modules and loaders
will help to transform the files.

module property for single loader

use property for multiple loaders

loaders are loaded in reverse order

[style-loader, css-loader]


### plugins
same as loaders and also they are different.
- loaders are used for per file basis
- plugins are ran on the bundle before it is sent for final output.
eg: minifaction

`sass-loader node-sass css-loader extract-text-webpack-plugin babel-core babel-loader babel-preset-es2015`

https://www.dropbox.com/s/injctyvdhgtnq60/Screenshot%202018-07-12%2001.00.43.png?dl=0
