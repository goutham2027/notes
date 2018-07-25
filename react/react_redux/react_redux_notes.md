## pluralsight course: Building Applications with React and Redux in ES6
  By Cory House

  reactjsconsulting.com


Test utils: Mocha, React test utils and Enzyme

#### Why Redux?
+ One Store
+ Reduced Boilerplate code than Flux
+ Isomorphic/Universal Friendly
+ Immutable store
+ Hot Reloading
+ Time-travel debugging
+ Smaller API (2k min gzip)

#### Environment Setup
react-slingshot - cookie cutter

https://github.com/coryhouse/pluralsight-redux-starter

Handle
  * Automated testing
  * Linting
  * Minification
  * Bundling
  * JSX compilation
  * ES6 transpilation
  * One command! to all of these.

Libraries used
* React
* React Router
* Redux
* Babel - transpiles es6 (ES2015) to es5
* Babel-polyfill - to transpile all of es6 features.
  Quite large 50k gzip min. Pull only required polyfills in production.
* Webpack - bundler. compile js to single bundle js file
* Mocha - Testing
* ESlint

Hot Reloading
babel-preset-react-hmre

Doesn't reload functional components
Doesn't reload container functions like mapStateToProps

#### npm scripts
- Easy to learn
- Simple
- No extra layer of abstraction
- No dependence on separate plugins
- Simpler debugging
- Better docs

Read: bit.ly/npmvsgulp

#### Webpack - Module Bundle
- Bundles the app for the web
