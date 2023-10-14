import { expect, test } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { MarkdownComponent } from "./MarkdownComponent";

test("markdown comoponent rendered", async () => {
  // Trick to test async server component with react-testing library
  const Component = await MarkdownComponent({ markdownBody: "bodybody" });
  // you can't do `render(<MarkdownComponent />)`
  render(Component);

  expect(screen.getByTestId("markdown")).toBeDefined();
});
