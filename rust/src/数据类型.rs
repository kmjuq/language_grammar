pub fn 基础类型() {
    // i8 表示是有符号类型，正号负号。u8 表示无符号类型，范围比i8多一倍
    // 位数有 8 16 32 64 128 位
    let uu8 = u8::MAX;
    let ii8 = i8::MIN;
    println!("{} {}", ii8, uu8);
    let uu128 = u128::MAX;
    let ii128 = i128::MIN;
    println!("{} {}", ii128, uu128);
    // 还有个动态的 isize 和 usize，该类型和CPU有关。
    // 32位cpu，则isize表示32位。
    // 64位cpu，则isize表示64位。
    let uusize = usize::MAX;
    let iisize = isize::MIN;
    println!("{} {}", iisize, uusize);

    // 数值类型浮点型字面量默认 f64
    let xf64 = 1.0;
    println!("{}", xf64);

    // 字符类型，使用单引号 '' 表示，是unicode编码规范
    let 中 = '中';
    println!("字符'中'占用了{}字节的内存大小", std::mem::size_of_val(&中));

    // 布尔类型
    let tt = true;
    let ff = false;
    if tt && !ff {
        println!("这是段毫无意义的代码 tt {} ff {}", tt, ff);
    }

    // 单元类型
    // 单元类型就是 ()，唯一的值也是 ()，没有指定返回值的函数默认返回 单元类型
}

pub fn 数值运算() {
    // + - * / % 都是其他语言中有的
    // 编译器会进行自动推导，给予twenty i32的类型
    let twenty = 20;
    // 类型标注
    let twenty_one: i32 = 21;
    // 通过类型后缀的方式进行类型标注：22是i32类型
    let twenty_two = 22i32;

    // 只有同样类型，才能运算
    let addition = twenty + twenty_one + twenty_two;
    println!(
        "{} + {} + {} = {}",
        twenty, twenty_one, twenty_two, addition
    );

    // 对于较长的数字，可以用_进行分割，提升可读性
    let one_million: i64 = 1_000_000;
    println!("{}", one_million.pow(2));

    // 定义一个f32数组，其中42.0会自动被推导为f32类型
    let forty_twos = [42.0, 42f32, 42.0_f32];

    // 打印数组中第一个值，并控制小数位为2位
    println!("{:.2}", forty_twos[0]);
}

pub fn 位运算() {
    // & | ^ ! << >> 都是其他语言中有的
    // 二进制为00000010
    let a: i32 = 2;
    // 二进制为00000011
    let b: i32 = 3;

    println!("(a & b) value is {}", a & b);
    println!("(a | b) value is {}", a | b);
    println!("(a ^ b) value is {}", a ^ b);
    println!("(!b) value is {} ", !b);
    println!("(a << b) value is {}", a << b);
    println!("(a >> b) value is {}", a >> b);

    let mut a = a;
    // 注意这些计算符除了!之外都可以加上=进行赋值 (因为!=要用来判断不等于)
    a <<= b;
    println!("(a << b) value is {}", a);
}

pub fn 数字类型不同进制表示() {
    // 数值类型字面量的不同表示。
    // 数值类型整型字面量默认 i32
    let bit2 = 0b1111_0000;
    let bit8 = 0o77;
    let bit10 = 98_222;
    let bit16 = 0xff;
    println!("{} {} {} {}", bit2, bit8, bit10, bit16);
}

pub fn 浮点数精度问题() {
    // 因为默认精度很高，因此更容易碰到精度问题
    let abc: (f32, f32, f32) = (0.1, 0.2, 0.3);
    let xyz: (f64, f64, f64) = (0.1, 0.2, 0.3);

    println!("abc (f32)");
    println!("   0.1 + 0.2: {:x}", (abc.0 + abc.1).to_bits());
    println!("         0.3: {:x}", (abc.2).to_bits());
    println!();

    println!("xyz (f64)");
    println!("   0.1 + 0.2: {:x}", (xyz.0 + xyz.1).to_bits());
    println!("         0.3: {:x}", (xyz.2).to_bits());
    println!();

    assert!(abc.0 + abc.1 == abc.2);
    assert!(xyz.0 + xyz.1 != xyz.2);
}

#[allow(non_snake_case)]
pub fn NaN() {
    // 处理数学上未定义的行为
    let x = (-42.0_f32).sqrt();
    if x.is_nan() {
        println!("未定义的数学行为")
    }
    // 必须指明f32类型，f64无is_nan方法
    let y: f32 = 1.0 / 0.0;
    if y.is_nan() {
        println!("未定义的数学行为")
    }
}

#[allow(non_snake_case)]
pub fn 使用As完成类型转换() {
    // 最常用于将原始类型转换为其他原始类型
}

pub fn 复杂类型() {
    // 元组 元组一般用来组织联系比较紧密的数据，比如二位坐标系的x,y、颜色空间的rgb三原色、IP地址等业务场景
    let coordinate = (10, 25);
    let rgb = (255, 255, 255);
    let ip = (127, 0, 0, 1);
    println!("坐标系{:?} 颜色空间{:?} ip地址{:?}", coordinate, rgb, ip);
    println!(
        "坐标系的 x 为 {}，坐标系的 y 为 {}",
        coordinate.0, coordinate.1
    );

    // 结构体 代表了一个抽象实体的数据集合
    // 结构体的定义
    #[derive(Debug)]
    struct User {
        active: bool,
        username: String,
        email: String,
        sign_in_count: u64,
    }
    // 创建实例
    let user = User {
        email: String::from("kemengjian@126.com"),
        username: String::from("kmj"),
        active: true,
        sign_in_count: 1,
    };
    println!("User 实例为 {:?}", user);
    println!("User 实例的邮箱为 {}", user.email);
    // 元组结构体
    // 主要用来定义已经是行业规范结构的结构体，比如说 RGB 3D空间坐标
    // RGB
    #[derive(Debug)]
    struct Color(i32, i32, i32);
    let rgb_tuple_struct = Color(255, 255, 255);
    println!(
        "元组结构体 struct Color(i32, i32, i32) ; {:?}",
        rgb_tuple_struct
    );
    // 3D 空间坐标
    struct Point(i32, i32, i32);
    // 单元结构体
    // 主要用来定义行为特征，无自己的数据。
    struct AlwaysEqual;

    // 枚举
    #[derive(Debug)]
    enum PokerSuit {
        Clubs,
        Spades,
        Diamonds,
        Hearts,
    }
    #[derive(Debug)]
    struct PokerCard {
        suit: PokerSuit,
        value: u8,
    }
    let card_a = PokerCard {
        suit: PokerSuit::Clubs,
        value: 1,
    };
    let card_q = PokerCard {
        suit: PokerSuit::Diamonds,
        value: 12,
    };
    println!("enum cardA {:?} , enum cardQ {:?}", card_a, card_q);
    // 可以用于不同的结构体归纳于一个枚举中
    #[derive(Debug)]
    enum Message {
        Quit,
        Move { x: i32, y: i32 },
        Write(String),
        ChangeColor(i32, i32, i32),
    }
    let m1 = Message::Quit;
    let m2 = Message::Move { x: 1, y: 1 };
    let m3 = Message::ChangeColor(255, 255, 0);
    println!("m1 {:?};m2 {:?};m3 {:?}", m1, m2, m3);

    // 数组类型
    // 直接初始化
    let months = [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    ];
    // 重复数据的初始化
    let num_arr = [3; 10];
    println!("months {:?}", months);
    println!("numArr {:?}", num_arr[3]);
}

pub fn 通过参数创建结构体的简化语法() {
    // 当参数和结构体字段同名时，可以使用简化语法
    struct Person {
        name: String,
        age: u8,
        height: u8,
        weight: u8,
    }

    fn build_user(name: String, age: u8) -> Person {
        Person {
            name,
            age,
            height: 168,
            weight: 168,
        }
    }
}

pub fn 通过已有结构体实例创建新实例的简化语法() {
    #[derive(Debug)]
    struct Person {
        name: String,
        age: u8,
        height: u8,
        weight: u8,
    }

    let kmj = Person {
        name: String::from("kmj"),
        age: 31,
        height: 168,
        weight: 168,
    };

    let lm = Person {
        age: 28,
        name: String::from("lm"),
        // Person 结构体的name属性没有copy特征，因此无法使用
        ..kmj
    };

    println!("kmj {:?} , lm {:?}", kmj, lm)
}

#[allow(non_snake_case)]
pub fn 用于空值处理的枚举Option() {
    // Option 用于表达值可能存在或者不存在
    // Option 枚举包含两个成员，一个成员表示含有值：Some(T), 另一个表示没有值：None
    // Option<T> 枚举是如此有用以至于它被包含在了 prelude 之中，无需使用 Option:: 前缀就可直接使用 Some 和 None
    fn plus_one(x: Option<i32>) -> Option<i32> {
        match x {
            None => None,
            Some(i) => Some(i + 1),
        }
    }
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);
    println!("five {:?}, six:{:?} none:{:?}", five, six, none);
}

pub fn 创建非基本类型重复元素的数组初始化() {
    let array: [String; 3] = std::array::from_fn(|_i| String::from("kmj"));
    println!("array {:?}", array);
}
