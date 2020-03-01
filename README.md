# hyperproductive
This is simple parallel library

## Use Case

### example1

```
	f := func() {
		fmt.Println("hoge")
	}

	h := hyperproductive.NewHyperProductiveGroup(10000, f)
	h.NotReportOperate()
```

### example2

```
	f := func() interface{} {
		return time.Now()
	}

	h := hyperproductive.NewHyperProductiveGroup(10000, f)
	results := h.IndividualOperate()

	for _, result := range results {
		fmt.Println(result)
	}
```

### example3

```
	f := func(incompletedParams ...interface{}) interface{} {
		params := incompletedParams[0].([]interface{})
		var strParams []string
		for _, param := range params {
			strParams = append(strParams, param.(string))
		}
		return strings.Join(strParams, "")
	}

	h := hyperproductive.NewHyperProductiveGroup(10000, f, "a", "b", "c")
	results := h.IndividualOperate()

	for _, result := range results {
		fmt.Println(result)
	}

```

## Available func.

```
func ()
func (...interface{})
func () interface{}
func (...interface{}) interface{}
```

## Availavle parametar
Pass the parameters used by the function to the variable length argument.
example3.


