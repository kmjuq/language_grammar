use std::{
    fmt::{self, Debug, Display},
};

#[allow(dead_code)]
pub fn 泛型语法() {
    // 结构体中的泛型语法
    #[derive(Debug)]
    #[allow(dead_code)]
    struct Point<T> {
        x: T,
        y: T,
    }
    // 方法中的泛型语法
    impl<T> Point<T> {
        fn x(&self) -> &T {
            &self.x
        }
    }
}

/// trait 关键字是用来定义特征的，特征类似于java的接口，用来抽象行为
/// 比如 你无法在当前作用域中，为 String 类型实现 Display 特征，因为它们俩都定义在标准库中。
/// 特征的方法可以有默认实现，java8的接口的默认实现类似
#[allow(dead_code)]
pub fn 特征的基本使用() {
    // 特征定义的语法
    pub trait Summary {
        fn summarize(&self) -> String;
        // 默认实现
        fn summarize_default(&self) -> String {
            String::from("(Read more...)")
        }
    }

    pub struct Post {
        title: String,
        author: String,
        content: String,
    }
    pub struct Weibo {
        username: String,
        content: String,
    }

    impl Summary for Post {
        fn summarize(&self) -> String {
            format!("文章 {}, 作者是{}", self.title, self.author)
        }
        fn summarize_default(&self) -> String {
            format!(
                "summarize_default 文章 {}, 作者是{}",
                self.title, self.author
            )
        }
    }

    impl Summary for Weibo {
        fn summarize(&self) -> String {
            format!("{} 发表了微博 {}", self.username, self.content)
        }
    }

    let post = Post {
        title: "Rust语言简介".to_string(),
        author: "Sunface".to_string(),
        content: "Rust棒极了!".to_string(),
    };
    let weibo = Weibo {
        username: "sunface".to_string(),
        content: "好像微博没Tweet好用".to_string(),
    };

    println!("{}", post.summarize());
    println!("{}", weibo.summarize());
    println!("{}", weibo.summarize_default());
    println!("{}", weibo.summarize_default());

    // 将特征当作函数参数
    pub fn notify(item: &impl Summary) {
        println!("Breaking news! {}", item.summarize());
    }
}

#[allow(dead_code)]
pub fn 特征约束() {
    pub trait Summary {
        fn summarize(&self) -> String;
    }

    // 普通约束
    pub trait NormalBound {
        fn notify1(item: &impl Summary);
        // 前者是后者的语法糖
        fn notify2<T: Summary>(item: &T);

        fn notify3(item1: &impl Summary, item2: &impl Summary);
        // 前者可以是两个不同的Summary实例，后者必须是相同的Summary实例
        fn notify4<T: Summary>(item1: &T, item2: &T);
    }

    // 特征约束
    pub trait MultiBound {
        fn notify5(item: &(impl Summary + Display));
        // 前者是语法糖
        fn notify6<T: Summary + Display>(item: &T);
    }

    // where 约束
    pub trait WhereBound {
        fn some_function1<T: Display + Clone, U: Clone + Debug>(t: &T, u: &U) -> i32;
        // 当特征约束变多时，使用where约束更容易理解。
        fn some_function2<T, U>(t: &T, u: &U) -> i32
        where
            T: Display + Clone,
            U: Clone + Debug;
    }

    // 特征定义中的特征约束
    // 说明特征 OutlinePrint，依赖特征 Display，需要同时实现 Display
    trait OutlinePrint: Display {
        fn outline_print(&self) {
            let output = self.to_string();
            let len = output.len();
            println!("{}", "*".repeat(len + 4));
            println!("*{}*", " ".repeat(len + 2));
            println!("* {} *", output);
            println!("*{}*", " ".repeat(len + 2));
            println!("{}", "*".repeat(len + 4));
        }
    }
    struct Point {
        x: i32,
        y: i32,
    }

    impl OutlinePrint for Point {}

    // 以下特征实现被注释会报错。
    impl fmt::Display for Point {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            write!(f, "({}, {})", self.x, self.y)
        }
    }
}

/// 只能返回单一的具体类型，只是用来简化返回类型，比如说返回是迭代器或者闭包
#[allow(dead_code)]
pub fn 函数可以返回impl_trait() {
    pub fn display() -> impl ToString {
        62
    }
}

/// 特征的方法不能返回Self、方法没有任何泛型参数
pub fn 特征对象() {
    pub trait Draw {
        fn draw(&self);
    }

    pub struct Button;

    impl Draw for Button {
        fn draw(&self) {
            println!("button box draw")
        }
    }

    struct SelectBox;

    impl Draw for SelectBox {
        fn draw(&self) {
            println!("select box draw")
        }
    }

    pub struct Screen {
        pub components: Vec<Box<dyn Draw>>,
    }

    impl Screen {
        fn run(&self) {
            for ele in self.components.iter() {
                ele.draw();
            }
        }
    }

    let screen = Screen {
        components: vec![Box::new(SelectBox), Box::new(Button)],
    };

    screen.run();
}

#[allow(dead_code)]
pub fn 关联类型() {
    pub trait CacheableItem: Clone + Default + fmt::Debug {
        type Address: AsRef<[u8]> + Clone + fmt::Debug + Eq;
        fn is_null(&self) -> bool;
    }
}

#[allow(dead_code)]
pub fn 默认泛型参数() {
    trait Add<RHS = Self> {
        type Output;

        fn add(self, rhs: RHS) -> Self::Output;
    }
}

#[allow(dead_code)]
pub fn 调用不同特征的同名方法() {
    trait Pilot {
        fn fly(&self);
    }

    trait Wizard {
        fn fly(&self);
    }

    struct Human;

    impl Pilot for Human {
        fn fly(&self) {
            println!("This is your captain speaking.");
        }
    }

    impl Wizard for Human {
        fn fly(&self) {
            println!("Up!");
        }
    }

    impl Human {
        fn fly(&self) {
            println!("*waving arms furiously*");
        }
    }

    let person = Human;
    Pilot::fly(&person); // 调用Pilot特征上的方法
    Wizard::fly(&person); // 调用Wizard特征上的方法
    person.fly();

    trait Animal {
        fn baby_name() -> String;
    }

    struct Dog;

    impl Dog {
        fn baby_name() -> String {
            String::from("Spot")
        }
    }

    impl Animal for Dog {
        fn baby_name() -> String {
            String::from("puppy")
        }
    }

    println!("A baby dog is called a {}", Dog::baby_name());
    println!("A baby dog is called a {}", <Dog as Animal>::baby_name());
}

/// 特征有个孤儿原则，如果你想要为类型 A 实现特征 T，那么 A 或者 T 至少有一个是在当前作用域中定义的！
#[allow(dead_code)]
pub fn 绕过孤儿原则() {
    // 绕过孤儿原则，给 Vec<String> 类型添加 Display 特征，通过创建 Wrapper 包装结构体
    struct Wrapper(Vec<String>);

    impl fmt::Display for Wrapper {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            write!(f, "[{}]", self.0.join(", "))
        }
    }

    fn main() {
        let w = Wrapper(vec![String::from("hello"), String::from("world")]);
        println!("w = {}", w);
    }
}
