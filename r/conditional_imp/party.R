library(party)
set.seed(my.seed)

pmodel <- cforest(LABEL ~ ., data=my.data, control=cforest_unbiased(mtry=5, ntree=51))
imp <- varimp(pmodel, conditional=F)
# imp <- varimp(pmodel, conditional=T)
fweights.party <- data.table(Features=names(imp), RealWeight=c(w, 0, 0), ConditionalImp=imp)
fweights.party <- fweights.party[order(-ConditionalImp)]
fweights.party$ConditionalImp <- as.integer(fweights.party$ConditionalImp * 100000)
fweights.party


result <- predict(pmodel, my.data, OOB=T, type="prob")
