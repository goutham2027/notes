import '../css/main.css';
import '../css/elements.css';
import { sayHello } from './say_hello';


//  function call_hello() {
//    console.log("Hello function");
//  }
sayHello()

const root = document.createElement("div")
root.innerHTML = `<p>Hello Webpack.</p>`
document.body.appendChild(root)