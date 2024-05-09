export const excludeHelper = <
  T extends Record<string, any>,
  Key extends keyof T,
>(
  obj: T,
  keys: Key[],
): Omit<T, Key> => {
  return Object.keys(obj).reduce(
    (acc, key) => {
      if (!keys.includes(key as Key)) {
        acc[key as keyof Omit<T, Key>] = obj[key];
      }
      return acc;
    },
    {} as Omit<T, Key>,
  );
};
