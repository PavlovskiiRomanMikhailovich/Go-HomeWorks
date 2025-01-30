package uniq

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
)

type Uniq struct {
	data          []string
	counter       map[int]int
	result        []string
	num_filds     int
	num_chars     int
	register_flag bool
}

func (u *Uniq) ReadData(path string) {
	f, err := os.Open(path)
	if err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			u.data = append(u.data, sc.Text())
		}
	}
	defer f.Close()
	return
}

func (u *Uniq) WriteData(path string) {
	f, err := os.Create(path)
	if err == nil {
		for i := 0; i < len(u.result); i++ {
			f.WriteString(u.result[i] + "\n")
		}
	}
	defer f.Close()

}
func (u *Uniq) DelDuplicate() {
	u.Counter()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] != 0 {
			u.result = append(u.result, u.data[i])
		}
	}
}

func (u *Uniq) Count_Number() {
	u.Counter()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] != 0 {
			u.result = append(u.result, strconv.Itoa(u.counter[i])+" "+u.data[i])
		}
	}

}

func (u *Uniq) Counter() {
	u.counter = make(map[int]int)
	for i := 0; i < len(u.data); i++ {
		i2 := i
		for i2 = i; i2 < len(u.data); i2++ {
			if u.register_flag {
				if strings.ToLower(u.SliceString(u.data[i2])) == strings.ToLower(u.SliceString(u.data[i])) {
					u.counter[i] += 1
				} else {
					break
				}
			} else {
				if u.SliceString(u.data[i2]) == u.SliceString(u.data[i]) {
					u.counter[i] += 1
				} else {
					break
				}
			}
		}
		i = i2 - 1

	}

	return
}

func (u *Uniq) SliceString(str string) string {
	split_string := strings.Split(str, " ")
	total := ""
	for index := 0; index < len(split_string); index++ {
		if index >= u.num_filds {
			total += " " + split_string[index]

		}

	}
	current := strings.Split(total, "")
	total_2 := ""
	for index := 0; index < len(current); index++ {
		if index > u.num_chars {
			total_2 += current[index]

		}

	}
	return total_2
}

func (u *Uniq) OnlyUnique() {
	u.Counter()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] == 1 {
			u.result = append(u.result, u.data[i])
		}

	}
}

func (u *Uniq) NotUnique() {
	u.Counter()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] > 1 {
			u.result = append(u.result, u.data[i])
		}
	}
}

func main() {
	// const
	FlagC := flag.Bool("c", false, "c")
	FlagD := flag.Bool("d", false, "d")
	FlagU := flag.Bool("u", false, "u")
	FlagI := flag.Bool("i", false, "i")
	//
	var slice_fields int
	var slice_chars int
	flag.IntVar(&slice_fields, "f", 0, "f")
	flag.IntVar(&slice_chars, "s", 0, "s")
	flag.Parse()
	// Create constructor
	t := Uniq{num_filds: slice_fields, num_chars: slice_chars, register_flag: *FlagI}
	input_file := os.Args[len(os.Args)-2]
	output_file := os.Args[len(os.Args)-1]
	t.ReadData(input_file)
	if *FlagC {
		t.Count_Number()
		t.WriteData(output_file)
	}
	if *FlagD {
		t.NotUnique()
		t.WriteData(output_file)
	}
	if *FlagU {
		t.OnlyUnique()
		t.WriteData(output_file)
	}
	if !*FlagC && !*FlagD && !*FlagU {
		t.DelDuplicate()
		t.WriteData(output_file)
	}

}
