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

    df1.loc[df1['key1'] == 'C', 'key1'] = 0
    print(df1)
    print(df2)
