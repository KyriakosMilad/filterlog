# filterlog

Go CLI tool to filter log file.

## Installation

    $ go get github.com/KyriakosMilad/filterlog

verify installation:

    $ filterlog

## Options

display help menu:

    -help

path to the log file (required string):

    -path <string>

filter/s to search for (required string):

    -filter <string>

separator to use when you want to filter multiple options (optional string):

    -separator <string>

print the results into file (optional boolean) (default: false):

    -output_file <true||false>

## Usage

You can filter the log file with any text you want, date, time, level, etc...

    $ filterlog -path <path_to_the_log_file> -filter <filter_to_search>

multiple filter options:

    $ filterlog -path <path_to_the_log_file> -filter <filter1,filter2,filter3> -separator <separator>

## Examples

one filter:

    $ filterlog -path example.log -filter INFO

multiple filters:

    $ filterlog -path example.log -filter "INFO,WARNING,ERROR" -separator ","
    $ filterlog -path example.log -filter "TRACE 04/21" -separator " "

output the results into file:

    $ filterlog -path example.log -filter "ERROR" -output_file true
