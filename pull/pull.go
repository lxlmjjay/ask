package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"os/exec"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/pull", func(c *gin.Context) {

		f, err := os.OpenFile("pull.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		command := "./pull.sh"
		fmt.Println("开始执行......")
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(f)
		if err != nil {
			log.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		} else {
			log.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
		}


	})

	r.Run(":8089") // listen and serve on 0.0.0.0:8080
}
