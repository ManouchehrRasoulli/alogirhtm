var value = "Hello World";

function Circular() {
    this.value = value
    this.self = this;
}

var circular = new Circle();

console.log(circular.value);
console.log(circular.self.value);

