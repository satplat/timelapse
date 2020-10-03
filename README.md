# TimeLapse
Select a region and get a timelapse video for it.
This project is powered by **satplat** APIs, you can find it from apieco.ir.  
At first this program will submit an area using provided coordinates, then it will download RGB images for this year (Every month may contain over 3 images) and saves it into **.image** directory and this directory will be deleted as soon as video is created.
 
## Dependencies
For converting images into video, you need ffmpeg script, which you can download it by running the command below.  
```$ /bin/bash ./ffmpeg```

## Select polygon
In order to select an area you need to put the coordinates of that area into _land.txt_ file provided inside the package. This file should be formatted like this:
> StartLat StartLon,MiddlesLat MiddlesLon,EndLat EndLon,StartLat StartLon

As an example, the following polygon is a correct area polygon points (coordinate) of Taleghan Lake.

> 36.1921193138716 50.6240608162892,36.1601194008489 50.624552535992,36.1602016419336 50.7429287600999,36.1925107789363 50.743325337844,36.1921193138716 50.6240608162892

## Run
Just run App as bellow:

```$ ./App```


## Modifications
You can modify the **.bash.sh** in order to save the images.  
If you modify the ***.go** files, you should build the project using the command below, which you would need a go compiler.  
```$ go build```

## Contribution
Any contribution is welcomed. 
