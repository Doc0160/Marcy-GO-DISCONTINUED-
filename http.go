package main
// import(
	// "net/http"
	// "sync"
	// "time"
	// "io/ioutil"
	// "fmt"
// )
// type httpClient struct{
	// Client *http.Client
	// mutex  *sync.Mutex
// }
// func NewHTTPClient()httpClient{
	// c:= httpClient{}
	// c.mutex = &sync.Mutex{}
	// c.Client=&http.Client{
		// Timeout: time.Second * 5,
		// Jar: nil,
	// }
	// return c
// }
// func (c*httpClient)Lock(){
	// c.mutex.Lock()
// }
// func (c*httpClient)Unlock(){
	// c.mutex.Unlock()
// }
// func (c*httpClient)Get(url string) ([]byte, error) {
	// req, err := http.NewRequest("GET", url, nil)
	// if err!= nil{
		// return nil,err
	// }else{
		// req.Header.Set("User-Agent", "I do things ? ... I'm a stupid slack bot ! (tristan.magniez@viacesi.fr)")
		// c.Lock()
		// r, err := c.Client.Do(req)
		// if err != nil {
			// return nil, err
		// } else {
			// buf, err := ioutil.ReadAll(r.Body)
			// if err != nil {
				// return nil, err
			// } else {
				// r.Body.Close()
				// c.Unlock()
				// return buf, nil
			// }
		// }
	// }
// }