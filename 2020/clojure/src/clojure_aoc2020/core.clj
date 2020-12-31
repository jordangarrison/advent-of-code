(ns clojure-aoc2020.core
  (:require [clojure-aoc2020.day1 :as day1]
            [clojure-aoc2020.day2 :as day2])
  (:gen-class))

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println
   (time (case (Integer/parseInt (first args))
           1 (day1/main)
           2 (day2/main)
           "Invalid day"))))
