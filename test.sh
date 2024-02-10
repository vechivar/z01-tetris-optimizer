for arg in $(ls sample)
do
echo "------------------------"
echo "TESTING : $arg"
echo "------------------------"
go run . "sample/$arg"
echo "------------------------"
done