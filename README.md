# Interleave Files

## Steps to run the script

`
go run main.go -f1 inputs/tab/file1.txt -f2 inputs/tab/file2.txt -d ,
`

### Interleaving 

I assume that single delimiter separates the words in each file - Handle double delimiter for now f1,,f2 -> [f1,f2] -remove empty strings

Space is the default delimiter 

Interleaving happens word by word and not line by line