# Accuknox--Assignments


api.go consists of POST, GET AND DELETE OPERATION FOR THE FOLLOWING QUESTION 

Problem Statement 1
1. Create an REST API server in Golang which implements the endpoints mentioned
below.
2. Containize the application using Docker and host the container image in DockerHub
(optional, preferred)
3. Push the code to GitHub and submit the URL of the git repository.

Note :- I have not processed with Docker because it was optional, if You wish to do say yes to do it. please comment below so i can process with it.



2ndquestion.md is the answer for following question 

Problem Statement 2
Explain the following code snippet. Explain what the code is attempting to do? You can explain
by: Giving use-cases of what this construct/pattern could be used for?

package main
import &quot;fmt&quot;
func main() {
cnp := make(chan func(), 10)
for i := 0; i &lt; 4; i++ {
go func() {
for f := range cnp {
f()
}
}()
}
cnp &lt;- func() {
fmt.Println(&quot;HERE1&quot;)
}
fmt.Println(&quot;Hello&quot;)
}
