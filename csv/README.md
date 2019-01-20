# Little programm demonstrationg the csv library

This package demonstrates with a little programm the csv module from the go std library. To have some data the programm loads a csv file from a https://archive.ics.uci.edu and uses this data to demonstrate the csv package.
If you want to load data form the file system you can use the file open from the os package.

    file, err := os.OpenFile("", os.O_RDONLY, 0)

The file will implement the io.Reader interface which is required by the "csv.Reader". Luckily the http.Response.Body also implements the io.Reader interface so we can stick the http request result directly into a "csv.Reader".

The rest of the programm demonstrates how to read and write csv values.