pub fn 基本使用场景() {
    // 模式匹配功能就是方便我们从某些数据结构中取数的
    // 最常使用的是 match 表达式以及 if let.
    // match
    #[derive(Debug)]
    enum Direction {
        East,
        West,
        North,
        South,
    }
    let dire = Direction::South;
    // match 匹配必须穷举出所有可能性，可以用 _ 来代表未列出的所有可能性
    // 每个分支都必须是一个表达式，且所有表达式的返回值必须相同
    match dire {
        Direction::East => println!("East"),
        // "｜" 符号代表"或"
        Direction::North | Direction::South => {
            println!("match dire Direction::South");
        }
        // _ 代表以上没有列举出来的其他匹配项，也可以使用变量来接收
        _ => println!("West"),
    };

    // if let
    // 如果碰到只有一个模式需要匹配，其他值忽略的场景可以使用 if let
    let some_u8 = Some(32u8);
    if let Some(uu8) = some_u8 {
        println!("if let Some(uu8) = some_u8 {:?}", uu8);
    }

    // while let
    // Vec是动态数组
    let mut stack = vec!["kmj", "jmk", "kkk"];
    print!("while let ");
    // stack.pop从数组尾部弹出元素
    while let Some(top) = stack.pop() {
        print!("{} ", top);
    }
    println!();

    // for 循环
    let v = vec!['a', 'b', 'c'];
    for (index, value) in v.iter().enumerate() {
        println!("{} is at index {}", value, index);
    }

    //
}

pub fn matches宏的使用() {
    // 可以将一个表达式跟模式进行匹配，然后返回匹配的结果 true or false。
    #[derive(Debug)]
    enum MyEnum {
        Foo,
        Bar,
        Zoo,
    }

    let v = vec![MyEnum::Foo, MyEnum::Bar, MyEnum::Zoo];
    let val = v.iter().filter(|x| matches!(x, MyEnum::Foo | MyEnum::Bar));
    println!("matches!宏的使用 {:?}", val);
}

pub fn 全模式列表() {
    匹配字面量();
    匹配命名变量();
    单分支多模式();
    通过序列匹配值的范围();
    解构();
}

pub fn 匹配字面量() {
    let x = 1;
    match x {
        1 => println!("one"),
        2 => println!("two"),
        3 => println!("three"),
        _ => println!("anything"),
    }
}

pub fn 匹配命名变量() {
    let x = Some(5);
    let y = 10;
    match x {
        Some(50) => println!("Got 50"),
        Some(y) => println!("Matched, y = {:?}", y),
        _ => println!("Default case, x = {:?}", x),
    }
    println!("at the end: x = {:?}, y = {:?}", x, y);
}

pub fn 单分支多模式() {
    let x = 1;
    match x {
        1 | 2 => println!("one or two"),
        3 => println!("three"),
        _ => println!("anything"),
    }
}

pub fn 通过序列匹配值的范围() {
    let x = 5;
    match x {
        1..=5 => println!("one through five"),
        _ => println!("something else"),
    }
}

pub fn 解构() {
    // 解构结构体
    struct Point {
        x: i32,
        y: i32,
    }
    let p = Point { x: 0, y: 7 };

    let Point { x, y } = p;
    assert_eq!(0, x);
    assert_eq!(7, y);

    // 解构枚举
    enum Color {
        Rgb(i32, i32, i32),
        Hsv(i32, i32, i32),
    }
    enum Message {
        Quit,
        Move { x: i32, y: i32 },
        Write(String),
        ChangeColor(Color),
    }
    let msg = Message::ChangeColor(Color::Hsv(0, 160, 255));
    match msg {
        Message::ChangeColor(Color::Rgb(r, g, b)) => {
            println!("Change the color to red {}, green {}, and blue {}", r, g, b)
        }
        Message::ChangeColor(Color::Hsv(h, s, v)) => {
            println!(
                "Change the color to hue {}, saturation {}, and value {}",
                h, s, v
            )
        }
        _ => (),
    }

    // 解构结构体和元组
    let ((feet, inches), Point { x, y }) = ((3, 10), Point { x: 3, y: -10 });
    println!("解构结构体和元组 {} {} {} {}", feet, inches, x, y);

    // 解构数组
    // 解构定长数组
    let arr: [u16; 2] = [114, 514];
    let [x, y] = arr;
    assert_eq!(x, 114);
    assert_eq!(y, 514);
    // 解构不定长数组，注意引用符号 &
    let arr: &[u16] = &[114, 514];
    if let [x, ..] = arr {
        assert_eq!(x, &114);
    }
    if let &[.., y] = arr {
        assert_eq!(y, 514);
    }
}

pub fn 忽略剩余值() {
    struct Point {
        x: i32,
        y: i32,
        z: i32,
    }
    let origin = Point { x: 0, y: 0, z: 0 };
    match origin {
        Point { x, .. } => println!("x is {}", x),
    }

    let numbers = (2, 4, 8, 16, 32);
    match numbers {
        (first, .., last) => {
            println!("Some numbers: {}, {}", first, last);
        }
    }
}

pub fn 匹配护卫() {
    let x = Some(5);
    let y = 10;
    match x {
        Some(50) => println!("Got 50"),
        Some(n) if n == y => println!("Matched, n = {}", n),
        _ => println!("Default case, x = {:?}", x),
    }
    println!("at the end: x = {:?}, y = {}", x, y);

    let x = 4;
    let y = false;
    match x {
        4 | 5 | 6 if y => println!("yes"),
        _ => println!("no"),
    }
}

pub fn 解构时绑定变量() {
    // 当你既想要限定分支范围，又想要使用分支的变量时，就可以用 @ 来绑定到一个新的变量上。
    enum Message {
        Hello { id: i32 },
    }

    let msg = Message::Hello { id: 5 };

    match msg {
        Message::Hello {
            id: id_variable @ 3..=7,
        } => {
            println!("Found an id in range: {}", id_variable)
        }
        Message::Hello { id: 10..=12 } => {
            println!("Found an id in another range")
        }
        Message::Hello { id } => {
            println!("Found some other id: {}", id)
        }
    }

    #[derive(Debug)]
    struct Point {
        x: i32,
        y: i32,
    }
    // 绑定新变量 `p`，同时对 `Point` 进行解构
    let p @ Point { x: px, y: py } = Point { x: 10, y: 23 };
    println!("x: {}, y: {}", px, py);
    println!("{:?}", p);

    let point = Point { x: 10, y: 5 };
    if let p @ Point { x: 10, y } = point {
        println!("x is 10 and y is {} in {:?}", y, p);
    } else {
        println!("x was not 10 :(");
    }

    match 1 {
        num @ (1 | 2) => {
            println!("{}", num);
        }
        _ => {}
    }
}
