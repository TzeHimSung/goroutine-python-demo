import numpy as np
import pandas as pd


def main():
    a = np.array([1, 2, 3, 4, 5])
    b = np.array([3, 4, 5, 6, 7])
    df = pd.DataFrame({
        'listA': a,
        'listB': b
    })
    df.to_csv('testProject/result.csv', encoding='utf-8')


if __name__ == '__main__':
    main()
