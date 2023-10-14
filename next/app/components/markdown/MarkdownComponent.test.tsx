/**
 * @jest-environment jsdom
 */
import { render, screen } from "@testing-library/react";
import { MarkdownComponent } from "./MarkdownComponent";

it("Markdown component rendered", () => {
  render(<MarkdownComponent markdownBody="markdownbody" />);
  expect(screen.getByTestId("markdown")).toHaveTextContent("markdownbody");
});
