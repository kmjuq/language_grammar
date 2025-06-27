use rand::Rng;

pub fn 函数() {
    // cal(i: i32, j: i32) -> i32 是我们创建的一个函数
    // 有编程语言经验的人应该能看出来函数的语法规范
    // 关键字 fn，方法名cal，参数列表i:i32 j:i32，返回值类型 i32，方法体{}
    // 需要注意的是函数体中的语句和表达式
    // 函数体也是一个表达式，如果不返回任何值，会隐式地返回一个 ()
    语句和表达式();
    println!("cal 33 4 {}", cal(33, 4));
}

pub fn 语句和表达式() {
    // 唯一需要注意的是rust 有语句和表达式的概念
    // 语句就是一个具体的命令，比如 cal 方法中的 let i1 = i * 2;
    // 能返回值的就是一个表达式，比如 i1 / j1 是有计算结果的，所以是表达式
    _ = cal(88, 3);
}

fn cal(i: i32, j: i32) -> i32 {
    let i1 = i * 2;
    let j1 = j + 1;
    i1 / j1
}

pub fn 流程控制() {
    // 条件判断语法
    let mut rng = rand::thread_rng();
    let generator: u8 = rng.gen();
    let number = if generator > 0 {
        1
    } else if generator == 0 {
        0
    } else {
        -1
    };
    println!("The value of number is: {}", number);

    // 循环语法
    // for 循环
    for i in 1..=5 {
        println!("{}", i);
    }
    let arr = [4, 3, 2, 1];
    // 同时获取索引和数据
    for (i, v) in arr.iter().enumerate() {
        println!("第 {} 个元素是{}", i + 1, v);
    }
    // continue 和 break;
    print!("continue break");
    for i in 1..8 {
        if i == 2 {
            continue;
        }
        if i == 7 {
            break;
        }
        print!("{}", i);
    }
    println!();
    // while 循环
    print!("while ");
    let mut while_count = 10;
    while while_count > 0 {
        print!("{}", while_count);
        while_count -= 1;
    }
    println!();
    // loop 循环
    let mut loop_count = 0;
    print!("loop ");
    loop {
        print!("{}", loop_count);
        loop_count += 1;
        if loop_count > 10 {
            break;
        }
    }
    println!();
}

///
/// |使用方法|等价使用方法|所有权|
/// |:--|:--|:--|
/// |for item in collection|for item in IntoIterator::into_iter(collection)|转移所有权|
/// |for item in &collection|for item in collection.iter()|不可变借用|
/// |for item in &mut collection|for item in collection.iter_mut()|可变借用|
pub fn 集合元素循环遍历中的所有权转移() {
    // 转移所有权
    let arr: [String; 3] = std::array::from_fn(|_| String::from("kmj"));
    print!("arr 数组内容为 ");
    IntoIterator::into_iter(arr).for_each(|val| print!("{} ", val));
    println!("");
    // 以下语句无法通过编译，因为 String 的数据所有权已被转移
    // println!("{:?}", arr);

    let arr: [String; 3] = std::array::from_fn(|_| String::from("kmj"));
    print!("arr 数组内容为");
    for val in arr {
        print!("{} ", val)
    }
    println!();
    // 以下语句无法通过编译，因为 String 的数据所有权已被转移
    // println!("{:?}", arr);

    // 不可变借用
    // 变量遮蔽
    let arr: [String; 3] = std::array::from_fn(|_| String::from("kmj"));
    print!("&arr 数组内容为 ");
    for val in &arr {
        print!("{} ", val)
    }
    println!();
    // &arr 的方式不会转移所有权，因此还可以使用
    println!("&arr 数组内容为 {:?}", arr);

    let arr: [String; 3] = std::array::from_fn(|_| String::from("kmj"));
    print!("&arr 数组内容为 ");
    // 这里为啥生成一个值？而上面不会？
    let _ = &arr.iter().for_each(|val| print!("{} ", val));
    println!();
    // &arr 的方式不会转移所有权，因此还可以使用
    println!("&arr 数组内容为 {:?}", arr);

    // 可变借用
    let mut arr: [String; 3] = std::array::from_fn(|_| String::from("kmj"));
    print!("&mut arr 数组内容为 ");
    for val in &mut arr {
        val.push_str("&mut ");
        print!("{} ", val);
    }
    println!();
    // &arr 的方式不会转移所有权，因此还可以使用
    println!("&mut arr 数组内容为 {:?}", arr);
}

pub fn loop循环搭配break的表达式() {
    let mut rng = rand::thread_rng();
    let mut generator: u8 = rng.gen();
    print!("loop ");
    let count = loop {
        print!("{} ", generator);
        if generator >= 10 {
            generator /= 10;
            continue;
        }
        if generator > 0 {
            break 1;
        } else if generator == 0 {
            break 0;
        } else {
            break -1;
        }
    };
    println!();
    println!("loop expr count {}", count);
}

pub fn 方法() {
    struct Person {
        name: String,
    }

    /**
     * impl 关键字加上结构体代表了该代码片段是结构体的函数，而含有self的函数即是方法，无的是关联函数。
     */
    impl Person {
        /**
         *  关联函数，相当于java里面类的静态方法，和结构体实例无关，在rust中一般用来创建结构体实例
         */
        fn new(name: String) -> Person {
            Person { name }
        }

        /**
         * 方法是用来定义结构体行为的函数，一般首个参数为 self,&self,&mut self ，并且该参数不用添加类型，是rust语法糖
         * self         会把当前实例的所有权转移，所以一般用的少。
         * &self        当前结构体的不可变引用
         * &mut self    当前结构体的可变应用
         *
         */
        fn eat(&self) {
            println!("{} 正在吃饭", &self.name);
        }
    }
    let person = Person::new(String::from("kmj"));
    person.eat();
}

pub fn 结构体可以定义多个impl() {
    struct Hero {
        name: String,
        rank: i32,
        attack: i32,
    }

    impl Hero {
        fn new(name: String) -> Hero {
            Hero {
                name,
                rank: 1,
                attack: 1,
            }
        }
    }

    impl Hero {
        fn upgrade(&mut self) {
            self.rank += 1;
            self.attack += 1;
            println!("{} 升级了，目前等级为{}", self.name, self.rank);
        }
    }

    impl Hero {
        fn attack(&self) {
            println!("{} 造成了伤害 {}", self.name, self.attack);
        }
    }

    let mut lunar = Hero::new(String::from("lunar"));
    lunar.upgrade();
    lunar.attack();
}

pub fn 枚举也可以定义impl() {
    enum Creature {
        Hero,
        Monster,
    }

    impl Creature {
        fn idle(&self) {
            println!("空闲站姿动画")
        }
    }

    let hero = Creature::Hero;
    hero.idle();
}
