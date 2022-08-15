# Interleave words from Files

## Steps to run the script

`
go run main.go -f1 inputs/tab/file1.txt -f2 inputs/tab/file2.txt
`

## Interleaving 

### Assumptions 

1. Words are required to be isolated by atleast one delimiter(space, comma)
2. These symbols could be allowed as part of the word Symbols ~$^+|<>[] For example., 20$
2. Following special characters such as !@#%*()_,; need not be considered for interleaving
4. Single length symbols are not considered as a proper word. Hence it's deleted

### Design Constraint 

1. Lack of defined input leads to open interpretation as mentioned in the assumptions
2. Funtionality accepts only 2 files as parameter. If number of inputs or type of inputs changes, then the function needs modification
3. No clear functional specifications