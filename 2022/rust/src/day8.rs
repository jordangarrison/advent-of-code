pub fn run(input: &str) {
    println!("Day 8");
    let (part1, part2) = problem(input);
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
}

fn problem(input: &str) -> (u64, u64) {
    // read input into a matrix
    let matrix = input
        .lines()
        .map(|line| {
            line.chars()
                .map(|c| c.to_digit(10).unwrap() as u64)
                .collect::<Vec<u64>>()
        })
        .collect::<Vec<Vec<u64>>>();

    // the matrix represents a grid of trees
    // trees are visible if the surounding trees are smaller than them
    // all trees on the edges are visible because they are on the edges
    // determine the number of visible trees
    let mut visible_trees = 0;
    let mut visible_grid = vec![vec![false; matrix[0].len()]; matrix.len()];
    let mut max_view_distance: u64 = 0;

    // determine for any given position in the matrix if that is the maximum value of
    // the row before it
    // or the row behind it
    // or the colume above it
    // or the colume below it
    // if it is then it is visible
    for i in 0..matrix.len() {
        for j in 0..matrix[i].len() {
            // get 4 vecs of the trees above, below, left, and right
            let mut above: Vec<u64> = vec![];
            let mut below: Vec<u64> = vec![];
            let mut left: Vec<u64> = vec![];
            let mut right: Vec<u64> = vec![];
            for k in 0..matrix[i].len() {
                if j > k {
                    left.push(matrix[i][k]);
                } else if j < k {
                    right.push(matrix[i][k]);
                }
            }
            for k in 0..matrix.len() {
                if i > k {
                    above.push(matrix[k][j]);
                } else if i < k {
                    below.push(matrix[k][j]);
                }
            }
            // check if the current tree is the max of any of the 4 vecs
            if matrix[i][j] > *above.iter().max().unwrap_or(&0)
                || matrix[i][j] > *below.iter().max().unwrap_or(&0)
                || matrix[i][j] > *left.iter().max().unwrap_or(&0)
                || matrix[i][j] > *right.iter().max().unwrap_or(&0)
                || i == 0
                || i == matrix.len() - 1
                || j == 0
                || j == matrix[i].len() - 1
            {
                visible_grid[i][j] = true;
                visible_trees += 1;
            }
            let view_distance = view_distance(matrix[i][j], above, below, left, right);
            if view_distance > max_view_distance
                && i != 0
                && i != matrix.len() - 1
                && j != 0
                && j != matrix[i].len() - 1
            {
                max_view_distance = view_distance;
            }
        }
    }
    (visible_trees, max_view_distance)
}

fn view_distance(
    height: u64,
    above: Vec<u64>,
    below: Vec<u64>,
    left: Vec<u64>,
    right: Vec<u64>,
) -> u64 {
    let mut above_distance = 1;
    let mut below_distance = 1;
    let mut left_distance = 1;
    let mut right_distance = 1;
    let above_rev = above.iter().rev().collect::<Vec<&u64>>();
    let left_rev = left.iter().rev().collect::<Vec<&u64>>();
    for i in 0..above_rev.len() {
        if *above_rev[i] < height && i != above_rev.len() - 1 {
            above_distance += 1;
        } else {
            break;
        }
    }
    for i in 0..below.len() {
        if below[i] < height && i != below.len() - 1 {
            below_distance += 1;
        } else {
            break;
        }
    }
    for i in 0..left_rev.len() {
        if *left_rev[i] < height && i != left_rev.len() - 1 {
            left_distance += 1;
        } else {
            break;
        }
    }
    for i in 0..right.len() {
        if right[i] < height && i != right.len() - 1 {
            right_distance += 1;
        } else {
            break;
        }
    }
    above_distance * below_distance * left_distance * right_distance
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_problem() {
        let input = "30373
25512
65332
33549
35390";
        assert_eq!(problem(input), (21, 8));
    }
}
