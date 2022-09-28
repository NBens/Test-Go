# Test-Go

To finish this REST API, we need to:

- [X] Make a request to XKCD and fetch the last post
- [X] Save the last posts to a local memory cache (map[int]xkcdJSONStruct)
- [X] Check the difference between the last post and the saved posts in cache and fetch the missing posts
- [X] Make a request to PoorlyDrawnLines RSS and get the latest posts
- [ ] Combine the RSS feed with the [int]xkcdJSONStruct map in JSON format based on date
- [ ] GET Endpoint
- [ ] Clean up
- [ ] Testing