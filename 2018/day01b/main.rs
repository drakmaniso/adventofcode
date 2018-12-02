use std::collections::HashMap;
use std::io;

fn main() {
	let input = read();
	let mut answer = 0;
	let mut previous = HashMap::new();
	previous.insert(0, true);
	'l: loop {
		for v in input.iter() {
			answer += v;
			if previous.contains_key(&answer) {
				break 'l;
			}
			previous.insert(answer, true);
		}
	}
	println!("Answer: {}", answer);
}

fn read() -> Vec<i32> {
	let mut input = Vec::new();
	loop {
		let mut v = String::new();
		io::stdin().read_line(&mut v).expect("failed to read input");
		if v.len() <= 1 {
			return input;
		}
		let v: i32 = v[..v.len() - 1].parse().expect("failed to parse input");
		input.push(v);
	}
}
