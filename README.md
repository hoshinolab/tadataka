# tadataka

![logo](./docs/tadataka-logo-small.png)

tadataka is the geospatial big data preprocessing tool for Japanese geospatial data. This tool is named after [Inō Tadataka](https://en.wikipedia.org/wiki/In%C5%8D_Tadataka), or Japanese greatest surveyor.

As of 25 Nov 2019, tadataka provides reverse geocoder (coordinate -> address) and subdivider for huge geospatial CSV data. In near future, this will also provide geocoder(address -> coordinate).

## Installation

Currently, tadataka only support UNIX-like OS. (GNU/Linux is recommended.) Windows will be supported soon.

### Redis

tadataka uses an in-memory database engine Redis to store address-coordinate data. Therefore you have to install Redis at first.

- [Redis](https://redis.io/) (version 4.0.0 or later is required)

Then, you have to run Redis by `$ redis-server` command.

### tadataka

Clone this repository and run `make`. Then add `./tadataka/bin` to `$PATH`. 
Pre-compiled binaries will be distributed in this GitHub repository.

### tadataka Basic setup

In first time, download Japanese address-coordinate data.

```
$ tadataka download
```

After download of address-coordinate data, run `stdby` command. This loads CSV files to Redis.

```
$ tadataka stdby
```


## Subcommands

Tadataka has sub commands.

- `download`: download Japanese address-coordinate data in `~/.tadataka`.
- `stdby`: import address-coordinate data from `~/.tadataka` to Redis. This subcommand is required to execute `gc` and `rgc`
- `subdiv`: subdivide large CSV file(s) into small CSV files with [Open Location Code (plus codes)](https://en.wikipedia.org/wiki/Open_Location_Code).
- `rgc`: reverse geocoder
- `version`: show version of Tadataka

### `download`

**requirement** : nothing (enough disk space)

stdby subcommand download address-coordinate data from Geospatial Information Authority of Japan (国土地理院, GSI) for geocoding/reverse geocoding into `~/.tadataka`

```
$ tadataka download
```

### `stdby`

**requirement** : downloaded address-coordinate data by `download` command, Redis

stdby subcommand imports address-coordinate data from `~/.tadataka` to Redis. You should run this subcommand before using `gc` and `rgc`

```
$ tadataka stdby
```


### `subdiv`

**requirement** : nothing (only target CSV files and enough disk space)

```sh
$ tadataka olc ./input/file/path.csv ./output/directory/path --lat 1 --lng 2 --header false
```

- `lat`: (number) Column number of latitude in CSV file. (begin from `0`)
- `lng`: (number) Column number of longitude in CSV file. (begin from `0`)
    - In CSV file like `user000,30.123456,145.456789,10,true`, `lat` is `1` and `lng` is `2`.
- `header`: (boolean) Whether CSV files have a header row or not. (default: `true`)


### `rgc`

**requirement** : downloaded address-coordinate data by `download` command, Redis

High speed reverse geocoder.

```sh
$ tadataka rgc ./input/file/path.csv ./output/file/path.csv --lat 1 --lng 2 --header false
```

- `lat`: (number) Column number of latitude in CSV file. (begin from `0`)
- `lng`: (number) Column number of longitude in CSV file. (begin from `0`)
    - In CSV file like `user000,30.123456,145.456789,10,true`, `lat` is `1` and `lng` is `2`.
- `header`: (boolean) Whether CSV files have a header row or not. (default: `true`)
