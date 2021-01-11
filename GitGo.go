package main

import (
	"os"
	"os/exec"
	"log"
	"time"
)

func main() {
	logTime("Start");

	runDirectory := os.Args[1]

	directories := getAllDirectories(runDirectory)

	writeToFile("log", runDirectory + "\n")

	for i := 0; i < len(directories); i++ {
		directory := runDirectory + "/" + directories[i]

		writeToFile("log", "\t./" + directories[i] + ": ")
		runGit(directory, "fetch")
		runGit(directory, "pull")
	}

	logTime("End");
	writeToFile("log", "\n")
}

func getAllDirectories(fileLocation string) []string {
	file, err := os.Open(fileLocation)
	
	if err != nil {
		fatalLog("getAllDirectories", err.Error())
	}

	directories, err := file.Readdirnames(0)

	if err != nil {
		fatalLog("getAllDirectories", err.Error())
	}

	return directories
}

func runGit(directory, command string) {
	output, err := exec.Command("git", "-C", directory, command).Output()

	if err != nil {
		fatalLog("runGit", err.Error())
	}

	writeToFile("log", string(output))
}

func fatalLog(location, information string) {
	fullLog := location + " : " + information

	writeToFile("fatal", fullLog + "\n")
	log.Fatal(fullLog);
}

func logTime(information string) {
	currentTime := time.Now()
	writeToFile("log", information + " " + currentTime.String() + "\n")
}

func writeToFile(fileName, text string) {
	file, err := os.OpenFile(fileName,  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 622)

    if err != nil {
		fatalLog("writeToFile", err.Error())
	}
	
	file.WriteString(text)
}