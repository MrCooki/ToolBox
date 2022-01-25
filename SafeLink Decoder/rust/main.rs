use urlencoding::decode;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let args: &str = &*args[1];

    let split_1: Vec<&str> = args.split("outlook.com/?url=").collect();
    let split_2: Vec<&str> = split_1[1].split("&data").collect();
    let url = decode(split_2[0]).expect("UTF-8");

    println!("Decoded URL: {}",url)



}
