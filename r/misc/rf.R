library(randomForest)

train <- seq.int(1, 150, by=2)
test <- setdiff(1:150, train)

Y <- iris[, colnames(iris) == "Species"]
X <- iris[, colnames(iris) != "Species"]

# tmpmodel <- tuneRF(X, Y, doBest=T)
# tmpmodel -> mtry=4 looks good

iris.rf <- randomForest(Species~., data=iris[train,], mtry=4)
iris.pred <- predict(iris.rf, iris[test,])
table(iris[test,5], iris.pred)

iris.pred <- predict(iris.rf, iris[test,], "prob")
iris.pred
