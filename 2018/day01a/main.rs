use std::io;

fn main() {
	let input = read();
	let mut answer = 0;
	for v in input.iter() {
		answer += v
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
