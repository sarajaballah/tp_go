// Begin gaining familiarity with Go-style concurrency primitives. 
// You will need to refer to the Go documentation.
package main

import (
	"fmt"
	"os"
    "strconv"
    "sync"
)

func main() {
	fmt.Println("========== Problem 1 : File processing ==========")
	Sum("numbers.txt", "sum.txt")

	
	fmt.Println("========== Problem 2 :  Concurrent map access==========")
	d := EnseirbDirectory{
		directory: make(map[int]string),
	}
	total := 1000
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			switch i % 3 {
			case 0:
				d.Add(1, "Aurore LI")
			case 1:
				d.Add(2, "Catie")
			case 2:
				d.Remove(2)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("d.directory = %v\n", d.directory)
}

// Problem 1: File processing with interfaces
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// You should use the interfaces for io.Reader and
// io.Writer to do this. 

// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.
// You should expect your input to end with a newline, and the output should
// have a newline after the result
func Sum(input, output string) {
	// TODO
	file, err := os.Open(input)
    if err != nil {
        fmt.Print(err)
    }
    lines := make([]byte, 1024)
    nbr, err := file.Read(lines)
    if err != nil {
        fmt.Print(err)
    }
    res := 0
    var str string
    for i := 0; i < nbr; i++ {
		if string(lines[i]) != "\n" {
        	if string(lines[i]) != "\r" {
            	str += string(lines[i])
        	} else {
            	r, _ := strconv.Atoi(str)
				res += r
            	str = ""
        	}
		}
	}
	fmt.Println("The result is : " + strconv.Itoa(res))
    outputF, _ := os.Create(output)
    outputF.Write([]byte(strconv.Itoa(res)))
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// You will build a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// EnseirbDirectory is a mapping from ID number to name (12345678 -> "Aurore LI").
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex
type EnseirbDirectory struct {
	// TODO
	mutex sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Enseirb Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *EnseirbDirectory) Add(id int, name string) {
	// TODO
	d.mutex.Lock()
    d.directory[id] = name
    d.mutex.Unlock()
}

// Get fetches a student from the Enseirb Directory by their ID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *EnseirbDirectory) Get(id int) string {
	// TODO
	d.mutex.RLock()
    data := d.directory[id]
    d.mutex.RUnlock()
    return data
}

// Remove a student to the Enseirb Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *EnseirbDirectory) Remove(id int) {
	// TODO
	d.mutex.Lock()
    delete(d.directory, id)
    d.mutex.Unlock()
}