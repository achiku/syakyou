library(lattis)

summary(iris)
str(iris)

histogram(~ Sepal.Length, data=iris)
histogram(~ Sepal.Length | Species, data=iris)
plot(Petal.Length ~ Sepal.Length, data=iris)

plot(iris)
pairs(iris[1:4])
pairs(iris[1:4], 
      main = "Iris Data setosa/versicolor/virginica",
      pch = 21,
      bg = c("red", "green3", "blue")[unclass(iris$Species)]
      ) 


plot( iris$Species, iris$Sepal.Length ) 
par ( new=T )
plot( iris$Sepal.Length, iris$Sepal.Width, pch = 21, bg = c("red", "green3", "blue") [unclass ( iris$Species ) ] ) 
par ( new=T )
plot( iris$Sepal.Length, iris$Sepal.Width, xlim=c( 2.0, 8.0 ), ylim=c( 2.0, 8.0 ))

par ( mfcol = c(1,3) )
plot( iris$Species, iris$Sepal.Length ) 
plot( iris$Sepal.Length, iris$Sepal.Width, pch = 21, bg = c("red", "green3", "blue") [unclass ( iris$Species ) ] ) 
plot( iris$Sepal.Length, iris$Sepal.Width, xlim=c( 2.0, 8.0 ), ylim=c( 2.0, 8.0 ))
par ( mfrow = c(1,1) )
