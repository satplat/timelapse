# TimeLapse
Select a region and get a timelapse video for it.
This project is powered by **satplat** APIs, you can find it from apieco.ir.  
At first this program will submit an area using provided coordinates, then it 
will download RGB images for this year (Every month may contain over 3 images) 
and saves it into **.image** directory and this directory will be deleted as 
soon as video is created.

## What's new
In the latest version you can select an index to download the index for the 
image that is downloaded. Go to _internal/downloader.go_, line 73. Put the index
you want, and it will create a csv file for every image.
 
## Dependencies
For converting images into video, you need ffmpeg script, which you can download it by running the command below.  
```$ /bin/bash ./ffmpeg```

## Select polygon
In order to select an area you need to put the coordinates of that area into _land.txt_ file provided inside the package. This file should be formatted like this:
> StartLat StartLon,MiddlesLat MiddlesLon,EndLat EndLon,StartLat StartLon

As an example, the following polygon is a correct area polygon points (coordinate) of Taleghan Lake.

> 36.1921193138716 50.6240608162892,36.1601194008489 50.624552535992,36.1602016419336 50.7429287600999,36.1925107789363 50.743325337844,36.1921193138716 50.6240608162892

## Run
In order to run please visit [apieco.ir](https://apieco.ir/api/%D9%88%D8%A8%D8%B3%D8%B1%D9%88%DB%8C%D8%B3-%D8%A7%D8%B7%D9%84%D8%A7%D8%B9%D8%A7%D8%AA-%D8%B2%D9%85%DB%8C%D9%86%D9%87%D8%A7%DB%8C-%DA%A9%D8%B4%D8%A7%D9%88%D8%B1%D8%B2%DB%8C-agricult/#)
to get api token.
Then open the **main.go** file and put the token in the token variable.  
After that you can run the program using ( You would need a go compiler )  
```$ go run main.go```  or create a binary file  
```$ go build``` and ```$ ./App```


## Modifications
You can modify the **.bash.sh** in order to save the images.  
If you modify the ***.go** files, you should build the project using the command below, which you would need a go compiler.  
```$ go build```

## Contribution
Any contribution is welcomed. 
