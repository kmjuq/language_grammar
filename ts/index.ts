let isDone: boolean = false;
let count: number = Infinity;
let myname: string = "111";
let mynull: null = null;
let myundefined: undefined = undefined;
let mysymbol: Symbol = Symbol();
let mybigint: BigInt = BigInt(3432432151243214325324123);

// void 为无返回值类型
function alterName(): void {
    alert('My name is Tom');
}

// any 任意类型，如果变量没有标注类型，则默认为any类型
let myFavoriteNumber: any = 'seven';
myFavoriteNumber = 7;

// 联合类型
let myFavoriteNumber1: string | number;
myFavoriteNumber1 = 'seven';
console.log(myFavoriteNumber1.length); // 5
myFavoriteNumber1 = 7;
// console.log(myFavoriteNumber1.length); // 5

// 对象的类型 --> 接口
interface Person {
    // 只读属性，只能在创建的时候被赋值
    // 注意，只读的约束存在于第一次给对象赋值的时候，而不是第一次给只读属性赋值的时候
    readonly id: number;
    name: string;
    // 可选类型就相当于 age: undefined | number;
    age?: number;
    // 任意属性 一旦定义了任意属性，那么确定属性和可选属性的类型都必须是它的类型的子集
    // 一个接口中只能定义一个任意属性。如果接口中有多个类型的属性，则可以在任意属性中使用联合类型
    [propName: string]: string | number | undefined;
}

// 数组类型
let fibonacci1: number[] = [1, 1, 2, 3, 5];
let fibonacci2: Array<number> = [1, 1, 2, 3, 5];
let list: any[] = ['xcatliu', 25, { website: 'http://xcatliu.com' }];

// 函数
function sum(x: number, y: number): number {
    return x + y;
}

let mySum: (x: number, y: number) => number = function (x: number, y: number): number {
    return x + y;
};

// 当使用接口定义函数签名时，只能定义一个，约束回调函数、高阶函数等
interface SearchFunc {
    (source: string, subString: string): boolean;
}

let mySearch: SearchFunc;
mySearch = function (source: string, subString: string) {
    return source.search(subString) !== -1;
}

// 接口定义了一个对象类型，包含多个方法属性（C, R, U, D）
// 约束对象必须包含这些方法，且每个方法符合对应的签名
// 约束对象的行为（如服务、控制器）
interface CRUD {
    C: (name: string, age: number) => boolean;
    R: () => { id: number, name: string, age: number };
    U: (id: string, name: string, age: number) => boolean;
    D: (id: number) => boolean;
}

// 可选参数
function buildName1(firstName: string = 'm', lastName?: string) {
    // TODO
}
// 给默认值的参数就是可选参数
function buildName2(firstName: string, lastName: string = 'k') {
    // TODO
}
// 剩余参数
function push(array: any[], ...items: any[]) {
    items.forEach(function (item) {
        array.push(item);
    });
}

// 函数重载
function reverse(x: number): number;
function reverse(x: string): string;
function reverse(x: number | string): number | string | void {
    if (typeof x === 'number') {
        return Number(x.toString().split('').reverse().join(''));
    } else if (typeof x === 'string') {
        return x.split('').reverse().join('');
    }
}

