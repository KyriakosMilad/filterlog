# filter-log

Simple Go CLI tool to filter log file.

## Install

1. ```  git clone https://github.com/KyriakosMilad/filter-log ```
2. ```  cd filter-log ```
3. ```  go install . ```
4. Now you can use filter-log command from anywhere.

## Usage

You can filter the log file with any text you want, date, time, level, etc..., but text must not include coma which is
used to separate between multiple filters.

     filter-log -path <path_to_the_log_file> -filter <filter_to_search>

multiple filter options:

    filter-log -path <path_to_the_log_file> -filter <filter1,filter2,filter3,...>

## Help

    filter-log -help