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
