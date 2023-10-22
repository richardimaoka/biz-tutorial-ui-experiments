import { renderHook, waitFor } from "@testing-library/react";
import { expect, test } from "vitest";
import { useEditorInstance } from "./useEditorInstance";
import { EditorBare } from "./EditorBare";
import { render } from "@testing-library/react";

test("useEditorInstance", () => {
  const { result } = renderHook(() => useEditorInstance());
  const [editorInstance, onMount] = result.current;

  render(<EditorBare onDidMount={onMount} />);

  waitFor(() => {
    expect(editorInstance).toBeTruthy();
  });
});
