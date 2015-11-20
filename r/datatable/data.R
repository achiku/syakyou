# http://www.slideshare.net/sercantahaahi/feature-selection-with-r-in-jp
library(MASS)

my.seed <- 12345
set.seed(my.seed)
n_var <- 20
n_obs <- 60

Sigma <- matrix(0, nrow=n_var, ncol=n_var)
Sigma[1:5, 1:5] <- 0.9
diag(Sigma) <- 1

A <- mvrnorm(n=n_obs, rep(0, n_var), Sigma)
eps <- rnorm(n_obs, mean=0, sd=0.5)

w <- rep(0, n_var); w[c(3, 5, 10, 15, 16, 20)] <- c(5, 5, 2, -5, -5, -4)


Y <- A %*% w + eps
Y <- ifelse(Y < 0, 0, 1)
my.data <- as.data.frame(cbind(A, Y))
names(my.data)[1:5] <- paste("CORR_VAR", 1:5, sep="")
names(my.data)[n_var+1] <- "LABEL"
my.data$LABEL <- factor(my.data$LABEL)

my.data$MANY_CTG <- factor(rep(1:30, each=n_obs/30))
my.data$FEWER_CTG <- factor(rep(1:10, each=n_obs/10))

set.seed(my.seed)
