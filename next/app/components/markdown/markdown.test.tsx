/**
 * @jest-environment jsdom
 */
import { render, screen } from "@testing-library/react";
import { MarkdownC } from "./markdown";

it("App Router: Works with Server Components", () => {
  render(<MarkdownC />);
  expect(screen.getByRole("heading")).toHaveTextContent("Markdown");
});
