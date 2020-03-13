# goweet 

CLI tool for tweeting using Twitter tokens.


**Usage:**
```
goweet  -handle <twitter handle> \ 
		-tweet <tweet body> \
		-consumer_key <consumer key for your twitter app> \
		-consumer_secret <consumer secret for your twitter app> \
		-access_token <access token for your twiter app> \
		-access_token_secret <access token secret for your twiter app>
		
```

Additionally, you can set the environment variables instead of using flags
which are shown as follows:

- `GOWEET_HANDLE --> handle`
- `GOWEET_CONSUMER_KEY --> consumer_key`
- `GOWEET_CONSUMER_SECRET --> consumer_secret`
- `GOWEET_ACCESS_TOKEN --> access_token`
- `GOWEET_ACCESS_TOKEN_SECRET --> access_token_secret`

**Acknowledgement:**
- https://github.com/dghubble/go-twitter