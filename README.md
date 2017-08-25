# Golang-Iteration-Performace
Track Performance &amp; Memory Allocation in 20.000.000 iteration in Golang

# Author
Thomas wangsa
github: https://github.com/Thomas-wangsa

# 1. Basic Information :
Benchmark System :
A. Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
B. 2x 4GiB SODIMM Synchronous 2133 MHz
C. go version go1.8.3 linux/amd64
D. Ubuntu 16.04.2 LTS(Xenial),non fresh installation
Library : https://github.com/julienschmidt/httprouter

# 2. Performance Result :
Basically i separated this test onto 4 Scenario :

I. Scenario1 : only looking for process time of each iteration.
II. Scenario2 : only looking for how many bytes used in heap memory allocated every iteration. 
III. Scenario3 : only looking for Allocation used every iteration.
IV. Scenario4 : Test Performance Scenario1-3 Together in 1 API Endpoint.

A. Scenario1 :
Moreless need 7341913 nano second or 0.36709565 nano second/1 Iteration

B. Scenario2 : 
Moreless 575592 bytes heap memory used when we hit the API before iteration and then will up to 32 bytes in every iteration. 
Source : https://golang.org/pkg/runtime/

C. Scenario3 :
Moreless need 560 - 575 Kb memory allocation when we hit the API before iteration and then will up to 0.03030303 kb in every iteration.
Source : https://golang.org/pkg/runtime/

D. Scenario4 :
Moreless 296333162956 nano second or 14816 nano second/1 Iteration
Moreless 575592 bytes heap memory used when we hit the API before iteration and then will up to 48 bytes in every iteration.
Moreless need 560 - 575 Kb memory allocation when we hit the API before iteration and then will up to 0.045454545 kb in every iteration.

# 3. Additional Info
Clone this Repository and run the main.go.
