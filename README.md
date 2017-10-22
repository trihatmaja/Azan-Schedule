# Azan-Schedule

This application to generate Azan Schedule or Sholat Schedule based on position(latitude, longitude, and timezone). This can be used at the smartphone if golang already support for android/ios.

#### How To Use
 1. Run `govendor sync`
 2. Build the application using `make build`
 3. Run `./azan --latitude=-6.18 --longitude=106.83 --timezone=+7 --city=jakarta`

### Next
1. Insert the azan file azan.ogg
2. Generate cron job file based on the time generated to execute azan file. But the user should install the cron job file by himself

### The MIT License (MIT)
Copyright (c) 2015 Dr. T. Djamaluddin, Wastono ST

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

