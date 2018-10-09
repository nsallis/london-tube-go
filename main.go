package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "time"
  "sort"
  "encoding/json"
)

const INPUT_FILE_NAME string = "lines.txt"
const MAX_STEPS int = 5

var nodes map[string]*Node

type Node struct {
  name string
  value int
  children []*Node

}

func buildNodes() {
  b, err := ioutil.ReadFile(INPUT_FILE_NAME)
  if(err != nil) {
    fmt.Print(err)
  }
  str := string(b)
  lines := strings.Split(str, "\n")
  for _, line := range lines {
    stationNames := strings.Split(line, ", ")
    // there will only ever be 2 stations unless there is a trailing new line. Check for that
    var station1, station2 *Node
    if(len(stationNames) == 2) {
      if _, ok := nodes[stationNames[0]]; ok {
        station1 = nodes[stationNames[0]]
      }else{
        station1 = &Node{name: stationNames[0], value: 9999}
        nodes[stationNames[0]] = station1
      }
      if _, ok := nodes[stationNames[1]]; ok {
        station2 = nodes[stationNames[1]]
      }else{
        station2 = &Node{name: stationNames[1], value: 9999}
        nodes[stationNames[1]] = station2
      }

      station1.children = append(station1.children, station2)
      station2.children = append(station2.children, station1)
    }
  }
}

func doNode(node *Node, value int) {
  if(value > MAX_STEPS) {
    return
  }
  if(node.value > value || node.value == 0) {
    node.value = value
    for _, child := range node.children {
      doNode(child, value + 1)
    }
  }
}

func main() {
  t1 := time.Now()
  nodes = make(map[string]*Node)
  buildNodes()
  root := nodes["East Ham"]
  doNode(root, 0)
  stations := []string{}
  for _, node := range nodes {
    if(node.value == MAX_STEPS){
      stations = append(stations, node.name)
    }
  }

  sort.Strings(stations)
  stationsJson, _ := json.Marshal(stations)
  fmt.Println(string(stationsJson))

  elapsed := time.Since(t1)
  fmt.Printf("finding routes took: %f milliseconds", float64(elapsed)/1000000)
}
