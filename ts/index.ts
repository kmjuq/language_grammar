let isDone: boolean = false;
let decLiteral: number = 6;
let hexLiteral: number = 0xf00d;
// ES6 中的二进制表示法
let binaryLiteral: number = 0b1010;
// ES6 中的八进制表示法
let octalLiteral: number = 0o744;
let notANumber: number = NaN;
let infinityNumber: number = Infinity;

let myname: string = "111";


function alterName(): void {
    alert('My name is Tom');
}

let myFavoriteNumber: any = 'seven';
myFavoriteNumber = 7;

let myFavoriteNumber1: string | number;
myFavoriteNumber1 = 'seven';
console.log(myFavoriteNumber1.length); // 5
myFavoriteNumber1 = 7;
// console.log(myFavoriteNumber1.length); // 5

interface Person {
    name: string;
    // 可选类型就相当于 age: undefined | number;
    age?: number;
    [propName: string]: string | number | undefined;
}

console.log("测试", isDone) 