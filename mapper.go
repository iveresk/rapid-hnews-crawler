package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func mapMerge(ms, ams map[string]string) map[string]string {
	for k, v := range ams {
		ms[k] = v
	}
	return ms
}

func blogEqual(ms, ams map[string]string) (map[string]string, bool) {
	if len(ms) != len(ams) {
		return nil, false
	}

	res := make(map[string]string)
	isequal := true

	for k, v := range ms {
		av := ams[k]
		if av != v {
			isequal = false
			res[k] = v
		}
	}
	return res, isequal
}

func mapReader(fn string) (map[string]string, error) {
	file, ferr := ioutil.ReadFile(fn)
	if ferr != nil {
		return nil, ferr
	}
	file_res := make(map[string]string)
	jerr := json.Unmarshal(file, &file_res)
	if jerr != nil {
		return nil, jerr
	}
	return file_res, nil
}

func mapWriter(ms map[string]string, fn string) error {
	jsonout, jerr := json.Marshal(ms)
	if jerr != nil {
		return jerr
	}
	err := ioutil.WriteFile(filename, jsonout, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// TODO mapreader and mapwriter for the MongoDB
