import { renderHook } from "@testing-library/react";
import { expect, test } from "vitest";
import useCounter from "./useCounter";

test("should use counter", () => {
  const { result } = renderHook(() => useCounter());

  expect(result.current.count).toBe(0);
  expect(typeof result.current.increment).toBe("function");
});
