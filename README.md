# Tadataka

<img src="./docs/tadataka-logo.svg" alt="logo" style="max-width:40%;">

Tadataka is the geospatial big data preprocessing tool. This tool is named after [Inō Tadataka](https://en.wikipedia.org/wiki/In%C5%8D_Tadataka), or Japanese greatest surveyor.

## Installation

//TODO

## Usage

### Basic setup

```
$ tadataka prep
$ tadataka stdby
```

## Subcommands

Tadataka has sub commands.

- `prep`: create `~/.tadataka` direcotry and download address-coordinate data from Geospatial Information Authority of Japan (国土地理院, GSI) for geocoding/reverse geocoding.
- `stdby`: import address-coordinate data from `~/.tadataka` to Redis. This subcommand is required to execute `gc` and `rgc`
- `subdiv`: subdivide large CSV file(s) into small CSV files with [Open Location Code (plus codes)](https://en.wikipedia.org/wiki/Open_Location_Code).
    - former: `olc`
- `gc`: geocoder (not implemented yet)
- `rgc`: reverse geocoder
- `version`: show version of Tadataka

### `prep`

**requirement** : nothing (enough disk space)

stdby subcommand download address-coordinate data from Geospatial Information Authority of Japan (国土地理院, GSI) for geocoding/reverse geocoding into `~/.tadataka`

```
$ tadataka prep
```

### `stdby`

**requirement** : downloaded address-coordinate data by `prep` command, Redis

stdby subcommand imports address-coordinate data from `~/.tadataka` to Redis. You should run this subcommand before using `gc` and `rgc`

```
$ tadataka stdby
```


### `subdivide`

**requirement** : nothing (only target CSV files and enough disk space)

former: `olc` command.


```sh
$ tadataka olc ./input/file/path.csv ./output/directory/path --lat 1 --lng 2 --header false
```

- `lat`: (number) Column number of latitude in CSV file. (begin from `0`)
- `lng`: (number) Column number of longitude in CSV file. (begin from `0`)
    - In CSV file like `id000,30.123456,145.456789,10,true`, `lat` is `1` and `lng` is `2`.
- `header`: (boolean) Whether CSV files have a header row or not. (default: `true`)



## Tadataka Config File

In some sub commands (like `olc` multiple file mode, `gc`, `rgc`), Tadataka requires a config file like below.

```json
{
    "input_dir":"~/input/directory/path",
    "output_dir":"~/output/directory/path",
    "lat_column":2,
    "lng_column":3,
    "header_row":true
}
```

- `input_dir`: (string) input directory
    - Tadataka reads all files in this designated directory
    - NOTICE: In current version, Tadataka does **NOT** support recursive mode.
- `output_dir`: (string) output directory
    - Tadataka writes CSV files in this directory keeping with file names in `input_dir`
- `lat_column`: (number) Column number of latitude in CSV file. (begin from `0`)
- `lng_column`: (number) Column number of longitude in CSV file. (begin from `0`)
    - In CSV file like `id000,30.123456,145.456789,10,true`, `lat_column` is `1` and `lng_column` is `2`.
- `header_row`: (boolean) Whether CSV files have a header row or not.