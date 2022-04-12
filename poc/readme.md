### flow 

- user joins using react app with a username (unique)
- assign the user with a socket session
- save the username v session id in the redis db 
- send the session with list of all the users (except him)
- as soon as a user joins, other sessions will be sent with updated list 
- as soon as a user leaves, other sessions will be sent with updated list 
- a user clicks on username and sends a direct message 
- all the messages are saved in redis 
- when user joins back with the same username, all messages are retrieved 

### infinite for loop 

What's happening is that `ReadMessage()` or `ReadJSON()` is not "breaking" out of your loop. It's just a blocking call. It's waiting for next request to come in or for the connection to be closed by the client. In other words, ReadMessage() does 2 things: first it waits for a message, then it reads it. Alternatively, when the client closes the connection, ReadMessage() will return a non-nil error. That will cause your code to print "Error in read" and return leaving both the for loop and the whole function.