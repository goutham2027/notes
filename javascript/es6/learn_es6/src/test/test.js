let path = require('path');
let chai = require('chai');

let Person = require(path.join(__dirname, '..', 'js/person'));

// let expect = chai.expect;
// let should = chai.should;

describe("Test say_hello", function() {
    it("should return a string with the given name", function() {
        p = new Person('Goutham')
        chai.expect(p.getName(), 'Goutham');
    })
});
// var assert = require('assert');
// describe('Array', function() {
//     describe('#indexOf()', function() {
//         it('should return -1 when the value is not present', function() {
//             assert.equal([1, 2, 3].indexOf(4), -1);
//         });
//     });
// });