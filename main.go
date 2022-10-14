package main

import (
	"sync"
	"bufio"
	"net/http"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main(){
	// fmt.Println("\n")
	fmt.Println("qsreplace Minxss'\" !!")
	fmt.Println(os.Args[0]+os.Args[1])
	fmt.Println("__________________")
	// fmt.Println("\n")
	colorReset := "\033[0m"
	colorRed := "\033[31m"
    //colorGreen := "\033[32m"


	sc := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i:= 0; i < 20; i++{

		wg.Add(1)
		go func(){
			defer wg.Done()
			for domain := range jobs {

				resp, err := http.Get(domain)
				if err != nil{
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
	      			fmt.Println(err)
	   			}
	   			sb := string(body)
	   			check_result := strings.Contains(sb , "Minxss\"'")
	   			// fmt.Println(check_result)
	   			if check_result != false {
	   				fmt.Println(string(colorRed),"\nPossible To XSS:", domain,string(colorReset))
	   			}else{
					fmt.Printf(".")
	   			}

			}
			
   		}()

	}



	for sc.Scan(){
		domain := sc.Text()
		jobs <- domain		
		

	}
	close(jobs)
	wg.Wait()

}
