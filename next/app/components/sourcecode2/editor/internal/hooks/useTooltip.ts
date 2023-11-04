import { editor } from "monaco-editor";
import { ReactNode, useCallback, useEffect, useRef, useState } from "react";
import { Root, createRoot } from "react-dom/client";
import { EditorBoundingRect } from "./useBoundingBox";

/**
 * Custm hook to handle Editor Tooltip
 *
 * @return nothing, as it is an effectful hook
 */
export function useTooltip(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,
  rect: EditorBoundingRect,

  // Prop to be passed from the parent component
  // `tooltip?` (i.e.) Optional parameter, because a React hook should handle
  // if/else *inside*, to avoid conditionals before calling the hook
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
  }
) {
  // React root to set the return of React.createRoot(), and call rootRef.render()
  // See - https://react.dev/reference/react-dom/client/createRoot
  const rootRef = useRef<Root | null>(null);

  // For content widget's root HTML element
  const [contentWidgetContainer] = useState<HTMLDivElement>(() =>
    document.createElement("div")
  );

  const resizeContentWidget = useCallback(() => {
    contentWidgetContainer.style.width = `${rect.width}px`;
  }, [contentWidgetContainer, rect]);

  // Render content widget
  useEffect(() => {
    if (editorInstance && tooltip) {
      // To avoid the following warning, you need to check if createRoot() is already called.
      //   > Warning: You are calling ReactDOMClient.createRoot() on a container that has already been passed to createRoot() before. Instead, call root.render() on the existing root instead if you want to update it.
      if (!rootRef.current) {
        rootRef.current = createRoot(contentWidgetContainer);
      }
      rootRef.current.render(tooltip.children);

      const contentWidget = createWidget(
        contentWidgetContainer,
        tooltip.lineNumber
      );

      editorInstance.addContentWidget(contentWidget);
      return () => {
        editorInstance.removeContentWidget(contentWidget);
      };
    }
  }, [contentWidgetContainer, editorInstance, tooltip]);

  // Resize content widget upon rect change
  useEffect(() => {
    resizeContentWidget();
  }, [resizeContentWidget, rect]);

  return { resizeContentWidget };
}

/**
 *  Required by Monaco-Editor's editorInstance.addContentWidget() method
 *  https://microsoft.github.io/monaco-editor/docs.html#interfaces/editor.ICodeEditor.html
 */
function createWidget(element: HTMLElement, lineNumber: number) {
  return {
    getId: function () {
      return "my.content.widget";
    },
    getDomNode: function () {
      return element;
    },
    getPosition: function () {
      return {
        position: {
          lineNumber: lineNumber,
          column: 1,
        },
        preference: [
          editor.ContentWidgetPositionPreference.BELOW,
          editor.ContentWidgetPositionPreference.ABOVE,
        ],
      };
    },
  };
}
