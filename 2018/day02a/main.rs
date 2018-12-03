use std::io;
use std::io::BufRead;

fn main() {
	let stdin = io::stdin();
	let input = stdin.lock().lines().filter_map(|line| line.ok());

	let mut twos = 0;
	let mut threes = 0;
	for l in input {

		let mut counts = [0; 26];
		for c in l.chars() {
			if c < 'a' || c > 'z' {
				println!("ERROR: not a letter");
				continue;
			}
			counts[(c as usize) - ('a' as usize)] += 1;
		}

		let mut has2 = false;
		let mut has3 = false;
		for c in counts.iter() {
			if *c == 2 {
				has2 = true;
			}
			if *c == 3 {
				has3 = true;
			}
		}
		if has2 {
			twos += 1;
		}
		if has3 {
			threes += 1;
		}

	}

	println!("Answer: {}", twos * threes);
}
