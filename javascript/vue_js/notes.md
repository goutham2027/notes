Pluralsight course: Vue.js Getting Started - Chad Campbell
============================================================

Declarative Binding - Glue that hold the UI and data together.
 - Remove the burden of managing the DOM.
 - update happens automatically

Mounting is when virtual dom is actually shown to the user.
Mounting an instance of vue will replace the html dom element with Vue
generated dom.

Life cycle of a View

4 main stages
each of these 4 stages have 2 hooks. hooks provide access to before and after stage.

creation - beforeCreate/created
mounting - beforeMount/mounted
updating - beforeUpdate/updated
destroy - beforeDestroy/destroyed

The creation stage happens when we call the view constructor.

creation
  before crate
  initialize state
  created

    compile template

mounting
  beforemount
  create virtual dom
  mounted

    listen for data changes

updating
  before updating
  re-render virtual dom (apply patches - snapdom algorithm)
  updated

destroy
  before destroy
  teardown virtual dom
  destroyed

