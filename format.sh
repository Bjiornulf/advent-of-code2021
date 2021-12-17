for dir in ./day*
do
	cd $dir
	go fmt .
	cd ..
done

