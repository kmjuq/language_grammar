pub fn 变量的绑定() {
    // 变量的绑定 将字面量 “hello world” 给 hello 变量
    let hello = "hello world";
    println!("{}", hello);
}

pub fn 变量的可变和不可变声明() {
    // 变量的可变、不可变声明
    let mut kmj_mut = String::from("kmj");
    let kmj = String::from("kmj");
    println!("{} {}", kmj_mut, kmj);
    kmj_mut = String::from("value");
    // error 不可变类型无法再次赋值
    // kmj = String::from("value");
    println!("{} {}", kmj_mut, kmj);
}

pub fn 使用下划线开头忽略未使用的变量() {
    let kmj = String::from("kmj");
    let _lm = String::from("lm");
    println!("{}", kmj);
}

pub fn 变量遮蔽() {
    let x = 5;
    // 在main函数的作用域内对之前的x进行遮蔽
    let x = x + 1;
    {
        // 在当前的花括号作用域内，对之前的x进行遮蔽
        let x = x * 2;
        println!("The value of x in the inner scope is: {}", x);
    }
    println!("The value of x is: {}", x);
}

/// 常量必须指定数据类型
pub fn 常量的绑定() {
    const KMJ_NAME: &str = "kemengjian";
    println!("{}", KMJ_NAME)
}
