# web-service-gin

Demo RESTFUL API webservice built from Go tutorial:

https://go.dev/doc/tutorial/web-service-gin

## Run service
To run first build and run the package from `/web-service-gin` using:
```
go build && go run main.go
```

## GET albums
Once running you can access the album directory either directly or via cURL at:
```
localhost:3001/albums
```
Or via id at:
```
localhost:3001/albums/[ID]
```

## POST albums
In order to add a new entry, you can make use of the POST endpoint /albums, for example:
```
curl http://localhost:3001/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
