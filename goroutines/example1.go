package main
import (
    "fmt"
)

func worker(id int, jobs <- chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Воркер %d начал задачу: %d\n", id, job)
        result := job * 2
        results <- result
        fmt.Printf("Воркер %d закончил задачу %d -> %d\n", id, job, result)
    }
}

func main() {
    const numJobs int = 5
    const numWorkers int = 2
    
    test := make(chan<- int, numJobs)  // канал в который можно только отправлять
    test1 := make(<-chan int, numJobs)  // канал из которого можно только получать

    jobs := make(chan int, numJobs)
    results := make(chan int, numWorkers)
    
    for w := 1; w < numWorkers; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j < numJobs; j++ {
        jobs <- j
    }

    close(jobs)

    for r := 1; r <= numJobs; r++ {
        fmt.Println(<- results)
    }
}
