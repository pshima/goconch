# goconch
GOCONnectivityCHecker is a simple TCP connectivity checker written in go.  Currently it only supports TCP connections and basic availability statistics and logging.

![GoConch](https://upload.wikimedia.org/wikipedia/commons/thumb/7/75/Sea_shell_%28Trinidad_%26_Tobago_2009%29.jpg/220px-Sea_shell_%28Trinidad_%26_Tobago_2009%29.jpg)

Put your ear to the shell - do you hear the ocean?

# example output

```
% ./goconch example.json
2015/11/04 23:08:49 [INFO] Waiting 5 seconds before starting
2015/11/04 23:08:49 [INFO] There are 3 total endpoints to be checked
2015/11/04 23:08:54 [INFO] Input Queue Length Is: 0
2015/11/04 23:08:54 [INFO] Output Queue Length Is: 0
2015/11/04 23:08:54 [FAIL] thissitedoesntexist.com:80 failed after 6.630973ms with error: dial tcp: lookup thissitedoesntexist.com: no such host
2015/11/04 23:08:54 [INFO] thissitedoesntexist.com:80 Availability: 0
2015/11/04 23:08:56 [SUCCESS] www.google.com:80: 2.058173093s
2015/11/04 23:08:56 [INFO] www.google.com:80 Availability: 100
2015/11/04 23:08:57 [FAIL] www.google2.com:80 failed after 3.005153033s with error: dial tcp: i/o timeout
2015/11/04 23:08:57 [INFO] www.google2.com:80 Availability: 0
2015/11/04 23:08:59 [INFO] Input Queue Length Is: 0
2015/11/04 23:08:59 [INFO] Output Queue Length Is: 0
2015/11/04 23:08:59 [FAIL] thissitedoesntexist.com:80 failed after 3.067223ms with error: dial tcp: lookup thissitedoesntexist.com: no such host
2015/11/04 23:08:59 [INFO] thissitedoesntexist.com:80 Availability: 0
2015/11/04 23:09:01 [SUCCESS] www.google.com:80: 2.028710688s
2015/11/04 23:09:01 [INFO] www.google.com:80 Availability: 100
2015/11/04 23:09:02 [FAIL] www.google2.com:80 failed after 3.000737785s with error: dial tcp: i/o timeout
2015/11/04 23:09:02 [INFO] www.google2.com:80 Availability: 0
2015/11/04 23:09:04 [INFO] Input Queue Length Is: 0
2015/11/04 23:09:04 [INFO] Output Queue Length Is: 0
2015/11/04 23:09:04 [FAIL] thissitedoesntexist.com:80 failed after 4.113368ms with error: dial tcp: lookup thissitedoesntexist.com: no such host
2015/11/04 23:09:04 [INFO] thissitedoesntexist.com:80 Availability: 0
```
