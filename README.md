# filter-log

Simple Go CLI tool to filter log file.

## Install

1. ```  git clone https://github.com/KyriakosMilad/filter-log ```
2. ```  cd filter-log ```
3. ```  go install . ```
4. Now you can use filter-log command from anywhere.

## Options

display help menu:

    -help

path to the log file (required string):

    -path <value>

filter/s to search for (required string):

    -filter <value>

separator to use when you want to filter multiple options (optional string):

    -separator <true||false>

print the results into file (optional boolean) (default: false):

    -output_file <true||false>

## Usage

You can filter the log file with any text you want, date, time, level, etc...

     filter-log -path <path_to_the_log_file> -filter <filter_to_search>

multiple filter options:

    filter-log -path <path_to_the_log_file> -filter <filter1,filter2,filter3> -separator <separator>

## Examples

one filter example:

    filter-log -path example.log -filter INFO

multiple filters example:

    filter-log -path example.log -filter "INFO,WARNING,ERROR" -separator ","
    filter-log -path example.log -filter "TRACE 04/21" -separator " "

output the results into file:

    filter-log -path example.log -filter "ERROR" -output_file true
