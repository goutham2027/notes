console.log('hello world');

class Person {
    constructor(name) {
        this._name = name;
    }

    get name() {
        return this._name;
    }

    set name(newValue) {
        this._name = newValue;
    }

    doWork() {
        return 'Parent class doWork';
    }
}

class Employee extends Person {
  constructor(name, title) {
      super(name);
      this._title = title;
  }

  get title() {
      return this._title;
  }

  doWork() {
      return super.doWork() + `${this._name} completed!`;
  }
}

let e = new Employee('Foo', 'Developer');
console.log(e.name);
console.log(e.title);
console.log(e.doWork());

// class Employee {
//     constructor(name) {
//         this._name = name;
//     }
//     doWork() {
//         return "complete!";
//     }

//     // getName() {
//     //     return this.name
//     // }

//     get name() {
//         return this._name
//     }

//     set name(newValue) {
//         this._name = newValue;
//     }
// }

// var e = new Employee("Foo");
// var result = e.doWork();
// var result = e.getName();
// console.log(e.name);
// console.log("Setting new value");
// e.name = "Bar";
// console.log(e.name);


// spread operators
// let sum = function(...numbers) {
//     let result = 0;
//     numbers.forEach(function(n) {
//         result += n;
//     })

//     console.log(result)
// };

// sum(1,21,3)
