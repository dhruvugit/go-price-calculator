1. As a idea we don't want to set the path for the packages for result and read file from internal packages, instead we want them to have in main function 
2. So if you see path in main, they were in package previously 
3. So to resolve above stuff we made our own struct in fileManager file for carrying input and outpur file paths(that looks better now)
4. So in code you'll be seeing FileManager struct as parameter in prices.go functions 
5. Now we also want the feature of swappable struct and their methods 
6. For that we have defined those method that the swappable structs are having 
7. Also we know in go that, go automatically looks for the structs which have those methods mentioned in the interface and treat them as they have implemented the inteface 
8. So interfaces in go are very useful if you want to have some flexiblity of using different method and struct as per the requirement.

Explanation of point A
Hey I am user A, i want my methods to use B(struct, method) and C(struct, method)
Go will say, hey just make sure to write a interface where those methods are named with input parameter and what they return then I will find those struct which have those methods in your case it should be B and C, also don't forge to use my interface type name for those struct 

This way you'll be able to add flexibility to add different type of structs and methods (B and C) to same user type(A)