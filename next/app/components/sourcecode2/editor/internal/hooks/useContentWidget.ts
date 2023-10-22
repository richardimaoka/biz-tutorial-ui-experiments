import { ReactNode, useEffect, useRef, useState } from "react";
import { editor } from "monaco-editor";
import { Root, createRoot } from "react-dom/client";

interface Tooltip {
  lineNumber: number;
  children: ReactNode;
}

type CallbackToResizeContentWidget = (width: number) => void;

/**
 * Custm hook to render tooltip as monaco-editor's content widget, isinde the editor
 *
 * @return nothing, as it is an effectful hook
 */

export function useTooltipWidget(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,

  // tooltip info, possibly undefined
  tooltip?: Tooltip
) {
  // For content widget
  const [contentWidgetContainer] = useState<HTMLDivElement>(() => {
    return document.createElement("div");
  });

  // To be created by React createRoot(), and render the content widget,
  const rootRef = useRef<Root | null>(null);

  // Add content widget
  useEffect(() => {
    if (editorInstance && tooltip) {
      // To avoid the following warning, you need to check if createRoot() is already called.
      //   > Warning: You are calling ReactDOMClient.createRoot() on a container that has already been passed to createRoot() before. Instead, call root.render() on the existing root instead if you want to update it.
      if (!rootRef.current) {
        rootRef.current = createRoot(contentWidgetContainer);
      }
      rootRef.current.render(tooltip.children);

      const contentWidget = createContentWidget(
        contentWidgetContainer,
        tooltip.lineNumber
      );

      editorInstance.addContentWidget(contentWidget);
      return () => {
        editorInstance.removeContentWidget(contentWidget);
      };
    }
  }, [contentWidgetContainer, editorInstance, tooltip]);
}

function createContentWidget(element: HTMLElement, lineNumber: number) {
  return {
    getId: function () {
      return "tooltip.content.widget";
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
