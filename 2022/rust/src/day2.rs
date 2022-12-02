use std::collections::HashMap;

pub fn run(input: &str) {
    println!("Day 2");
    println!("Part 1: {}", part1(input));
    println!("Part 2: {}", part2(input));
}

fn get_input(input: &str) -> Vec<(&str, &str)> {
    // parse the input into a vector of <u64, u64>
    input
        .lines()
        .map(|line| {
            let mut iter = line.split_whitespace();
            let player1 = iter.next().unwrap();
            let player2 = iter.next().unwrap();
            (player1, player2)
        })
        .collect::<Vec<(&str, &str)>>()
}

fn part1(input: &str) -> u64 {
    // parse the input into a vector of <u64, u64>
    let games = get_input(input);

    // Game rules for part 1 is a map of strings to numbers
    let mut game_rules = HashMap::new();
    game_rules.insert("X".to_string(), 1);
    game_rules.insert("Y".to_string(), 2);
    game_rules.insert("Z".to_string(), 3);

    // Apply the game rules to the first item in each tuple of the games vector and sum the results
    games
        .iter()
        .map(|(player1, player2)| {
            // for each game, apply the game rules to player 2
            let player2_choice_score = game_rules.get(*player2).unwrap();
            // evaluate the game results for player 2 as win, loss, or draw
            match *player1 {
                "A" => match *player2 {
                    "X" => 3 + player2_choice_score.clone(),
                    "Y" => 6 + player2_choice_score.clone(),
                    "Z" => player2_choice_score.clone(),
                    _ => 0,
                },
                "B" => match *player2 {
                    "X" => player2_choice_score.clone(),
                    "Y" => 3 + player2_choice_score.clone(),
                    "Z" => 6 + player2_choice_score.clone(),
                    _ => 0,
                },
                "C" => match *player2 {
                    "X" => 6 + player2_choice_score.clone(),
                    "Y" => player2_choice_score.clone(),
                    "Z" => 3 + player2_choice_score.clone(),
                    _ => 0,
                },
                _ => 0,
            }
        })
        .sum()
}

fn part2(input: &str) -> u64 {
    let games = get_input(input);

    // Game rules for part 1 is a map of strings to numbers
    let mut game_rules = HashMap::new();
    game_rules.insert("A".to_string(), 1);
    game_rules.insert("B".to_string(), 2);
    game_rules.insert("C".to_string(), 3);

    // Wins loss rules
    let mut win_loss_rules = HashMap::new();
    win_loss_rules.insert("X".to_string(), 0);
    win_loss_rules.insert("Y".to_string(), 3);
    win_loss_rules.insert("Z".to_string(), 6);

    // Apply the game rules to the first item in each tuple of the games vector and sum the results
    games
        .iter()
        .map(|(player1, player2)| {
            // for each game, apply the game rules to player 2
            let player2_win_loss_score = win_loss_rules.get(*player2).unwrap();
            // evaluate the game results for player 2 as win, loss, or draw
            match *player1 {
                "A" => match *player2 {
                    "X" => game_rules.get("C").unwrap() + player2_win_loss_score.clone(),
                    "Y" => game_rules.get("A").unwrap() + player2_win_loss_score.clone(),
                    "Z" => game_rules.get("B").unwrap() + player2_win_loss_score.clone(),
                    _ => 0,
                },
                "B" => match *player2 {
                    "X" => game_rules.get("A").unwrap() + player2_win_loss_score.clone(),
                    "Y" => game_rules.get("B").unwrap() + player2_win_loss_score.clone(),
                    "Z" => game_rules.get("C").unwrap() + player2_win_loss_score.clone(),
                    _ => 0,
                },
                "C" => match *player2 {
                    "X" => game_rules.get("B").unwrap() + player2_win_loss_score.clone(),
                    "Y" => game_rules.get("C").unwrap() + player2_win_loss_score.clone(),
                    "Z" => game_rules.get("A").unwrap() + player2_win_loss_score.clone(),
                    _ => 0,
                },
                _ => 0,
            }
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "A Y
B X
C Z";
        assert_eq!(part1(input), 15);
    }

    #[test]
    fn test_part2() {
        let input = "A Y
B X
C Z";
        assert_eq!(part2(input), 12);
    }
}
