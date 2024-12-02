import gleam/io
import gleam/hackney
import gleam/http.{Get}
import gleam/http/request
import envoy
import simplifile

pub fn pull_day(day: String, output_dir: String) {
  let session_token = case envoy.get("AOC_SESSION_TOKEN") {
    Ok(token) -> token
    Error(_) -> {
      panic as "AOC_SESSION_TOKEN not set"
    }
  }
  let output_file = output_dir <> "/day" <> day <> ".txt"
  io.println("Pulling day " <> day <> " to " <> output_file)

  // Get data from aoc website
  let req =
    request.new()
    |> request.set_method(Get)
    |> request.set_host("adventofcode.com")
    |> request.set_path("/2024/day/" <> day <> "/input")
    |> request.set_header("Cookie", "session=" <> session_token)
    |> request.set_header(
      "User-Agent",
      "Jordan Garrison/aoc2024 gleaming my inputs",
    )

  let res = case hackney.send(req) {
    Ok(res) -> res
    Error(_) -> {
      panic as "Failed to fetch input"
    }
  }

  let file = simplifile.write(to: output_file, contents: res.body)
  case file {
    Ok(_) -> io.println("Pulled day " <> day)
    Error(_) -> io.println("Failed to pull day " <> day)
  }
}
