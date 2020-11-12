package handlers

func queryBuyers() string {
	resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")
	if err != nil {
		log.Fatalln(err)
	}