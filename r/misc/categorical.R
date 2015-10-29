summary(iris)
x <- iris$Sepal.Length
x
x2 <- cut( x, breaks = c(4,5,6,7,8), labels= c("4-5","5-6","6-7","7-8"), right=FALSE, ordered_result=TRUE )
x2
table(x2)

y <- iris$Sepal.Width
y2 <- cut(y, breaks = c(2.0,2.5,3.0,3.5,4.0,4.5), 
          labels= c("2.0-2.5","2.5-3.0","3.0-3.5","3.5-4.0","4.0-4.5"),
          right=FALSE, ordered_result=TRUE)
y2
table(y2)