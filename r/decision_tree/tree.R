library(rpart)
library(rpart.plot)

data(iris)
iris.rp <- rpart(Species~., data=iris)
iris.rp
rpart.plot(iris.rp, type=1, uniform=T, extra=1, under=1, faclen=0)
