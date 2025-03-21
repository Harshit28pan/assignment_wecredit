package main

import (
        "log"
        "math/rand"
        "os"
        "time"
)

func main() {
        logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
                log.Fatal(err)
        }
        defer logFile.Close()

        logger := log.New(logFile, "", log.LstdFlags)

        rand.Seed(time.Now().UnixNano())
        for i := 0; i < 100; i++ {
                level := rand.Intn(4)

                switch level {
                case 0:
                        logger.Println("[INFO] Everything is good.")
                case 1:
                        logger.Println("[WARN] Something might be wrong.")
                case 2:
                        logger.Println("[DEBUG] Checking a variable.")
                case 3:
                        logger.Println("[ERROR] Something went wrong.")
                }
                time.Sleep(time.Millisecond * 500)
        }
}
