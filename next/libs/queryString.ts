export const queryParamToString = (
  queryString: string | string[] | undefined
): string | undefined => {
  // if object, it must be string[]
  if (typeof queryString == "object") {
    if (queryString.length > 0) {
      return queryString[0];
    } else {
      return undefined;
    }
  } else if (typeof queryString == "string") {
    return queryString;
  } else {
    return undefined;
  }
};
