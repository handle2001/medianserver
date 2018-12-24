# Medianserver v1.0.0

## Overview:
Medianserver performs a simple task: it accepts a named set of integers in JSON format and returns their median value.

## Examples:

```
$ curl -X POST -d '{"name":"deezNuts","set":[90,92,93,88,95,88,97,87,98]}' http://localhost:8080 | jq
```
Returns:

```
{
  "Name": "deezNuts",
  "Type": "Median",
  "Value": [
    92
  ]
}
```