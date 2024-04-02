package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type MyStu struct {
	Name  string
	Score float64
}

func sortStudentScore(className string, stuScoreMp map[string]interface{}) {
	stus := make([]MyStu, 0, len(stuScoreMp))
	for name, val := range stuScoreMp {
		if score, ok := val.(float64); ok {
			stu := MyStu{
				Name:  name,
				Score: score,
			}
			stus = append(stus, stu)
		} else {
			log.Fatal("convert to float64 failed")
		}
	}
	sort.Slice(stus, func(i, j int) bool {
		return stus[i].Score > stus[j].Score
	})
	preScore := -1
	preNum := 1
	for i := 0; i < len(stus); i++ {
		if stus[i].Score == float64(preScore) {
			fmt.Printf("%s %s %d\n", className, stus[i].Name, preNum)
		} else {
			fmt.Printf("%s %s %d\n", className, stus[i].Name, i+1)
			preNum = i + 1
			preScore = int(stus[i].Score)
		}
	}
}

func main() {
	jsonData, err := ioutil.ReadFile("score.json")
	if err != nil {
		log.Fatal(err)
	}
	scoreMap := make(map[string]interface{}) // key: className, value: map[string]int key: studentName value: score
	err = json.Unmarshal([]byte(jsonData), &scoreMap)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range scoreMap {
		if val, ok := v.(map[string]interface{}); ok {
			sortStudentScore(k, val)
		} else {
			log.Fatal("convert to map failed")
		}
	}
}
