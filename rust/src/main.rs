use std::fs::read_to_string;
use itermore::IterMore;

fn main() {
    let i: Vec<u32> = get_input("src/pkg/adventofcode/one/input.txt");
    // println!("{:?}", i);
    let pt1: usize = part_one(&i);
    println!("Part One: {:?}", pt1);
    
    
    let pt2: usize = part_two(&i);
    println!("Part Two: {:?}", pt2);
}

fn get_input(f: &str) -> Vec<u32> {
    let input = read_to_string(&f)
        .expect("File not found");

    let res: Vec<u32> = input
        .lines()
        .filter_map(|dep| dep.parse().ok())
        .collect();
    
        res
}

fn part_one(i: &[u32]) -> usize {
    let s: Vec<u32> = i.to_vec();
    get_increase_count(&s)
}

fn part_two(i: &[u32]) -> usize {
    let s: Vec<u32> = get_window_sums(i);

    get_increase_count(&s)
}

fn get_window_sums(i: &[u32]) -> Vec<u32> {
    i.windows(3)
    .map(|w| w.iter().sum())
    .collect()
}

fn get_increase_count(i: &Vec<u32>) -> usize {
    i.windows(2).filter(|w| w[0] < w[1]).count()
}