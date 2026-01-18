
#[test]
fn log2(){
    Solution::reordered_power_of2(1);
}
struct Solution;

impl Solution {
    pub fn reordered_power_of2(n: i32) -> bool {
        let (min,max): (i32, i32) = Self::get_range_from_number(n);
        let binary_num_str_sort = (min.ilog2()+1 ..= max.ilog2()).map(|x| {
        let mut chars = 2i32.pow(x).to_string().chars().collect::<Vec<_>>();
            chars.sort();
            chars.into_iter().collect::<String>()
        }).collect::<Vec<_>>();
        let mut num_char_vec_sort = n.to_string().as_str().chars().collect::<Vec<_>>();
        num_char_vec_sort.sort();
        let num_str_sort = num_char_vec_sort.into_iter().collect::<String>();
        binary_num_str_sort.contains(&num_str_sort)
    }

    fn count_digits_i32(n: i32) -> u32 {
        match n {
            0..=9 => 1u32,
            10..=99 => 2u32,
            100..=999 => 3u32,
            1000..=9999 => 4u32,
            10000..=99999 => 5u32,
            100000..=999999 => 6u32,
            1000000..=9999999 => 7u32,
            10000000..=99999999 => 8u32,
            100000000..=999999999 => 9u32,
            _ => 10,  // i32 最大值为 2147483647（10 位）
        }
    }

        /// 根据位数获取最小值和最大值
    fn get_range_by_digits(digits: u32) -> (i32, i32) {
        if digits == 0 {
            panic!("位数不能为0");
        }
        // 计算 10^(digits-1) 作为最小值
        let min = 10i32.pow(digits - 1);
        // 计算 10^digits - 1 作为最大值
        let max = 10i32.pow(digits) - 1;
        (min as i32, max as i32)
    }

    /// 根据整数获取其位数对应的最大最小值
    fn get_range_from_number(n: i32) -> (i32, i32) {
        let digits = Self::count_digits_i32(n);
        Self::get_range_by_digits(digits)
    }
}