(ns clojure-aoc2020.day3
  (:require [clojure.string :as string]
            [clojure.java.io :as io]))

(defn get-data
  [f]
  (->> f
       slurp
       string/trim
       string/split-lines
       (map seq)))

(defn map-metadata
  [data]
  (let [length (count data)
        width (first (map #(count %) data))]
    {:length length :width width}))

(defn numeric-map
  [data]
  (mapv #(mapv {\# 1 \. 0} %) data))

(defn get-coordinate
  [[x y] map-metadata data]
  (let [width (:width map-metadata)]
    (get-in data [y (mod x width)])))

(defn get-path
  [[dx dy] map-meta]
  (let [l (:length map-meta)
        w (:width map-meta)]
    (->> (range (count (range (/ l dy))))
         (map (fn
                [i]
                (let [x (mod (* dx i) w)
                      y (* dy i)]
                  [x y]))))))

(defn solve
  [[dx dy] metadata number-map]
  (->> (get-path [dx dy] metadata)
       (map #(get-coordinate % metadata number-map))
       (reduce +)))

(comment
  (def data (get-data "data/day3.test.dat"))
  (def meta (map-metadata data))
  (def numeric-data (numeric-map data))
  (get-coordinate [0 1] meta numeric-data)
  (def part1-path (get-path [3 2] meta))
  (map #(get-coordinate % meta numeric-data) part1-path)
  (->> (get-path [3 1] meta)
       (map #(get-coordinate % meta numeric-data))
       (reduce +))
  (map #(mod (* 1 %) 11) (range (/ 20 2))))

(defn main
  [& args]
  (let [data (get-data "data/day3.dat")
        meta (map-metadata data)
        num-map (numeric-map data)
        part1 (solve [3 1] meta num-map)
        part2 (->> [[1 1] [3 1] [5 1] [7 1] [1 2]]
                   (map #(solve % meta num-map))
                   (reduce *))]
    (str "Part 1: " part1 "\nPart 2: " part2)))
