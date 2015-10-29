x <- 10
y <- 10
x
y
x + y
z = 'test string'
z

xs <- c(1, 2, 3, 4, 5)
xs
ys <- c(1:10)
ys
zs <- ys * 10
zs

categories <- c('great', 'ok', 'suck')
categories

m1 <- matrix(c(1, 2, 3, 4, 5, 6, 7, 8), nrow=2, ncol=4)
m1
m1[,1]
m1[1,]

a <- c(1:4)
b <- c(5:8)
df <- data.frame(a, b)
df
df <- data.frame(a, b, a)
df

l1 <- list(a, b)
l1
l2 <- list(A=a, B=b)
l2
names(l2)
str(l2)
l2$A

i1 <- c(1:10)
val3 <- c(101:110)
val4 <- c(41:50)
df1 <- data.frame(Sample=i1, Value1=val3, Value2=val4)
df1

i2 <- c(31:40)
val1 <- c(501:510)
val2 <- c(61:70)
df2 <- data.frame(Sample=i1, Value1=val1, Value2=val2)
df2
