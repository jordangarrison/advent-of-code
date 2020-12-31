(ns clojure-aoc2020.day1
  "Advent of Code 2020: Day 1 Challenge"
  (:require [clojure.string :as string]
            [clojure.math.combinatorics :as combo]))

(defn get-data
  [f]
  (-> f
      slurp
      string/trim
      (#(string/split % #"\n"))
      (#(map read-string %))))

(defn generate-combinations
  [data number]
  (combo/combinations data number))

(defn filter-combination
  [data target]
  (filter #(= target (apply + %)) data))

(defn get-multiplied-combination
  [data n total]
  (-> data
      (#(generate-combinations % n))
      (#(filter-combination % total))
      first
      (#(reduce * %))))

(defn main
  [& args]
  (let [;; data (get-data "data/day1.test.dat")
        data (get-data "data/day1.dat")
        part1 (get-multiplied-combination data 2 2020)
        part2 (get-multiplied-combination data 3 2020)]
    (str "Part 1: " part1 ", Part 2: " part2)))
