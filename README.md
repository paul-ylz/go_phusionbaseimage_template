# golang & phusion/baseimage template

This template is my 'project starter kit'. It uses gin as main, one-and-only server but can be easily swapped out with any other go app server of your choice. Similarly, postgres is included but can easily be exchanged for another database. It is merely my current go-to stack.

It is based on a work flow of developing locally using docker (and docker-compose), and can be easily adapted to deploy via a shell script to a integration / staging server (see script at deploy/staging.sh). It is expected that production is on some service like ECS, and that the produced image is easily adaptable to whatever method chosen.

Suggestions / discussion for improvement are most welcome. I am relatively new to Golang, coming from a background of developing with Sinatra/Passenger and then Sinatra/Docker. As a developer that has to hop from project to project frequently, I find the phusion/baseimage fat container paradigm a great time and energy saver, especially if the project requires routine jobs (cron). Routine jobs may be better organized on Cloudwatch Events in 2018 but if you are a small shop with many projects in development, it is faster and more encapsulated to tweak a crontab that lives in the project dir.

## Next:
### Standard Package Layout
The next thing I would like to add to this template is a seed for using [Ben Johnson's Standard Package Layout] (https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).  I have been looking for a good way to organize Go code for a while and after reading that post that TJHolowaychuk sortof-endorsed, have to admit looks like the best way I've seen yet.
