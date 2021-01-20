import pandas as pd


if __name__ == '__main__':
    df1 = pd.DataFrame({
        'data1': range(6),
        'key1': ['A', 'B', 'C', 'A', 'B', 'C'],
        'key2': ['A', 'B', 'C', 'A', 'B', 'C']
        })
    df2 = pd.DataFrame({
        'data2': range(3),
        'key1': ['A', 'C', 'D'],
        'key2': ['A', 'E', 'B']
        })

    # print(df1)
    # print(df2)

    # dfm1 = pd.merge(df1, df2)
    # print(dfm1)

    print(pd.merge(df1, df2, on=['key1', 'key2']))

    dfm2 = pd.merge(df1, df2, how='outer')
    print(dfm2)

    print(df1.drop(['key1', 'data1'], axis=1))
    print(df1)
