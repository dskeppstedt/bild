# bild
📝 ➡️ 📸 ➡️ 🖼

_Bild_ is a simple tool that takes screenshots of websites.

# Build and run
Follow these steps to build and run this program. 

1. Install [go](https://golang.org/doc/)
2. Install [wkhtmltopdf](https://wkhtmltopdf.org)
3. > mkdir screenshots

4. create a file with urls

5. > go run bild.go urls.txt


# How to scale this?

_bild_ is a pretty slow and bad service. Improving it would require splitting it up into smaller services that could run on a set of servers. 

The services could be:
 * the frontend where applications or users can submit jobs.
 * the request service that gets a URL and saves a html request to a distributed file system.
 * the html to image converts service that would use the html.

 Another improvement would be to instead of running wkhtmltopdf as a child process it could be added as a c++ library to the html/image converter service. 

The frontend service would add all the URLs to a database, e.g mongodb, and push the URLs to a message queue. The request service on a given server would take the next URL in the queue and process it and once done adds the file path another message queue. The html/image converter takes the next file path from the queue and processes it and marks it as done in the database. The frontend could show the status of each URL in the database and link to the actual screenshot.

I believe this setup can be able to scale to handle the required number of requests.
