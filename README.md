# dbconnect
Offers endpoints to add or read contents from a database. The database used is currently Redis but more databases could be incorporated with time.

Offers the following endpoints:

**:9050/addtodb** - adds a key, field, value to the database
The request body should be of this form:
{"key":"somekey",
"field":"somefield",
"value":"somevalue"}

**:9050/getfromdb** - gets a value from the database when a key and field are specified. The request body should be of the same form as mentioned above.

### Pre-requisites
The only prerequisite (currently) is that Redis should be running on port 6379.

**To run**, migrate to the project directory and type go run main.go in a terminal.

This database API is being used by github.com/meghashyamc/scraper. 
