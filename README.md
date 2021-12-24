# filterlog

Go CLI tool to filter log file.

## Installation

    $ go install github.com/KyriakosMilad/filterlog

verify installation:

    $ filterlog

## Options

display help menu:

    -help

path to the log file (required string):

    -path <string>

filter/s to search for (required string):

    -filters <string>

separator to use when you want to filter multiple options (optional string):

    -separator <string>

print the results into file (optional boolean) (default: false):

    -export_results <true||false>

## Usage

You can filter the log file with any text you want, date, time, level, etc...

    $ filterlog -path <path_to_the_log_file> -filters <filter_to_search>

multiple filters:

    $ filterlog -path <path_to_the_log_file> -filters <filter1,filter2,filter3> -separator <separator>

## Examples

one filter:

    $ filterlog -path example.log -filter INFO

multiple filters:

    $ filterlog -path example.log -filter "INFO,WARNING,ERROR" -separator ","
    $ filterlog -path example.log -filter "TRACE 04/21" -separator " "

output the results to a file:

    $ filterlog -path example.log -filters "ERROR" -export_results true
