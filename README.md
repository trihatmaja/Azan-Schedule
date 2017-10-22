# Azan-Schedule

This application to generate Azan Schedule or Sholat Schedule based on position(latitude, longitude, and timezone). This can be used at the smartphone if golang already support for android/ios.

## How To Use

### Get Dependency

1. Run `govendor sync`

### Build CLI

If you want to generate files that set the azan schedule, build the cli

1. Build the application using `make cli`
2. Run `./azan_cli --latitude=-6.18 --longitude=106.83 --timezone=+7 --city=jakarta`

### Build API Server

If you want to setup server for azan schedule, you can build the api server

#### Dependency

1. MYSQL
2. MEMCACHED

#### Build

1. Build the api server using `make api`
2. Change env.sample to .env and edit the value as you need
3. Run `./azan_api`
4. Apps will listen on port 1234

## How To Develop

### Calculation

You may develop your own calculation implementation, see `provider.go` for the contract about the calculation function and place it under calculation package

### Database

You may develop your own database implementation (if you want to use other db other than mysql), see `provider.go` for the contract about database function and place it under database package

### Cache

You may develop your own cache implementation (if you want to use other cache engine other than memcache), see `provider.go` for the contract about cache function and place it under cache package

## The MIT License (MIT)
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

