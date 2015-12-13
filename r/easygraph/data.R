
a <- as.numeric(c(0:10))
b <- as.numeric(c(10:20))
c <- as.numeric(c(20:30))
flag <- c(1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0)

df <- data.frame(a=a, b=b, c=c, flag=flag)

easygraph <- function(df) {
    on <- df[df$flag == 1,]
    on <- on[,colnames(df) != "flag"]
    off <- df[df$flag == 0,]
    off <- off[,colnames(df) != "flag"]
    cols <- colnames(on)
    for (x in cols) {
      hist(on[x], col="#ff00ff40")
      hist(off[x], col="#0000ff40", add = TRUE)
      legend("topright", legend=c("on", "off"), col=c("red", "blue"), pch=c(1, 1))
  }
}

easygraph(df)
