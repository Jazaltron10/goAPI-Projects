# Weather Forecast Service

This is a simple weather forecast service implemented in Go using the Echo framework.

## Getting Started

Follow these steps to run the service:

1. Build the Docker image.
2. Run the Docker container.

...

"""
Only USA

Only JSON

We need to implement an endpoint (Echo Framework) that, given a list of cities, returns a list of forecast for those cities:

/weather?city=los%20angels,%20new%20york,chicago

For each single city we would like to know the daily forecast plus the forecasts of the next 2 days (72h)

Hints:

    in order to convert a string (New York) in a pair of coordinates (40.71427, -74.00597) you can use this service https://nominatim.openstreetmap.org/search?q=new%20york,usa&format=json - (Json marshal/unmarshal), structs and json tags (-, omitempty)... 

    in order to get the forecast of a pair of coordinates you can use this service https://api.weather.gov/points/40.71427,-74.00597 -> `properties`` -> `forecast` (this is a URL) -> make another request -> `properties` -> `periods`
    the forecast for the day comes back in different "periods" ... Check the first and more, if its the current time if its not discard... and accept any that are now or later (under 72h - 2 days)

What we would like to see

    1 .tiny service that works just fine (that we can launch it ideally via docker or command line).
    2. thoughtful design & tests.
    3. implement a file cache (memory or file based) system without using any framework.
    ~ ReadMe file with the instructions for using the service.

response.json
```
{
  "forecast": [
    {
      "name": "Los Angels",
      "detail": [
        {
          "startTime": "2022-05-13T00:00:00-07:00",
          "endTime": "2022-05-13T06:00:00-07:00",
          "description": "Mostly clear, with a low around 58. East northeast wind 5 to 10 mph."
        },
        {
          "startTime": "2022-05-14T18:00:00-07:00",
          "endTime": "2022-05-14T06:00:00-07:00",
          "description": "Mostly clear, with a low around 61. North wind 5 to 15 mph."
        },
        {
          "startTime": "2022-05-15T18:00:00-07:00",
          "endTime": "2022-05-15T06:00:00-07:00",
          "description": "Patchy fog after 11pm. Mostly cloudy, with a low around 61. Southeast wind 5 to 15 mph, with gusts as high as 20 mph."
        }
      ]
    },
    {
      "name": "New York",
      "detail": [
        {
          "startTime": "2022-05-13T02:00:00-04:00",
          "endTime": "2022-05-13T06:00:00-04:00",
          "description": "A slight chance of rain and a slight chance of drizzle. Cloudy, with a low around 60. Northeast wind around 6 mph. Chance of precipitation is 20%."
        },
        {
          "startTime": "2022-05-14T18:00:00-04:00",
          "endTime": "2022-05-14T06:00:00-04:00",
          "description": "A slight chance of rain and a slight chance of drizzle before 11pm, then areas of fog and a slight chance of rain and a slight chance of drizzle. Cloudy, with a low around 61. Southeast wind 2 to 6 mph. Chance of precipitation is 20%."
        },
        {
          "startTime": "2022-05-15T18:00:00-04:00",
          "endTime": "2022-05-15T06:00:00-04:00",
          "description": "A chance of rain and a chance of drizzle before 8pm, then a chance of rain and a chance of drizzle and patchy fog. Cloudy, with a low around 62. South wind 5 to 9 mph. Chance of precipitation is 40%. New rainfall amounts less than a tenth of an inch possible."
        }
      ]
    }
  ]
}
```

Echo framework (server)
Client -> to make http GET requests -> To any links mentioned/needed
A lot of Json to Structs

Sending out the response -> 
"""