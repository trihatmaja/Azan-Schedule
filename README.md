# Azan-Schedule

This application to generate Azan Schedule or Sholat Schedule based on position(latitude, longitude, and timezone). This can be used at the smartphone if golang already support for android/ios.

#### How To Use
 1. Install dependency using goop, or using go get. (see Goopfile for the dependency)
 2. Build the application. Go build azan.go 
 3. Run ./azan --latitude=-6.18 --longitude=106.83 --timezone=+7 --city=jakarta

### Next
1. Insert the azan file azan.ogg
2. Generate cron job file based on the time generated to execute azan file. But the user should install the cron job file by himself

