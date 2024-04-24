# ping-checker

ping-checker helps to check ICMP status for multiple IPs with csv input file and generate the output.csv as a result.

### HOW TO USE

_Download the relevent os package from [here](https://github.com/pnkj-kmr/ping-checker/releases)_

_create a **input.csv** file into your current working directory_

```
<ip_address1>,<tag>
<ip_address2>,<tag>
<ip_address3>,<tag>

...
```

_OR_

```
<ip_address1>
<ip_address2>
<ip_address3>

...
```

_After creating the file run the executable binary as_

```
./pingchecker
```

### OUTPUT

_As a result **output.csv** file will be created after completion_

```
ip,tag,ping_result,error_if_any



```

### HELP

```
./pingchecker --help
```

:)
