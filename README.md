# Golang Program to list 100 most common three word sequences in the text 

**Objective:**  
To write a Go program to provide a list of the 100 most common three word sequences in the text,
along with a count of how many times each occurred in the text.

**Requirements:**  

Before proceeding, please ensure that you have Docker/Go  installed and running on your system. 


## Steps to Run with Docker 
1. **Step 1: Run Command** Run the command ```make start-the-process``` on your terminal at the root of the project. 
    - This downloads all the dependency, builds docker images for application and startup the container

2. **Step 2: Check the terminal or docker logs for output
3. **Step 3: Termination and Cleanup** Run the command ```make stop-the-process``` on your terminal at the root of the project.
    - This removes the docker image, stops the docker containers which was create during *Step 1* and removes them from system 

## Steps to run in your system
1. **Step1: Make sure you have golang Installed (1.20)
2. **Step2: ``` go run server/main.go file2.txt``` to run the code with sample file2.txt
3. **Step3: If you want to provide input from STDIN```go build main.go``` and then ```cat file2.txt | ./main -```

The sample files used are file1.txt and file2.txt. And the sample output is below for file2.txt.
"file1.txt" has also been tested and it has more lines.
```
Top 100 three-word sequences:
9 - of natural selection
6 - conditions of life
4 - have been formed
4 - unity of type
4 - we have seen
3 - conditions of existence
3 - have seen that
3 - the conditions of
2 - a continuous area
2 - at any one
2 - at the same
2 - been formed by
2 - by means of
2 - by natural selection
2 - by unity of
2 - could not have
2 - follows from the
2 - formed by natural
2 - habits of life
2 - in all cases
2 - in concluding that
2 - in many cases
2 - in this chapter
2 - intermediate variety will
2 - means of natural
2 - must often have
2 - natural selection acts
2 - natural selection in
2 - natural selection is
2 - naturally follows from
2 - new conditions of
2 - not have been
2 - of the same
2 - of type is
2 - of useful structures
2 - partly because the
2 - principle of natural
2 - process of natural
2 - species under new
2 - that a part
2 - the inhabitants of
2 - the larger country
2 - the principle of
2 - the same class
2 - the same time
2 - the standard of
2 - the theory of
2 - theory of natural
2 - through natural selection
2 - under new conditions
2 - we know of
2 - we should be
2 - will have been
1 - a bat for
1 - a few forms
1 - a great advantage
1 - a long series
1 - a multitude of
1 - a part formerly
1 - a part or
1 - a species that
1 - a species under
1 - a swim bladder
1 - absolute perfection be
1 - absolute perfection nor
1 - accordance with the
1 - account for the
1 - accumulated by means
1 - acknowledged that all
1 - acquired by means
1 - acquirement of any
1 - acquisition through natural
1 - action of external
1 - acts by either
1 - acts of creation
1 - acts only on
1 - acts through the
1 - adaptations being aided
1 - adaptations that of
1 - adapted them during
1 - adapting the varying
1 - advantage of and
1 - advantage over the
1 - affected by the
1 - again naturally follows
1 - against the theory
1 - agreement in structure
1 - aided by the
1 - aided in many
1 - air breathing lung
1 - air we have
1 - all cases at
1 - all cases to
1 - all organic beings
1 - all those of
1 - allied species now
1 - almost always be
1 - alone of the
1 - also believe that
1 - although the belief
```