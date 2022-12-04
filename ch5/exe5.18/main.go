package main

// func fetch(url string) (filename string, n int64, err error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", 0, err
// 	}
// 	defer resp.Body.Close()

// 	local := path.Base(resp.Request.URL.Path)
// 	if local == "/" {
// 		return "", 0, err
// 	}
// 	f, err := os.Create(local)
// 	if err != nil {
// 		return "", 0, err
// 	}
// 	defer func() {
// 		if closeErr := f.Close(); err == nil {
// 			err = closeErr
// 		}
// 	}()

// 	n, err = io.Copy(f, resp.Body)
// 	return local, n, err
// }
