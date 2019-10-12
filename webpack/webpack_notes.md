## Webpack
https://frontendmasters.com/courses/webpack-fundamentals/
https://frontendmasters.com/courses/performance-webpack/

## Books
Javascript Design Patterns

### Why webpack?
- Unmaintainable scripts
- Scope
- Size
- Readability
- Fragility
- Monolith files

IIFE - Immediately Invoked Function Expression
```javascript
const whatever = (function(datanowusedinside){
    return {
        a: "apple"
    }
})(1)
```

Treat each file as one IIFE. Concatenate each files into a single file.

Make, Grunt, Gulp, Broccoli, Brunch, StealJS

Problem with these tools: Full rebuilds everytime.


### Modules

### Static Analysis

NPM + Node + Modules

CommonJS resolution loader is slow.

Bundlers/Linkers

ESM - ECMA script modules
- Reusable
- Encapsulated
installing webpack
- Organized

```
# node version manager
sudo npm install -g n

# to get latest node version
sudo n latest

yarn add webpack
```

References:
https://www.youtube.com/watch?v=GU-2T7k9NfI

### Webpack - How to use it?
#### Config
```javascript
module.exports = [
    entry: {
        vendor: '',
        main: ''
    },
    output: {
        path: '',
        filename: '
    }

]
```
#### CLI

#### Node API
```javascript
var webpack = require("webpack");
webpack({

})
```

Webpack expects entry and output properties, but by default in webpack4 it looks at `src/index.js`


### Webpack Core Concepts
#### 1. Entry
The first javascript file to load to kick-off your app.

#### 2. Output
Tells Webpack Where and how to distribute bundles (compilations). Works with Entry.

```javascript
module.exports = {
    entry: "",
    output: {
        path: './dist',
        filename: './bundle.js'
    },
}
```
#### 3. Loaders + Rules
Tells webpack how to interpret and translate files. Transformed on a per-file basis before adding to the depdendency graph.
Tells webpack how to modify files before its added to dependency graph. Loaders are also javascript modules that takes the source file, and returns it in a modified state.

```javascript
module: {
    rules: [
        {test: /\.ts$/, use: 'ts-loader'},

        {test: /\.js$/, use: 'babel-loader'},

        {test: /\.css$/, use: 'css-loader'},
    ]
}

// rules
module: {
    rules: [
        {
            test: regex,
            use: (Array|String|Function),
            include: RegExp[],
            exclude: RegExp[],
            issuer: (RegExp|String)[],
            enforce: "pre"|"post"
        }
    ]
}

// test: A regular expression that instructs the compiler which files to run the loader against
// use: An array/string/function that returns loader objects
// enforce: tells webpack to run this rule before or after all other rules.
// include: an array of regular expression that instruct the compiler which folders/files to include.
// will only search paths provided with the include.
// exclude: An array of regular expression that instructs the compuler which folders/files to ignore.
```

#### Chaining loaders
Right to left
style(css(less())) = ["style", "css", "less"]

```javascript
rules: [
    {
        test: /\.less$/,
        use: ["style", "css", "less"]
    }
]
```

#### 4. Plugins
Instance that has an apply property in chaining loaders.

Objects with apply property.

Allow you to hook into the entire compilation lifecycle.

Webpack has a variety of built in plugins.

```javascript
var BellOnBundlerErrorPlugin = require('bell-on-error');
var webpack = require('webpack');

module.exports = {
    plugins: [
        new BellOnBundlerErrorPlugin(),

        new webpack.optimize.CommonsChunkPlugin('vendors')
    ]
}

```



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

### Webpack config
```javascript
module.exports = {
    output: {
        filename: "bundle.js"
    }

};
```

To access environment variables, pass `--env.mode`

```
module.exports = (env) => {
    console.log(env);
    return {
    output: {
            filename: "bundle.js"
        }

    };
};
```

### Adding webpack plugins
html-webpack-plugin

It injects  output assets into html file.

webpack-dev-server

wepback-merge

### Using CSS with webpack

### hot module replacement

mincss extract plugin

### File loader and URL loader
yarn add file-loader url-loader

### Limit Filesize option in URL loader
```javascript
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const webpackMerge = require("webpack-merge");

const modeConfig = env => require(`./build-utils/webpack.${env}`)(env);

module.exports = ({ mode, presets } = { mode: "production", presets: [] }) => {
  return webpackMerge(
    {
      mode,
      module: {
          rules: [
              {
                test: /\.jpe?g$/,
                use: [{loader: "url-loader", options: {
                    limit: 5000, // 5000 bytes
                }}],
              }
          ]
      },
      output: {
        filename: "bundle.js"
      },
      plugins: [new HtmlWebpackPlugin(), new webpack.ProgressPlugin()]
    },
    modeConfig(mode)
  );
};
```

### Implementing presets

### Bundle Analyzer Presets
yarn add webpack-bundle-analyzer --dev

### Compression
compression-webpack-plugin

```javascript
const CompressionWebpackPlugin = require("compression-webpack-plugin");
module.exports = () => ({
    plugins: [new CompressionWebpackPlugin()]
})

```

### Code splitting CSS

### Source Maps
https://webpack.js.org/configuration

devtool is the property responsible for creating source maps

Source maps gives the entire directory structure in the browser. Used to
do debug driven development.


loaders vs plugins

Lazy loading = code splitting in webpack

```javascript
const loadFooter = () => import("./footer");
const button = makeButton("A button")
button.addEventListener("click", e => {
  loadFooter().then(m => {

    document.body.appendChild(m.footer)
  })
})
```


For loaders and plugins
github.com/webpack-contrib
