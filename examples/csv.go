/*
Simple example of loading a CSV and filtering.
Dataset is from https://github.com/vincentarelbundock/Rdatasets
*/

package main

import (
	"fmt"
	df "github.com/kniren/gota/data-frame"
	"io/ioutil"
)

const (
	path      = "./data/LakeHuron.csv"
	threshold = 580.58
)

func main() {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dataframe := df.ReadCSV(string(data), "float")
	err = dataframe.Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d rows and %d columns \n", dataframe.Nrow(), dataframe.Ncol())
	dataframe = dataframe.Filter(df.F{Colname: "LakeHuron", Comparator: ">=", Comparando: threshold})
	for _, record := range dataframe.Col("LakeHuron").Records() {
		fmt.Printf("%s >= %f\n", record, threshold)
	}
}
