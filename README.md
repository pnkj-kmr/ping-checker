# ping-checker

ping-checker helps to check ICMP status for multiple IPs with csv input file and generate the output.csv as a result.

### HOW TO USE

_Download the relevent os package from [here](https://github.com/pnkj-kmr/ping-checker/releases)_

_create a **input.csv** file_

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

_OR_

_create a **input.json** file_

```
[
    {
        "ip": "127.0.0.1",
        "tag": "test1", # omitempty
        "count": 4      # omitempty
        "timeout": 4    # omitempty
    },
    ...
]
```

_After creating the file run the executable binary as_

```
./pingchecker
```

### OUTPUT

_As a result **output.csv** file will be created after completion_

```
ip,tag,result,packetloss,error


```

_OR_

_As a result **output.json** file will be created after completion_

```
[
  {
    "input": { "ip": "127.0.0.1", ... },
    "ok": true,
    "error": "",        # omitempty
    "packet_loss": 0,
    "avg_rtt": 265500,
    "std_dev_rtt": 70687
  },

  ...
]
```

### HELP

```
./pingchecker --help

----------------------
Usage of ./pingchecker:
  -c int
        packet count (default 4)
  -f string
        give a file name (default "input.csv")
  -json
        file type - default[csv]
  -o string
        output file name (default "output.csv")
  -t int
        ping timeout [secs] (default 5)
  -w int
        number of workers (default 4)

-------
Example:

./pingchecker -f x.csv -t 30 -w 20

```

## OPTIONS

---

### `-f` (DEFAULT: "input.csv")

Different input file if any

```
./pingchecker -f ./new_input_file.csv
```

### `-o` (DEFAULT: "output.csv")

Different output file if required

```
./pingchecker -o new_output.csv
```

### `-w` (DEFAULT: 4)

Increase worker processes if needed

```
./pingchecker -w 1000
```

### `-t` (DEFAULT: 5 (secs))

Increase end IP ping timeout

```
# timeout 10 seconds
./pingchecker -t 10
```

### `-json` (DEFAULT: csv)

Change of input/output to json

```
./pingchecker -json
```

### `-c` (DEFAULT: 4)

Ping packet count to send

```
./pingchecker -c 2
```

:)
