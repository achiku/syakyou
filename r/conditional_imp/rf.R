library(data.table)
library(randomForest)

source("./data.R")

rfmodel <- randomForest(LABEL~., data=my.data, ntree=51, importance=T, scale=F)

imp1 <- importance(rfmodel, type=1)[, 1]
imp1 <- data.table(Feature=names(imp1), RealWeight=c(w, 0, 0), MeanDecreaseAccuracy=imp1)
imp1 <- imp1[order(-MeanDecreaseAccuracy)]

imp2 <- importance(rfmodel, type=2)[, 1]
imp2 <- data.table(Feature=names(imp2), RealWeight=c(w, 0, 0), MeanDecreaseGini=imp2)
imp2 <- imp2[order(-MeanDecreaseGini)]

varImpPlot(rfmodel)