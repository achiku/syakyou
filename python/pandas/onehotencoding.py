import pandas as pd


if __name__ == '__main__':
    df1 = pd.DataFrame({
        'data1': range(6),
        'key1': ['A', 'B', 'C', 'A', 'B', 'C'],
        'key2': ['A', 'B', 'C', 'A', 'B', 'C']
        })

    df3 = pd.get_dummies(df1, columns=['key1'], sparse=False)
    print(df3)
