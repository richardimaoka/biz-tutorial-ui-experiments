import { expect, test } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { MarkdownConfigurable } from "./MarkdownConfigurable";

test("markdown comoponent rendered", async () => {
  // Trick to test async server component with react-testing library
  const Component = await MarkdownConfigurable({ markdownBody: "bodybody" });
  // you can't do `render(<MarkdownComponent />)`
  render(Component);

  expect(screen.getByText("bodybody")).toBeDefined();
});
