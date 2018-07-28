console.log('hello world');

// static resolve, reject
// var previousPromise = Promise.resolve(3);
// var promise = Promise.resolve(previousPromise);
// var promise = Promise.catch(previousPromise);

// Nesting promises. Resolve one promise only if other is resolved.
// let promise = new Promise(function(resolve, reject) {
//     resolve(3);
// });

// let newPromise = new Promise(function(resolve, reject) {
//     resolve(promise);
// });

// newPromise.then(function(data) {
//     console.log(data);
// });


// let promise = new Promise(function(resolve, reject) {
//     // resolve(1);
//     reject(Error('oh no!'));
// });

// promise.then(function(data) {
//     console.log(data);
// }, function(error) {
//     console.log(error);
// });
// just to use error mechanism of promise use catch
// promise.catch(function(error) {
//     console.log(error);
// });
// let octal = 0o71;
// console.log(octal);
// Comprehensions
// var numbers = [for (n of [1, 2, 3, 4]) n * n];
// console.log(numbers)

// import {Person} from './person';

// let p = new Person('Goutham');
// console.log(p.getName());

// export let say_hello = (name) => `Hello ${name}`

// Using generators to build iterables
// let numbers = function* () {
//     yield 1;
//     yield 2;
//     yield 3;
// };

// let numbers = function* (start, end) {
//     for (let i = start; i <= end; i++) {
//         console.log(i);
//         yield i;
//     }
// };

// let sum = 0;
// let iterator = numbers(1, 10);
// let next = iterator.next();
// while (!next.done) {
//    sum += next.value;
//    console.log('next');
//    next = iterator.next();
// }

// for (let n of numbers(1, 10)) {
//     sum += n;
//     console.log('next');
// }

// console.log(sum);
// Building an iterable
// class Company {
//     constructor() {
//         this.employees = [];
//     }

//     addEmployees(...names) {
//         this.employees = this.employees.concat(names);
//     }

//     [Symbol.iterator]() {
//         return new ArrayIterator(this.employees);
//     }
// }

// class ArrayIterator {
//     constructor(array) {
//         this.array = array;
//         this.index = 0;
//     }

//     next() {
//         let result = {value: undefined, done: true};
//         if (this.index < this.array.length) {
//             result.value = this.array[this.index];
//             result.done = false;
//             this.index += 1;
//         }
//         return result;
//     }
// }

// let count = 0;
// let company = new Company();
// company.addEmployees('Tim', 'Sue', 'abc');
// for (let employee of company) {
//     count += 1;
//     employee;
// }

// console.log(count);

// Syntactic sugar for iterable

// let numbers = [1, 2, 3, 4];
// for (let i of numbers) {
//     console.log(i);
// }
// Iterables
// let sum = 0;
// let numbers = [1, 2, 3, 4];
// // for loop
// for (let i = 0; i < numbers.length; i++) {
//     sum += numbers[i];
// }

// // for in
// sum = 0;
// for (let i in numbers) {
//     if (i >= 0) {
//         sum += numbers[i];
//     }
// }

// // iterator
// sum = 0;
// let iterator = numbers.values(); // returns an iterator
// let next = iterator.next();
// while (!next.done) {
//     sum += next.value;
//     next = iterator.next();
// }
// console.log(sum);

// Arrow functions
// let add = (x, y) => x + y;

// console.log(add(3, 5));

// class

// class Person {
//     constructor(name) {
//         this._name = name;
//     }

//     get name() {
//         return this._name;
//     }

//     set name(newValue) {
//         this._name = newValue;
//     }

//     doWork() {
//         return 'Parent class doWork';
//     }
// }

// class Employee extends Person {
//   constructor(name, title) {
//       super(name);
//       this._title = title;
//   }

//   get title() {
//       return this._title;
//   }

//   doWork() {
//       return super.doWork() + `${this._name} completed!`;
//   }
// }

// let e = new Employee('Foo', 'Developer');
// console.log(e.name);
// console.log(e.title);
// console.log(e.doWork());

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
