(ns clojure-aoc2020.day2
  "Advent of Code 2020: Day 2 Challenge"
  (:require [clojure.string :as string]))

(defstruct passwd :positions :character :password)

(defn structure-data
  [line]
  (let [fields (string/split line #"\s+")]
    (struct-map passwd :positions (map read-string (string/split (first fields) #"-"))
                :character (read-string (string/replace (second fields) ":" ""))
                :password (nth fields 2))))

(defn get-data
  [f]
  (-> f
      slurp
      string/trim
      (#(string/split-lines %))
      (#(map structure-data %))))

(defn -validate-part1
  [pass-line]
  (let [c-min (first (:positions pass-line))
        c-max (second (:positions pass-line))
        pass (:password pass-line)
        c (:character pass-line)
        c-count (count (apply str (filter #(= (str c) (str %)) pass)))]
    (and (>= c-count c-min) (<= c-count c-max))))

(defn -validate-part2
  [pass-line]
  (let [positions (map dec (:positions pass-line))
        pass (:password pass-line)
        c (:character pass-line)
        index-chars (map #(nth pass %) positions)]
    ;; only 1 character index can have the correct character
    (= 1 (count (filter #(= (str c) (str %)) index-chars)))))

(defn main
  [& args]
  (let [;; data (get-data "data/day2.test.dat")
        data (get-data "data/day2.dat")
        part1 (count (filter -validate-part1 data))
        part2 (count (filter -validate-part2 data))]
    (str "Part 1: " part1 "\nPart 2: " part2)))
