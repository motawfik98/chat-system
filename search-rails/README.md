This project only contains the search endpoint which could be accessed using the following URL

POST `http://localhost:4500/applications/:appToken/chats/:chatNumber/search`

With a json body of two key value pairs
1. `message` which is the text to search for
2. `operator` a string value to indicate whether an or/and search is required

It's first access the database to make sure that a chat with the specified app token and chat number exists.

Then it'll connect to elasticsearch server to fetch the results.

*N.B.* The endpoint will throw an error if the index hasn't been created yet. So, make sure to create messages using the Golang app before trying to search for it.