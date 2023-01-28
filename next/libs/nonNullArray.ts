const nonNull = <ElementType>(
  element: ElementType | null
): element is ElementType => {
  return element !== null && element != undefined;
};

export const nonNullArray = <ElementType>(arr: (ElementType | null)[]) => {
  return arr.filter(nonNull);
};
