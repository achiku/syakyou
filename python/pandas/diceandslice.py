import pandas as pd


if __name__ == '__main__':
    df1 = pd.DataFrame({
        'data1': range(6),
        'key1': ['A', 'B', 'C', 'A', 'B', 'C'],
        'key2': ['A', 'B', 'C', 'A', 'B', 'C']
        })

    # bool index
    print(df1[df1['data1'] >= 4])

    # query
    print(df1.query('data1 >= 4'))
    print(df1.query('data1 >= 4 and key1 == "B"'))

    print(df1)
    # inplace replace
    df1.query('data1 == 1', inplace=True)
    print(df1)
