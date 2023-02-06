# Run the app

```shell
cp .env.dist .env && docker-compose up --build
```

### Available routes
`GET /locations`


creates a new location

`POST /locations`

`GET /locations/:location_id`

`GET /locations/:location_id/devices`

`GET /locations/:location_id/outlets`


creates a new device

`POST /devices`


`GET /devices/:device_id`

Creates a network outlet to which the device is connected

`POST /outlets`