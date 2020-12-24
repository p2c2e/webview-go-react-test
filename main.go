package main

import (
	"encoding/json"
	"net/http"
	"os"
	//  "strings"
	//  "fmt"
	//  "io"

	"github.com/gobuffalo/packr"
	"github.com/webview/webview"
)

// Message : struct for message
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Bind folder path for packaging with Packr
	folder := packr.NewBox("./ui/build")

	// Handle to ./static/build folder on root path
	http.Handle("/", http.FileServer(folder))

	// Handle to showMessage func on /hello path
	http.HandleFunc("/hello", showMessage)

	http.HandleFunc("/recvfile", receiveFile)

	// Run server at port 8000 as goroutine
	// for non-block working
	go http.ListenAndServe(":8000", nil)

	// Let's open window app with:
	//  - name: Golang App
	//  - address: http://localhost:8000
	//  - sizes: 800x600 px
	//  - resizable: true
	w := webview.New(true)
	w.SetTitle("Golang App")
	w.SetSize(800, 600, webview.HintNone)
	w.Bind("mycallback", cbfunc)
	w.Bind("hello", func() string { os.Create("/tmp/touch.hello"); return "World!" })

	w.Bind("noop", func() string {
		return "hello"
	})
	w.Bind("add", func(a, b int) int {
		return a + b
	})
	w.Bind("quit", func() {
		w.Terminate()
	})
	/*
	   w.Navigate(`data:text/html,
	           <!doctype html>
	           <html>
	               <body>hello</body>
	               <script>
	                   window.onload = function() {
	                       document.body.innerText = ` + "`hello, ${navigator.userAgent}`" + `;
	                       noop().then(function(res) {
	                           console.log('noop res', res);
	                           add(1, 2).then(function(res) {
	                               console.log('add res', res);
	                           });
	                       });
	                   };
	               </script>
	           </html>
	       `)
	*/
	w.Navigate("http://localhost:8000/")
	w.Run()
}

func cbfunc(s string) string {
	os.Create("/tmp/touch." + s)
	return "Hello"
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	// Create Message JSON data
	message := Message{"World"}

	// Return JSON encoding to output
	output, err := json.Marshal(message)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write output
	w.Write(output)
}

func receiveFile2(w http.ResponseWriter, r *http.Request) {
	// Create Message JSON data
	message := Message{"TestBlock"}

	// Return JSON encoding to output
	output, err := json.Marshal(message)

	// Catch error, if it happens
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write output
	w.Write(output)
}

func receiveFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // limit your max input length!
	// in your case file would be fileupload
	//var buf bytes.Buffer
	//r.ParseForm()
	message := Message{"DEFAULT A"}
	//for k,_ := range r.Form {
	//	message = Message{ k  + "X"}
	//}
	_, header, err := r.FormFile("fileselect")
	//
	//message := Message{"DEFAULT"}
	if err != nil {
		message = Message{err.Error() + r.Method + " " + string(r.ContentLength) + "="}
	} else {
		//name := strings.Split(header.Filename, ".")
		//fmt.Printf("File name %s\n", name[0])
		message = Message{header.Filename}
	}
	/*
	       defer file.Close()

	       // Copy the file data to my buffer
	       io.Copy(&buf, file)
	       // do something with the contents...
	       // I normally have a struct defined and unmarshal into a struct, but this will
	       // work as an example
	       contents := buf.String()
	       fmt.Println(contents)
	       // I reset the buffer in case I want to use it again
	       // reduces memory allocations in more intense projects
	       buf.Reset()
	   	// Create Message JSON data
	*/
	//message := Message{"Called"}

	// Return JSON encoding to output
	output, _ := json.Marshal(message)

	// Catch error, if it happens
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Write output
	w.Write(output)
}
